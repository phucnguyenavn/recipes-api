package model

import "time"

type Recipe struct {
	id string `json:"id"`

	name string `json:"name"`

	tags []string `json:"tags"`

	ingredients []string `json:"ingredients"`

	instructions []string `json:"instructions"`

	publishedAt time.Time `json:"publishedAt"`
}

func (r *Recipe) Id() string {
	return r.id
}

func (r *Recipe) SetId(id string) {
	r.id = id
}

func (r *Recipe) Name() string {
	return r.name
}

func (r *Recipe) SetName(name string) {
	r.name = name
}

func (r *Recipe) Tags() []string {
	return r.tags
}

func (r *Recipe) SetTags(tags []string) {
	r.tags = tags
}

func (r *Recipe) Ingredients() []string {
	return r.ingredients
}

func (r *Recipe) SetIngredients(ingredients []string) {
	r.ingredients = ingredients
}

func (r *Recipe) Instructions() []string {
	return r.instructions
}

func (r *Recipe) SetInstructions(instructions []string) {
	r.instructions = instructions
}

func (r *Recipe) PublishedAt() time.Time {
	return r.publishedAt
}

func (r *Recipe) SetPublishedAt(publishedAt time.Time) {
	r.publishedAt = publishedAt
}
