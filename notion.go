package notion

import (
	"bytes"
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
	token         string

	common    service
	Databases *DatabaseService
	Pages     *PageService
	Blocks    *BlockService
	Comments  *CommentService
	Search    *SearchService
	Users     *UserService
}

type service struct {
	client *Client
}

func NewClient(token string) *Client {
	c := &Client{
		client:        http.DefaultClient,
		baseUrl:       baseUrl,
		notionVersion: notionVersion,
		apiVersion:    apiVersion,
		token:         token,
	}

	c.common.client = c
	c.Databases = (*DatabaseService)(&c.common)
	c.Pages = (*PageService)(&c.common)
	c.Blocks = (*BlockService)(&c.common)
	c.Comments = (*CommentService)(&c.common)
	c.Search = (*SearchService)(&c.common)
	c.Users = (*UserService)(&c.common)

	return c
}

func (c *Client) SetBaseUrl(url string) *Client {
	c.baseUrl = url
	return c
}

func (c *Client) SetApiVersion(v string) *Client {
	c.apiVersion = v
	return c
}

func (c *Client) SetNotionVersion(version string) *Client {
	c.notionVersion = version
	return c
}

type query map[string]string
type header map[string]string
type reqParam struct {
	method string
	header header
	query  query
	body   any
}

func (c *Client) req(path string, param *reqParam, result any) error {
	api := c.baseUrl + "/" + c.apiVersion + path

	u, err := url.Parse(api)
	if err != nil {
		return err
	}

	// set query
	q := u.Query()
	if param != nil && param.query != nil {
		for k, v := range param.query {
			q.Add(k, v)
		}
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

	// make request
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return err
	}

	// set default header
	req.Header.Add("Authorization", "Bearer "+c.token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Notion-Version", c.notionVersion)

	// set extend header
	if param != nil && param.header != nil {
		for k, v := range param.header {
			req.Header.Add(k, v)
		}
	}

	// send request
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("network error status code %d with error %s ", res.StatusCode, string(b))
	}

	err = json.Unmarshal(b, result)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) get(path string, param query, result any) error {
	return c.req(path, &reqParam{
		method: http.MethodGet,
		query:  param,
	}, result)
}

func (c *Client) post(path string, body any, result any) error {
	return c.req(path, &reqParam{
		method: http.MethodPost,
		body:   body,
	}, result)
}
