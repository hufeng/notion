package notion

import (
	"fmt"
)

type User struct {
	Object    string `json:"object"`
	Id        string `json:"id"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
	Person    struct {
		Email string `json:"email,omitempty"`
	}
	Bot struct {
		Owner struct {
			Type      string `json:"type,omitempty"`
			Workspace bool   `json:"workspace,omitempty"`
		} `json:"owner,omitempty"`
	} `json:"bot,omitempty"`
}

type UserApi interface {
	List()
	Retrive(string) (*User, error)
	RtriveBot(string) (*User, error)
}

type UserService service

type UserPagination struct {
	Object     string  `json:"object"`
	Results    []*User `json:"results"`
	NextCursor string  `json:"next_cursor"`
	HasMore    bool    `json:"has_more"`
	Type       string  `json:"type"`
}

func (u *UserService) List(start string, size int) (*UserPagination, error) {
	var result UserPagination
	err := u.client.get("/users", query{
		"start_cursor": start,
		"page_size":    fmt.Sprintf("%d", size),
	}, &result)
	return &result, err
}

func (u *UserService) Retrive(userId string) (*User, error) {
	var user User
	err := u.client.get(fmt.Sprintf("/users/%s", userId), nil, &user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (u *UserService) RetriveBot(botId string) (*User, error) {
	var user User
	err := u.client.get(fmt.Sprintf("/users/%s", botId), nil, &user)
	if err != nil {
		return nil, err
	}
	return &user, err
}
