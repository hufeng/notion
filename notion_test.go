package notion

import "testing"

func TestNotionNew(t *testing.T) {
	client := NewClient()

	if client.baseUrl != baseUrl {
		t.Errorf("Expected %s, got %s", baseUrl, client.baseUrl)
	}
	if client.notionVersion != notionVersion {
		t.Errorf("Expected %s, got %s", notionVersion, client.notionVersion)
	}
}

func TestNotionNewWithSetParam(t *testing.T) {
	testUrl := "https://api-test.notion.com"
	testNotionVersion := "2022-06-29"

	client := NewClient().SetBaseUrl(testUrl).SetNotionVersion(testNotionVersion)

	if client.baseUrl != testUrl {
		t.Errorf("Expected %s, got %s", testUrl, client.baseUrl)
	}
	if client.notionVersion != testNotionVersion {
		t.Errorf("Expected %s, got %s", testNotionVersion, client.notionVersion)
	}
}
