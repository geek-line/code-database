package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"code-database/config"
	"code-database/models"
	"code-database/structs"

	"github.com/gorilla/csrf"
)

const lenPathKnowledge = len(config.UserKnowledgePath)

// KnowledgeHandler /knowledgesに対するハンドラ
func KnowledgeHandler(w http.ResponseWriter, r *http.Request, auth bool) {
	header := newHeader(false)
	if auth {
		header.IsLogin = true
	}
	suffix := r.URL.Path[lenPathKnowledge:]
	if suffix != "" {
		var userDetailPage structs.UserDetailPage
		var id int
		var err error
		if id, err = strconv.Atoi(suffix); err != nil {
			log.Print(err.Error())
			StatusNotFoundHandler(w, r, auth)
			return
		}
		selectedCategory, err := models.GetCategoryFromKnowledgeID(id)
		if err != nil {
			log.Print(err.Error())
			StatusNotFoundHandler(w, r, auth)
			return
		}
		userDetailPage.Knowledge, err = models.GetKnowledgePublished(id)
		switch {
		case err == sql.ErrNoRows:
			log.Println("レコードが存在しません")
			StatusNotFoundHandler(w, r, auth)
			return
		case err != nil:
			log.Print(err.Error())
			StatusInternalServerError(w, r, auth)
			return
		default:
			userDetailPage.SelectedTags, err = models.GetTagFromKnowledgeID(userDetailPage.Knowledge.ID)
			if err != nil {
				log.Print(err.Error())
				StatusInternalServerError(w, r, auth)
				return
			}
			funcMap := template.FuncMap{
				"safehtml": func(text string) template.HTML { return template.HTML(text) },
			}
			t := template.Must(getTemplateWithFuncs(funcMap, "dist/template/user_details.html", "dist/template/_header.html", "dist/template/_footer.html"))
			if err := t.Execute(w, struct {
				Header           structs.Header
				SelectedCategory structs.Category
				DetailPage       structs.UserDetailPage
				CsrfTag          string
				CsrfToken        string
			}{
				Header:           header,
				SelectedCategory: selectedCategory,
				DetailPage:       userDetailPage,
				CsrfTag:          "csrfField",
				CsrfToken:        csrf.Token(r),
			}); err != nil {
				log.Print(err.Error())
				StatusInternalServerError(w, r, auth)
				return
			}
		}
	} else {
		KnowledgesHandler(w, r, auth)
		return
	}
}
