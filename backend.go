package main

import (
  	"encoding/json"
  	"fmt"
  	"io/ioutil"
	"log"
	"net/http"
)
type AutoGenerated struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			ID int `json:"id"`
			FromID int `json:"from_id"`
			OwnerID int `json:"owner_id"`
			Date int `json:"date"`
			PostType string `json:"post_type"`
			Text string `json:"text"`
			IsPinned int `json:"is_pinned"`
			Attachments []struct {
				Type string `json:"type"`
				Photo struct {
					ID int `json:"id"`
					AlbumID int `json:"album_id"`
					OwnerID int `json:"owner_id"`
					UserID int `json:"user_id"`
					Photo75 string `json:"photo_75"`
					Photo130 string `json:"photo_130"`
					Photo604 string `json:"photo_604"`
					Photo807 string `json:"photo_807"`
					Photo1280 string `json:"photo_1280"`
					Width int `json:"width"`
					Height int `json:"height"`
					Text string `json:"text"`
					Date int `json:"date"`
					AccessKey string `json:"access_key"`
				} `json:"photo"`
			} `json:"attachments"`
			Comments struct {
				Count int `json:"count"`
			} `json:"comments"`
			Likes struct {
				Count int `json:"count"`
			} `json:"likes"`
			Reposts struct {
				Count int `json:"count"`
			} `json:"reposts"`
		} `json:"items"`
	} `json:"response"`
}

func GETHTTP(data string)[]byte{
	res, err := http.Get(data)
	if err != nil {
		log.Fatal(err)
	}
	response, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return response
}

func main() {
	data:=GETHTTP("https://api.vk.com/method/wall.get?domain=cinemavilyui&count=1&filter=owner&extended=0&v=5.45")
	// fmt.Println(string(data))
	var m AutoGenerated
	err := json.Unmarshal(data, &m)
	if err != nil {
		fmt.Println("error:", err)
	}
		fmt.Println(m.Response.Items[0].Text)
		fmt.Println(len(m.Response.Items[0].Attachments))


}