package model

import "time"

type Recipe struct {
	name string `json:"name"`

	tags []string `json:"tags"`

	ingredients []string `json:"ingredients"`

	instructions []string `json:"instructions"`

	publishedAt time.Time `json:"publishedAt"`
}
