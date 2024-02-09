package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"code-database/config"
	"code-database/models"
)

// AdminSaveHandler /admin/saveに対するハンドラ
func AdminSaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	category := r.FormValue("category")
	content := r.FormValue("content")
	rowContent := r.FormValue("row_content")
	eyecatchSrc := r.FormValue("eyecatch_src")
	switch {
	case r.Method == "POST":
		createdAt := time.Now()
		updatedAt := time.Now()
		knowledgeID, err := models.PostKnowledge(title, content, rowContent, createdAt, updatedAt, eyecatchSrc, category)
		if err != nil {
			log.Print(err.Error())
			return
		}
		if r.FormValue("tags") != "" {
			tags := strings.Split(r.FormValue("tags"), ",")
			for _, tag := range tags {
				tagID, _ := strconv.Atoi(tag)
				err := models.PostKnowledgesTags(int(knowledgeID), tagID, createdAt, updatedAt)
				if err != nil {
					log.Print(err.Error())
					return
				}
			}
		}
	case r.Method == "PUT":
		knowledgeID, _ := strconv.Atoi(r.FormValue("id"))
		updatedAt := time.Now()
		err := models.UpdateKnowledge(knowledgeID, title, content, rowContent, updatedAt, eyecatchSrc, category)
		if err != nil {
			log.Print(err.Error())
			return
		}
		if err := models.DeleteKnowledgesTagsFromKnowledgeID(knowledgeID); err != nil {
			log.Print(err.Error())
			return
		}
		if r.FormValue("tags") != "" {
			tags := strings.Split(r.FormValue("tags"), ",")
			createdAt := time.Now()
			for _, tag := range tags {
				tagID, _ := strconv.Atoi(tag)
				err := models.PostKnowledgesTags(knowledgeID, tagID, createdAt, updatedAt)
				if err != nil {
					log.Print(err.Error())
					return
				}
			}
		}
		return
	default:
		break
	}
	http.Redirect(w, r, config.AdminKnowledgesPath, http.StatusFound)
}
