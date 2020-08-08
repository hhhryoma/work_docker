package main

import (
	"fmt"
	"net/http"

	"github.com/ChimeraCoder/anaconda"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/tweet", search)
	e.POST("/tweet/post", doTweet)
	e.Logger.Fatal(e.Start(":1323"))
}

type Tweet struct {
	User string `json: "user"`
	Text string `json: "text"`
}

// Tweetの配列
type Tweets *[]Tweet

func connectTwitterApi() *anaconda.TwitterApi {
	return anaconda.NewTwitterApiWithCredentials("1095304436660940807-EvZaso6B7tXZ8GMALsbjP4xXGECYom",
		"okOHnBDjuG13P8iEjlz6kGewYkAIFCAW8wzaIeAw6VrnA",
		"7LymMc7LV1Pcc6ehQa23RwnLU",
		"bXinljyaCPhlgK37B70QWD7jdRQZgn9nPFaoHodvbw2oxcLzEJ")

}

func search(c echo.Context) error {
	keyword := c.FormValue("keyword")
	api := connectTwitterApi()

	searchResult, _ := api.GetSearch(`"`+keyword+`"`, nil)
	tweets := make([]*Tweet, 0)

	for _, data := range searchResult.Statuses {
		fmt.Println(data.FullText)
		tweet := new(Tweet)
		tweet.User = data.User.Name
		tweet.Text = data.FullText

		tweets = append(tweets, tweet)

	}
	return c.JSON(http.StatusOK, tweets)
}

func doTweet(c echo.Context) error {
	api := connectTwitterApi()
	text := c.FormValue("text")
	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println(tweet.Text)
	return nil
}
