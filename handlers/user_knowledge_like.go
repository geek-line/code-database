package handlers

import (
	"log"
	"net/http"
	"strconv"

	"code-database/models"
)

// KnowledgeLikeHandler /knowledges/likeに対するハンドラ
func KnowledgeLikeHandler(w http.ResponseWriter, r *http.Request, auth bool) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Print(err.Error())
		StatusInternalServerError(w, r, auth)
		return
	}
	if r.Method == "POST" {
		if err := models.IncrementLikes(id); err != nil {
			log.Print(err.Error())
			StatusInternalServerError(w, r, auth)
			return
		}
	} else {
		if err := models.DecrementLikes(id); err != nil {
			log.Print(err.Error())
			StatusInternalServerError(w, r, auth)
			return
		}
	}
}
