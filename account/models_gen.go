// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package account

type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AccountInput struct {
	Name string `json:"name"`
}

type PaginationInput struct {
	Skip *int `json:"skip"`
	Take *int `json:"take"`
}