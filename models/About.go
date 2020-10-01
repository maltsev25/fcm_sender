package models

type Route struct {
	Method string `json:"method"`
	URL    string `json:"url"`
	Fields []Field
}

type Field struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
}
