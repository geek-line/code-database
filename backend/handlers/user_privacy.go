package handlers

import (
	"code-database/structs"
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
	t := template.Must(getTemplate("dist/template/user_privacy.html", "dist/template/_header.html", "dist/template/_footer.html"))
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
