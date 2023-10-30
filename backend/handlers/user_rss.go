package handlers

import (
	"code-database/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/feeds"
)

func RssHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		knowledges, err := models.GetAllPublishedKnowledges()
		if err != nil {
			log.Print(err.Error())
			// エラー処理
		}
		var feedItems []*feeds.Item
		for _, k := range knowledges {
			feedItem := &feeds.Item{
				Id:          strconv.Itoa(k.ID),
				Title:       k.Title,
				Link:        &feeds.Link{Href: fmt.Sprintf("https://code-database.com/knowledges/%d", k.ID)},
				Description: k.Title,
				Created:     time.Now(),
			}
			feedItems = append(feedItems, feedItem)
		}
		feed := &feeds.Feed{
			Title:       "Code database Knowledges Feed",
			Link:        &feeds.Link{Href: "https://code-database.com/"},
			Description: "",
			Author:      &feeds.Author{Name: "Rei Sugiura and Tomoya Imamura"},
			Created:     time.Now(),
			Items:       feedItems,
		}
		feed.WriteRss(w)
		// w.Header().Set("Content-Type", "application/xml")
	} else {
		return
	}
}
