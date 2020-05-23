package handlers

import (
	"html/template"
	"log"
	"net/http"
)

//PrivacyHandler /privacyに対するハンドラ
func PrivacyHandler(w http.ResponseWriter, r *http.Request, auth bool) {
	header := newHeader(false)
	if auth {
		header.IsLogin = true
	}
	t := template.Must(template.ParseFiles("template/user_privacy.html", "template/_header.html", "template/_footer.html"))
	if err := t.Execute(w, nil); err != nil {
		log.Print(err)
		StatusInternalServerError(w, r, auth)
		return
	}
}
