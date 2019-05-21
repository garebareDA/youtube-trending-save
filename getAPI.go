package main

import(
	"time"
	"net/http"
	"github.com/antonholmquist/jason"
	"log"
)

const getURL = "https://www.googleapis.com/youtube/v3/videos?part=snippet&chart=mostPopular&regionCode=JP&maxResults=25&key=token"

func getYoutubeTrending() {

	list := []channel{}

	db := connectDB()
	defer db.Close()

	res, err := http.Get(getURL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()



	getJson, err := jason.NewObjectFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	items, err := getJson.GetObjectArray("items")

	for _, item := range items {
		t := time.Now()
		const layout = "20060102"

		id, _:= item.GetString("id")
		title, _:= item.GetString("snippet", "title")
		channelTitle, _:= item.GetString("snippet","channelTitle")
		channelID, _ := item.GetString("snippet","channelId")
		description, _ := item.GetString("snippet", "description")

		descriptionCut := isCut(description)
		list = append(list, channel{URL:id, Title:title,ChannelID:channelID, ChannelTitle:channelTitle,Description:descriptionCut, Date:t.Format(layout)})
		log.Println(id, title, channelTitle, channelID, descriptionCut)
	}

	insert(list,db)
}

func trendingTicker() {
	for {
		getYoutubeTrending()
		time.Sleep(24 * time.Hour)
	}
}

func isCut(description string) string {

	if len([]rune(description)) > 105{

		description:= string([]rune(description)[0:105])
		return description

	}

	return description
}
