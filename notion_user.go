package notion

import (
	"context"
	"fmt"
)

type Base struct {
	Object    string `json:"object"`
	Id        string `json:"id"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
}

type User struct {
	Base
	Person struct {
		Email string `json:"email"`
	} `json:"person"`
}

type Bot struct {
	Base
	Bot struct {
		Owner struct {
			Type      string `json:"type"`
			Workspace bool   `json:"workspace"`
		} `json:"owner"`
	} `json:"bot"`
}

type UserService service

func (u *UserService) retrive(ctx context.Context, userId string) (*Base, error) {
	var user Base
	err := u.client.get(fmt.Sprintf("/users/%s", userId), nil, &user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (u *UserService) retriveBot(ctx context.Context, botId string) (*Bot, error) {
	var bot Bot
	err := u.client.get(fmt.Sprintf("/users/%s", botId), nil, &bot)
	if err != nil {
		return nil, err
	}
	return &bot, err
}
