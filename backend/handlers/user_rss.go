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
	layout := "2006-01-02 03:04:05"
	if r.Method == "GET" {
		knowledges, err := models.GetAllPublishedKnowledges()
		if err != nil {
			log.Print(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var feedItems []*feeds.Item
		for _, k := range knowledges {
			created, _ := time.Parse(layout, k.CreatedAt)
			updated, _ := time.Parse(layout, k.UpdatedAt)
			feedItem := &feeds.Item{
				Id:          strconv.Itoa(k.ID),
				Title:       k.Title,
				Link:        &feeds.Link{Href: fmt.Sprintf("https://code-database.com/knowledges/%d", k.ID)},
				Description: k.Title,
				Content:     k.Content,
				Created:     created,
				Updated:     updated,
			}
			feedItems = append(feedItems, feedItem)
		}
		feed := &feeds.Feed{
			Title:       "Code database Knowledges Feed",
			Link:        &feeds.Link{Href: "https://code-database.com/"},
			Description: "Code Databaseはプログラミング中級者向けのサイトです。すぐに開発に応用できる知識を、サンプルコードを用いた丁寧な解説と一緒に記事にしてお届けしています。",
			Author:      &feeds.Author{Name: "Rei Sugiura and Tomoya Imamura"},
			Created:     time.Now(),
			Items:       feedItems,
			Image:       &feeds.Image{Url: "https://s3-ap-northeast-1.amazonaws.com/code-database.com/images/code-database-ogp.png"},
		}
		log.Print(feed.ToRss())
		feed.WriteRss(w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
