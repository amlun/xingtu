package api

import (
	"github.com/dghubble/sling"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

var authClient = &http.Client{
	Timeout: 5 * time.Second,
}

func init() {
	u := &url.URL{
		Scheme: "https",
		Host:   "www.xingtu.cn",
	}
	cookies := []*http.Cookie{
		{
			Name:  "sessionid",
			Value: "923ca131b8ea701b968b6466d3bc6e57",
		},
		{
			Name:  "gftoken",
			Value: "OTIzY2ExMzFiOHwxNjM0NTQ0NTEyMjR8fDAGBgYGBgY",
		},
	}
	authClient.Jar, _ = cookiejar.New(nil)
	authClient.Jar.SetCookies(u, cookies)

}

type Client struct {
	sling *sling.Sling
}

type BaseResponse struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}
