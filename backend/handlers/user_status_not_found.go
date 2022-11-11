package handlers

import (
	"html/template"
	"net/http"

	"code-database/structs"
)

// StatusNotFoundHandler に対するハンドラ
func StatusNotFoundHandler(w http.ResponseWriter, r *http.Request, auth bool) {
	header := newHeader(false)
	if auth {
		header.IsLogin = true
	}
	t := template.Must(getTemplate("dist/template/404.html", "dist/template/_header.html", "dist/template/_footer.html"))
	if err := t.Execute(w, struct {
		Header structs.Header
	}{
		Header: header,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
