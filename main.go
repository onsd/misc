package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Follower struct {
	Login string `json:"login"`
	//ID                int    `json:"id"`
	//NodeID            string `json:"node_id"`
	AvatarURL string `json:"avatar_url"`
	//GravatarID        string `json:"gravatar_id"`
	/*
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	*/
}

func fetch(username string) ([]Follower, error) {
	url := "https://api.github.com/users/" + username + "/followers"
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
		return nil, fmt.Errorf("Unable to get this url : http status %d", res.StatusCode)
	}
	defer res.Body.Close()

	// jsonを読み込む
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// jsonを構造体へデコード
	var articles []Follower
	if err := json.Unmarshal(body, &articles); err != nil {
		return nil, err
	}
	return articles, nil
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		/*if c.Req.Method == "OPTIONS" {
			fmt.Println("options")
			c.Abort(200)
			return
		}*/

		c.Next()
	}
}
func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		follower, err := fetch(name)
		if err != nil {
			//log.Fatalf("Error!: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
			return
		}

		//構造体をJSONに戻す
		var buf bytes.Buffer
		b, _ := json.Marshal(follower)
		buf.Write(b)
		c.String(http.StatusOK, buf.String())
	})
	router.Run(":8080")
}
