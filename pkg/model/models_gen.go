// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewKey struct {
	Content string `json:"content"`
	Name    string `json:"name"`
	UserID  string `json:"userId"`
}

type NewUser struct {
	Name     *string `json:"name"`
	Username string  `json:"username"`
}
