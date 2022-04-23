package handlers

import (
	"code-database/config"
	"code-database/models"
	"code-database/structs"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

//AdminCategoriesHandler /admin/categories/に対するハンドラ
func AdminCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	header := newHeader(true)
	switch {
	case r.Method == "GET":
		categories, err := models.GetAllCategories()
		if err != nil {
			log.Print(err.Error())
		}
		eyecatches, err := models.GetAllEyecatches()
		if err != nil {
			log.Print(err.Error())
			return
		}
		t := template.Must(getTemplate("dist/template/admin_categories.html", "dist/template/_header.html"))
		if err := t.Execute(w, struct {
			Header     structs.Header
			Categories []structs.Category
			Eyecatches []structs.Eyecatch
		}{
			Header:     header,
			Categories: categories,
			Eyecatches: eyecatches,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case r.Method == "POST":
		name := r.FormValue("name")
		eyecatchSrc := r.FormValue("eyecatch_src")
		summary := r.FormValue("summary")
		createdAt := time.Now()
		updatedAt := time.Now()
		err := models.PostCategory(name, eyecatchSrc, summary, createdAt, updatedAt)
		if err != nil {
			log.Print(err.Error())
			return
		}
		http.Redirect(w, r, config.AdminCategoriesPath, http.StatusFound)
	case r.Method == "PUT":
		id, _ := strconv.Atoi(r.FormValue("id"))
		name := r.FormValue("name")
		eyecatchSrc := r.FormValue("eyecatch_src")
		summary := r.FormValue("summary")
		updatedAt := time.Now()
		err := models.UpdateCategory(id, name, eyecatchSrc, summary, updatedAt)
		if err != nil {
			log.Print(err.Error())
			return
		}
	case r.Method == "DELETE":
		id, _ := strconv.Atoi(r.FormValue("id"))
		err := models.DeleteCategory(id)
		if err != nil {
			log.Print(err.Error())
			return
		}
	}
}
