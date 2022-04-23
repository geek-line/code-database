package handlers

import (
	"code-database/structs"
	"html/template"
	"log"
	"net/http"
)

//AboutHandler /aboutに対するハンドラ
func AboutHandler(w http.ResponseWriter, r *http.Request, auth bool) {
	header := newHeader(false)
	if auth {
		header.IsLogin = true
	}
	t := template.Must(getTemplate("dist/template/user_about.html", "dist/template/_header.html", "dist/template/_footer.html"))
	if err := t.Execute(w, struct {
		Header structs.Header
	}{
		Header: header,
	}); err != nil {
		log.Print(err)
		StatusInternalServerError(w, r, auth)
		return
	}
}
