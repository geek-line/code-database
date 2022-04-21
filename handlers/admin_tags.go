package handlers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"code-database/config"
	"code-database/models"
	"code-database/structs"
)

//AdminTagsHandler /admin/tagsに対するハンドラ
func AdminTagsHandler(w http.ResponseWriter, r *http.Request) {
	header := newHeader(true)
	switch {
	case r.Method == "GET":
		tags, err := models.GetAllTags()
		if err != nil {
			log.Print(err.Error())
			return
		}
		t := template.Must(template.ParseFiles("template/admin_tags.html", "template/_header.html"))
		if err := t.Execute(w, struct {
			Header    structs.Header
			Tags      []structs.Tag
			BuildMode string
		}{
			Header:    header,
			Tags:      tags,
			BuildMode: config.BuildMode,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case r.Method == "POST":
		name := r.FormValue("name")
		createdAt := time.Now()
		updatedAt := time.Now()
		err := models.PostTag(name, createdAt, updatedAt)
		if err != nil {
			log.Print(err.Error())
			return
		}
		http.Redirect(w, r, config.AdminTagsPath, http.StatusFound)
	case r.Method == "PUT":
		id, _ := strconv.Atoi(r.FormValue("id"))
		name := r.FormValue("name")
		updatedAt := time.Now()
		err := models.UpdateTag(id, name, updatedAt)
		if err != nil {
			log.Print(err.Error())
			return
		}
	case r.Method == "DELETE":
		id, _ := strconv.Atoi(r.FormValue("id"))
		err := models.DeleteTag(id)
		if err != nil {
			log.Print(err.Error())
			return
		}
		err = models.DeleteKnowledgesTagsFromKnowledgeIDFromTagID(id)
		if err != nil {
			log.Print(err.Error())
			return
		}
	}
}
