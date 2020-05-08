package handlers

import (
	"log"
	"net/http"
	"strconv"

	"code-database/config"
	"code-database/models"
)

const lenPathDelete = len(config.AdminDeletePath)

//AdminDeleteHandler admin/deleteに対するハンドラ
func AdminDeleteHandler(w http.ResponseWriter, r *http.Request) {

	suffix := r.URL.Path[lenPathDelete:]
	var id int
	id, _ = strconv.Atoi(suffix)
	err := models.DeleteKnowledge(id)
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = models.DeleteKnowledgesTagsFromKnowledgeID(id)
	if err != nil {
		log.Print(err.Error())
		return
	}
	http.Redirect(w, r, config.AdminKnowledgesPath, http.StatusFound)
}
