package handlers

import (
	"html/template"
	"log"
	"net/http"

	"code-database/config"
	"code-database/models"
	"code-database/structs"
)

func newHeader(isLogin bool) structs.Header {
	return structs.Header{IsLogin: isLogin}
}

//AdminNewHandler /admin/newに対するハンドラ
func AdminNewHandler(w http.ResponseWriter, r *http.Request) {
	tags, err := models.GetAllTags()
	if err != nil {
		log.Print(err.Error())
		return
	}
	eyecatches, err := models.GetAllEyecatches()
	if err != nil {
		log.Print(err.Error())
		return
	}
	categories, err := models.GetAllCategories()
	if err != nil {
		log.Print(err.Error())
		return
	}
	t := template.Must(template.ParseFiles("template/admin_new.html", "template/_header.html"))
	header := newHeader(true)
	if err := t.Execute(w, struct {
		Header     structs.Header
		Tags       []structs.Tag
		Eyecatches []structs.Eyecatch
		Categories []structs.Category
		BuildMode  string
	}{
		Header:     header,
		Tags:       tags,
		Eyecatches: eyecatches,
		Categories: categories,
		BuildMode:  config.BuildMode,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
