package notion

const (
	baseUrl       = "https://api.notion.com"
	notionVersion = "2022-06-28"
)

type client struct {
	baseUrl       string
	notionVersion string
}

func NewClient() *client {
	return &client{
		baseUrl:       baseUrl,
		notionVersion: notionVersion,
	}
}

func (c *client) SetBaseUrl(url string) *client {
	c.baseUrl = url
	return c
}

func (c *client) SetNotionVersion(version string) *client {
	c.notionVersion = version
	return c
}
