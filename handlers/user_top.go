package handlers

import (
	"code-database/structs"
	"html/template"
	"log"
	"net/http"
)

//TopHandler トップページに関するハンドラ
func TopHandler(w http.ResponseWriter, r *http.Request, auth bool) {
	header := newHeader(false)
	if auth {
		header.IsLogin = true
	}
	t := template.Must(template.ParseFiles("template/user_top.html", "template/_header.html", "template/_footer.html"))
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
