package handlers

import (
	"net/http"
	"strconv"

	"code-database/models"
	"code-database/routes"
)

const lenPathAdminPublish = len(routes.AdminPublishPath)

//AdminPublishHandler /admin/publish/に対するハンドラ
func AdminPublishHandler(w http.ResponseWriter, r *http.Request) {
	suffix := r.URL.Path[lenPathAdminPublish:]
	id, _ := strconv.Atoi(suffix)
	if r.Method == "POST" {
		models.SetPublicKnowledge(id)
	} else {
		models.SetPrivateKnowledge(id)
	}
}
