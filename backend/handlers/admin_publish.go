package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"code-database/config"
	"code-database/models"
	"code-database/xml_update/helpers"
)

const lenPathAdminPublish = len(config.AdminPublishPath)

// AdminPublishHandler /admin/publish/に対するハンドラ
func AdminPublishHandler(w http.ResponseWriter, r *http.Request) {
	suffix := r.URL.Path[lenPathAdminPublish:]
	id, _ := strconv.Atoi(suffix)
	var message string
	if r.Method == "POST" {
		if err := models.SetPublicKnowledge(id); err != nil {
			AdminKnowledgesHandler(w, r)
		}
		message = fmt.Sprintf("{\"text\":\"https://code-database.com/knowledges/%d が公開されました!\nサイトマップにも反映済みです。\"}", id)

	} else {
		if err := models.SetPrivateKnowledge(id); err != nil {
			AdminKnowledgesHandler(w, r)
		}
		message = fmt.Sprintf("{\"text\":\"https://code-database.com/knowledges/%d が非公開になりました!\nサイトマップにも反映済みです。\"}", id)

	}
	if config.BuildMode == "prod" {
		publishedKnowledges, err := models.GetAllPublishedKnowledges()
		if err != nil {
			AdminKnowledgesHandler(w, r)
		}
		urlSet := helpers.MakeKnowledgesXMLSitemap(publishedKnowledges)
		if err = urlSet.UpdateXMLSitemap("knowledges.xml"); err != nil {
			AdminKnowledgesHandler(w, r)
		}
		resp, err := http.Post("https://hooks.slack.com/services/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", "applicotion/json", strings.NewReader(message))
		if err != nil {
			AdminKnowledgesHandler(w, r)
		}
		defer resp.Body.Close()
	}
}
