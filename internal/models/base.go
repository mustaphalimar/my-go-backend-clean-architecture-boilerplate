package models

type Metadata struct {
	Tags       *[]string `json:"tags"`
	Reminder   *string   `json:"reminder"`
	Color      *string   `json:"color"`
	Difficulty *int      `json:"difficulty"`
}
