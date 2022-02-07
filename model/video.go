package model

type Video struct {
	Id      string `json:"item_id"`
	Title   string `json:"title"`
	Play    int    `json:"play"`
	Like    int    `json:"like"`
	Comment int    `json:"comment"`
	IsHot   bool   `json:"is_hot"`
	Share   int    `json:"share"`
	Url     string `json:"url"`
}
