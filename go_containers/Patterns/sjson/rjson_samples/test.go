package main

import (
	"fmt"

	"github.com/winterssy/gjson"
)

func main() {
	const dummyData = `
{
  "code": 200,
  "data": {
    "list": [
      {
        "artist": "周杰伦",
        "album": "周杰伦的床边故事",
        "name": "告白气球"
      },
      {
        "artist": "周杰伦",
        "album": "说好不哭 (with 五月天阿信)",
        "name": "说好不哭 (with 五月天阿信)"
      }
    ]
  }
}
`
	var s struct {
		Code int `json:"code"`
		Data struct {
			List gjson.Array `json:"list"`
		} `json:"data"`
	}
	err := gjson.UnmarshalFromString(dummyData, &s)
	if err != nil {
		panic(err)
	}

	fmt.Println(s.Data.List.Index(0).ToObject().GetString("name"))
}

