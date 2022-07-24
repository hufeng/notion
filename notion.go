package notion

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	baseUrl       = "https://api.notion.com"
	notionVersion = "2022-06-28"
	apiVersion    = "v1"
)

type Client struct {
	client        *http.Client
	baseUrl       string
	notionVersion string
	apiVersion    string
	Token         string

	common   service
	Database *Database
	Page     *Page
	Block    *Block
	Comment  *Comment
	Search   *Search
	User     *User
}

type service struct {
	client *Client
}

func NewClient() *Client {
	c := &Client{
		client:        http.DefaultClient,
		baseUrl:       baseUrl,
		notionVersion: notionVersion,
		apiVersion:    apiVersion,
	}
	c.common.client = c
	c.Database = (*Database)(&c.common)
	c.Page = (*Page)(&c.common)
	c.Block = (*Block)(&c.common)
	c.Comment = (*Comment)(&c.common)
	c.Search = (*Search)(&c.common)
	c.User = (*User)(&c.common)
	return c
}

func (c *Client) SetBaseUrl(url string) *Client {
	c.baseUrl = url
	return c
}

func (c *Client) SetNotionVersion(version string) *Client {
	c.notionVersion = version
	return c
}

type req struct {
	method string
	header map[string]any
	query  map[string]any
	body   any
}

func (c *Client) req(ctx context.Context, path string, param *req, result any) error {
	api := c.baseUrl + "/" + c.apiVersion + "/" + path
	u, err := url.Parse(api)
	if err != nil {
		return err
	}

	// set query
	q := u.Query()
	for k, v := range param.query {
		q.Add(k, v.(string))
	}

	// set body
	var buf io.Reader
	if param != nil && param.body != nil {
		body, err := json.Marshal(param.body)
		if err != nil {
			return err
		}
		buf = bytes.NewBuffer(body)
	}

	// set method
	method := param.method
	if method == "" {
		method = http.MethodGet
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return err
	}

	// set header
	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Notion-Version", c.notionVersion)

	if param != nil && param.header != nil {
		for k, v := range param.header {
			req.Header.Add(k, v.(string))
		}
	}

	res, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("network error status code %d with error %s ", res.StatusCode, string(b))
	}

	err = json.Unmarshal(b, result)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) get(path string, param map[string]any, result any) error {
	return c.req(context.Background(), path, &req{
		method: http.MethodGet,
		query:  param,
	}, result)
}

func (c *Client) post(path string, body any, result any) error {
	return c.req(context.Background(), path, &req{
		method: http.MethodPost,
		body:   body,
	}, result)
}
