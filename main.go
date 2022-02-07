package main

import (
	"encoding/json"
	"fmt"
	"github.com/qianshouapp/xingtu/api"
	"github.com/xuri/excelize/v2"
	"strings"
)

func main() {
	// 创建Excel
	f := excelize.NewFile()
	streamWriter, err := f.NewStreamWriter("Sheet1")
	if err != nil {
		fmt.Println(err)
	}
	// A, B, C, D, E, F ,G, H, I, J
	streamWriter.SetRow("A1", []interface{}{"ID", "达人名称", "达人标签", "视频报价1-20s", "视频报价21-60s", "视频报价60s+", "播放量", "点赞量", "评论量", "视频链接"})
	index := 2
	client := new(api.Client)
	// 表头
	// a.Id, a.Name, a.Title, a.IsHot, a.Play, a.Like, a.Comment, a.Share
	var start, end int
	for i := 1; i <= 100; i++ {
		authors, err := client.GetAuthorList("36", i)
		if err != nil {
			break
		}
		for _, author := range authors {
			var priceArr = make([]int, 3)
			var tagArr []string
			for _, price := range author.PriceInfo {
				if price.VideoType == 1 {
					priceArr[0] = price.Price
				}
				if price.VideoType == 2 {
					priceArr[1] = price.Price
				}
				if price.VideoType == 71 {
					priceArr[2] = price.Price
				}
			}
			json.Unmarshal([]byte(author.Tags), &tagArr)
			tags := strings.Join(tagArr, ",")
			// 个人信息
			distributions, err := client.GetAuthorWatchedDistribution(author.Id)
			if err != nil {
				continue
			}
			for _, distribution := range distributions {
				// 性别分布
				if distribution.Type == 1 && len(distribution.Image) > 0 && distribution.Image[0] != "女性居多" {
					continue
				}
			}

			// 视频列表
			videos, err := client.GetAuthorShowItems(author.Id)
			if err != nil {
				continue
			}
			// 打印视频
			start = index
			for _, video := range videos {
				streamWriter.SetRow(fmt.Sprintf("A%d", index), []interface{}{author.Id, author.NickName, tags, priceArr[0], priceArr[1], priceArr[2], video.Play, video.Like, video.Comment, video.Url})
				end = index
				index++
			}
			fmt.Println(start, end)
			// 按达人合并单元格
			streamWriter.MergeCell(fmt.Sprintf("A%d", start), fmt.Sprintf("A%d", end))
			streamWriter.MergeCell(fmt.Sprintf("B%d", start), fmt.Sprintf("B%d", end))
			streamWriter.MergeCell(fmt.Sprintf("C%d", start), fmt.Sprintf("C%d", end))
			streamWriter.MergeCell(fmt.Sprintf("D%d", start), fmt.Sprintf("D%d", end))
			streamWriter.MergeCell(fmt.Sprintf("E%d", start), fmt.Sprintf("E%d", end))
			streamWriter.MergeCell(fmt.Sprintf("F%d", start), fmt.Sprintf("F%d", end))
		}
	}
	// 保存Excel
	if err := streamWriter.Flush(); err != nil {
		fmt.Println(err)
	}
	if err := f.SaveAs("抖音星达人-生活.xlsx"); err != nil {
		fmt.Println(err)
	}
}