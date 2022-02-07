package api

import "testing"

func TestClient_GetAuthorWatchedDistribution(t *testing.T) {
	c := Client{}
	distributions, err := c.GetAuthorWatchedDistribution("6774630017209991179")
	if err != nil {
		t.Error(err)
	}
	t.Log(distributions)
}

func TestClient_GetAuthorShowItems(t *testing.T) {
	c := Client{}
	videos, err := c.GetAuthorShowItems("6774630017209991179")
	if err != nil {
		t.Error(err)
	}
	t.Log(videos)
}