package model

type PriceInfo []PriceDetail

type PriceDetail struct {
	Price     int    `json:"price"`
	Desc      string `json:"desc"`
	VideoType int    `json:"video_type"`
}
