import (
  	"encoding/json"
  	"fmt"
  	"io/ioutil"
	"log"
	"net/http"
	"time"
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

type Movie struct{
	Text string `json:"text"`
	Attachments []struct {
		Photo struct {
				Photo75 string `json:"photo_75"`
				Photo130 string `json:"photo_130"`
				Photo604 string `json:"photo_604"`
				Photo807 string `json:"photo_807"`
				Photo1280 string `json:"photo_1280"`
			} 
		} 
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

func DECODEJSON() Movie{
	var m AutoGenerated
	data:=GETHTTP("https://api.vk.com/method/wall.get?domain=cinemavilyui&count=1&filter=owner&extended=0&v=5.103&access_token=TEXT")

	err := json.Unmarshal(data, &m)
	if err != nil {
		fmt.Println("error:", err)
	}

	var movie Movie
	
	movie.Attachments=make([]struct {

	Photo struct {
			Photo75 string `json:"photo_75"`
			Photo130 string `json:"photo_130"`
			Photo604 string `json:"photo_604"`
			Photo807 string `json:"photo_807"`
			Photo1280 string `json:"photo_1280"`
			} 
	} ,len(m.Response.Items[0].Attachments))

		for i := 0; i < len(m.Response.Items[0].Attachments); i++ {
			movie.Attachments[i].Photo.Photo75=m.Response.Items[0].Attachments[i].Photo.Photo75
			movie.Attachments[i].Photo.Photo130=m.Response.Items[0].Attachments[i].Photo.Photo130
			movie.Attachments[i].Photo.Photo604=m.Response.Items[0].Attachments[i].Photo.Photo604
			movie.Attachments[i].Photo.Photo807=m.Response.Items[0].Attachments[i].Photo.Photo807
			movie.Attachments[i].Photo.Photo1280=m.Response.Items[0].Attachments[i].Photo.Photo1280
		}
	movie.Text=m.Response.Items[0].Text
return movie
}

func MARSHALL_HTTP(w http.ResponseWriter, r *http.Request) {
    movie:=DECODEJSON()
	b, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("error:", err)
	}	 
		w.Header().Set("Content-Type", "application/json charset=utf-8")
		w.Write(b)
}		

func main() {

    ticker := time.NewTicker(time.Millisecond * 3600000)//каждый час
		for t:= range ticker.C {
			http.HandleFunc("/", MARSHALL_HTTP)
		        http.ListenAndServe(":8080", nil)
			fmt.Println("debug Tick", t)
		}
}