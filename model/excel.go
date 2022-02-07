package model

import "fmt"

type Star struct {
	Id   string // 作者ID
	Name string // 作者名称
	Play    int // 播放数
	Like    int // 点赞数
	Comment int // 评论数
	//Title   string // 视频标题
	//IsHot   bool   // 是否热门
	//Share   int    // 分享数
	Url string // 视频链接
}

func (a Star) String() string {
	return fmt.Sprintf("%s,%s,%d,%d,%d,%s", a.Id, a.Name, a.Play, a.Like, a.Comment, a.Url)
}
