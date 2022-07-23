package notion

const (
	baseUrl       = "https://api.notion.com"
	notionVersion = "2022-06-28"
	apiVersion    = "v1"
)

type Client struct {
	baseUrl       string
	notionVersion string
}

func NewClient() *Client {
	return &Client{
		baseUrl:       baseUrl,
		notionVersion: notionVersion,
	}
}

func (c *Client) SetBaseUrl(url string) *Client {
	c.baseUrl = url
	return c
}

func (c *Client) SetNotionVersion(version string) *Client {
	c.notionVersion = version
	return c
}
