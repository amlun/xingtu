package api

import (
	"github.com/dghubble/sling"
	"github.com/qianshouapp/xingtu/model"
)

const (
	// page = %d, tag = 1,72
	// ?limit=20&need_detail=true&page=%d&platform_source=1&order_by=score&disable_replace_keyword=false&marketing_target=1&task_category=1&tag=%s&is_filter=true
	AuthorListApi = "https://www.xingtu.cn/v/api/"
)

type GetAuthorListQuery struct {
	Limit        int    `url:"limit"`
	Page         int    `url:"page"`
	TaskCategory int    `url:"task_category"`
	Tag          string `url:"tag"`
}

type getAuthorListResponse struct {
	BaseResponse
	Data *struct {
		Authors []model.Author `json:"authors"`
	} `json:"data"`
}

func (c *Client) GetAuthorList(tag string, page int) (authors []model.Author, err error) {
	base := sling.New().Base(AuthorListApi).Client(authClient)
	q := &GetAuthorListQuery{
		Limit:        20,
		Page:         page,
		Tag:          tag,
		TaskCategory: 1,
	}
	resp := new(getAuthorListResponse)
	_, err = base.New().Get("demand/author_list").QueryStruct(q).ReceiveSuccess(resp)
	if err != nil {
		return nil, err
	}
	if resp.Data != nil && resp.Data.Authors != nil {
		authors = resp.Data.Authors
	}
	return authors, nil
}
