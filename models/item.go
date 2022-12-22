package models

import (
	"fmt"
	"net/http"
)

type Item struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type ItemList struct {
	Items []Item `json:"items"`
}

func (item *Item) Bind(r *http.Request) error {
	if item.Name == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}

func (item *Item) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (item *ItemList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
