package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"code-database/config"
	"code-database/models"
)

const lenPathAdminPublish = len(config.AdminPublishPath)

//AdminPublishHandler /admin/publish/に対するハンドラ
func AdminPublishHandler(w http.ResponseWriter, r *http.Request) {
	suffix := r.URL.Path[lenPathAdminPublish:]
	id, _ := strconv.Atoi(suffix)
	if r.Method == "POST" {
		if err := models.SetPublicKnowledge(id); err != nil {
			AdminKnowledgesHandler(w, r)
		}
		message := fmt.Sprintf("{\"text\":\"https://code-database.com/knowledges/%d が公開されました\"}", id)
		resp, err := http.Post("https://hooks.slack.com/services/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", "applicotion/json", strings.NewReader(message))
		if err != nil {
			AdminKnowledgesHandler(w, r)
		}
		defer resp.Body.Close()
	} else {
		if err := models.SetPrivateKnowledge(id); err != nil {
			AdminKnowledgesHandler(w, r)
		}
		message := fmt.Sprintf("{\"text\":\"https://code-database.com/knowledges/%d が非公開になりました\"}", id)
		resp, err := http.Post("https://hooks.slack.com/services/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", "applicotion/json", strings.NewReader(message))
		if err != nil {
			AdminKnowledgesHandler(w, r)
		}
		defer resp.Body.Close()
	}
}
