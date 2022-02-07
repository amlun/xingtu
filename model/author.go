package model

type Author struct {
	Id                    string    `json:"id"`                         // 抖音ID
	NickName              string    `json:"nick_name"`                  // 昵称
	Gender                int       `json:"gender"`                     // 性别
	Tags                  string    `json:"tags"`                       // 标签
	PriceInfo             PriceInfo `json:"price_info"`                 // 价格信息
	Follower              int       `json:"follower"`                   // 粉丝数
	AvgPlay               int       `json:"avg_play"`                   // 平均播放量
	ExpectedPlayNum       int       `json:"expected_play_num"`          // 预期播放量
	InteractRateWithin30d float32   `json:"interact_rate_within_30d"`   // 30天的互动率
	PlayOverRateWithin30d float32   `json:"play_over_rate_within_30_d"` // 30天的完播率
}
