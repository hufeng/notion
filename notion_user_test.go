package notion

import (
	"fmt"
	"os"
	"testing"
)

var (
	NotionToken = os.Getenv("NOTION_TOKEN")
	NotionBotId = os.Getenv("NOTION_BOT_ID")
)

func TestUserServiceRetrive(t *testing.T) {
	notion := NewClient(NotionToken)
	user, err := notion.Users.Retrive(NotionBotId)
	if err != nil {
		t.Fatalf("error %s\n", err)
	}
	fmt.Printf("%+v\n", user)
}

func TestUserServiceRetriveBot(t *testing.T) {
	notion := NewClient(NotionToken)
	bot, err := notion.Users.RetriveBot(NotionBotId)
	if err != nil {
		t.Fatalf("error %s\n", err)
	}
	fmt.Printf("%+v\n", bot)
}

func TestUserServiceList(t *testing.T) {
	notion := NewClient(NotionToken)
	users, err := notion.Users.List("", 100)
	if err != nil {
		t.Fatalf("user list error %s", err)
	}
	fmt.Printf("%#v\n", users)
}
