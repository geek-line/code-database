package handlers

import (
	"html/template"
	"net/http"

	"code-database/structs"
)

// StatusInternalServerError に対するハンドラ
func StatusInternalServerError(w http.ResponseWriter, r *http.Request, auth bool) {
	header := newHeader(false)
	if auth {
		header.IsLogin = true
	}
	t := template.Must(getTemplate("dist/template/500.html", "dist/template/_header.html", "dist/template/_footer.html"))
	t.Execute(w, struct {
		Header structs.Header
	}{
		Header: header,
	})
}
