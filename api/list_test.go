package api

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_GetAuthorList(t *testing.T) {
	c := Client{}
	authors, err := c.GetAuthorList("1", 1)
	if err != nil {
		t.Error(err)
	}
	for _, author := range authors {
		fmt.Println(author.Tags)

		var tagArr []string
		json.Unmarshal([]byte(author.Tags), &tagArr)
		fmt.Println(tagArr)

		break
	}

}
