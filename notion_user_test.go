package notion

import (
	"context"
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
	user, err := notion.Users.retrive(context.Background(), NotionBotId)
	if err != nil {
		t.Fatalf("error %s\n", err)
	}
	fmt.Printf("%+v\n", user)
}

func TestUserServiceRetriveBot(t *testing.T) {
	notion := NewClient(NotionToken)
	bot, err := notion.Users.retriveBot(context.Background(), NotionBotId)
	if err != nil {
		t.Fatalf("error %s\n", err)
	}
	fmt.Printf("%+v\n", bot)
}
