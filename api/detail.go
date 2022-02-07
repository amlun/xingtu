package api

import (
	"crypto/md5"
	"fmt"
	"github.com/dghubble/sling"
	"github.com/qianshouapp/xingtu/model"
)

const (
	AuthorDetailApi = "https://www.xingtu.cn/h/api/"

	// "service_methodIMTotalUnreadservice_namegeneric.AdStarGenericServicesign_strict1" + "e39539b8836fb99e1538974d3ac1fe98"
	SecretKey = "e39539b8836fb99e1538974d3ac1fe98"
	// l = (7) ['o_author_id', 'platform_source', 'platform_channel', 'type', 'service_name', 'service_method', 'sign_strict']
	/*
		o_author_id: "6774630017209991179"
		platform_channel: 1
		platform_source: 1
		service_method: "GetAuthorWatchedDistribution"
		service_name: "data.AdStarDataService"
		sign_strict: 1
		type: 1

		"o_author_id6774630017209991179platform_channel1platform_source1service_methodGetAuthorWatchedDistributionservice_namedata.AdStarDataServicesign_strict1type1e39539b8836fb99e1538974d3ac1fe98"
	*/
)

type GetAuthorWatchedDistributionQuery struct {
	AuthorId        string `url:"o_author_id"`
	PlatformSource  string `url:"platform_source"`
	PlatformChannel string `url:"platform_channel"`
	Type            string `url:"type"`
	ServiceName     string `url:"service_name"`   // data.AdStarDataService
	ServiceMethod   string `url:"service_method"` // GetAuthorWatchedDistribution
	SignStrict      string `url:"sign_strict"`
	Sign            string `url:"sign"`
}

func (q *GetAuthorWatchedDistributionQuery) genSign() string {
	var str = "o_author_id" + q.AuthorId + "platform_channel" + q.PlatformChannel + "platform_source" + q.PlatformSource +
		"service_method" + q.ServiceMethod + "service_name" + q.ServiceName + "sign_strict" + q.SignStrict + "type" + q.Type + SecretKey
	re := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", re)
}

type GetAuthorWatchedDistributionResponse struct {
	BaseResponse
	Data *struct {
		Distributions []model.Distribution `json:"distributions"`
	}
}

/*
	o_author_id: "6774630017209991179"
	platform_channel: 1
	platform_source: 1
	service_method: "GetAuthorWatchedDistribution"
	service_name: "data.AdStarDataService"
	sign_strict: 1
	type: 1

	"o_author_id6774630017209991179platform_channel1platform_source1service_methodGetAuthorWatchedDistributionservice_namedata.AdStarDataServicesign_strict1type1e39539b8836fb99e1538974d3ac1fe98"
*/
func (c *Client) GetAuthorWatchedDistribution(id string) (distributions []model.Distribution, err error) {
	base := sling.New().Base(AuthorDetailApi).Client(authClient)
	q := &GetAuthorWatchedDistributionQuery{
		AuthorId:        id,
		PlatformSource:  "1",
		PlatformChannel: "1",
		Type:            "1",
		ServiceName:     "data.AdStarDataService",
		ServiceMethod:   "GetAuthorWatchedDistribution",
		SignStrict:      "1",
	}
	q.Sign = q.genSign()
	resp := new(GetAuthorWatchedDistributionResponse)
	_, err = base.New().Get("gateway/handler_get/").QueryStruct(q).ReceiveSuccess(resp)
	if err != nil {
		return nil, err
	}
	if resp.Data != nil && resp.Data.Distributions != nil {
		distributions = resp.Data.Distributions
	}
	return distributions, nil
}

type GetAuthorShowItemsQuery struct {
	AuthorId        string `url:"o_author_id"`
	PlatformSource  string `url:"platform_source"`
	PlatformChannel string `url:"platform_channel"`
	Limit           string `url:"limit"`
	Type            string `url:"type"`
	ServiceName     string `url:"service_name"`   // data.AdStarDataService
	ServiceMethod   string `url:"service_method"` // GetAuthorShowItemsV2
	SignStrict      string `url:"sign_strict"`
	Sign            string `url:"sign"`
}

func (q *GetAuthorShowItemsQuery) genSign() string {
	var str = "limit" + q.Limit + "o_author_id" + q.AuthorId + "platform_channel" + q.PlatformChannel + "platform_source" + q.PlatformSource +
		"service_method" + q.ServiceMethod + "service_name" + q.ServiceName + "sign_strict" + q.SignStrict + SecretKey
	re := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", re)
}

type GetAuthorShowItemsResponse struct {
	BaseResponse
	Data *struct {
		Videos []model.Video `json:"latest_item_info"`
	}
}

/*
	o_author_id: 6774630017209991179
	platform_source: 1
	platform_channel: 1
	limit: 10
	service_name: author.AdStarAuthorService
	service_method: GetAuthorShowItemsV2
	sign_strict: 1
	sign: 44e1e53565a9105a2c6207fc90bbb9c1

	"limit10o_author_id6774630017209991179platform_channel1platform_source1service_methodGetAuthorShowItemsV2service_nameauthor.AdStarAuthorServicesign_strict1"
*/
func (c *Client) GetAuthorShowItems(id string) (videos []model.Video, err error) {
	base := sling.New().Base(AuthorDetailApi).Client(authClient)
	q := &GetAuthorShowItemsQuery{
		AuthorId:        id,
		PlatformSource:  "1",
		PlatformChannel: "1",
		Limit:           "10",
		Type:            "1",
		ServiceName:     "author.AdStarAuthorService",
		ServiceMethod:   "GetAuthorShowItemsV2",
		SignStrict:      "1",
	}
	q.Sign = q.genSign()
	resp := new(GetAuthorShowItemsResponse)
	_, err = base.New().Get("gateway/handler_get/").QueryStruct(q).ReceiveSuccess(resp)
	if err != nil {
		return nil, err
	}
	if resp.Data != nil && resp.Data.Videos != nil {
		videos = resp.Data.Videos
	}
	return videos, nil
}
