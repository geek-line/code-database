package handlers

import (
	"code-database/models"
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
	//ここでおすすめ記事の番号を指定
	arg := []string{"93", "96", "111"}
	recomendedKnowledges, err := models.GetRecomendedElems(arg)
	if err != nil {
		log.Print(err.Error())
		StatusInternalServerError(w, r, auth)
		return
	}
	likedKnowledges, err := models.Get20SortedElems("likes", 0, 3)
	if err != nil {
		log.Print(err.Error())
		StatusInternalServerError(w, r, auth)
		return
	}
	recentKnowledges, err := models.Get20SortedElems("updated_at", 0, 3)
	if err != nil {
		log.Print(err.Error())
		StatusInternalServerError(w, r, auth)
		return
	}
	userTopPage := structs.UserTopPage{
		RecomendedKnowledges: recomendedKnowledges,
		LikedKnowledges:      likedKnowledges,
		RecentKnowledges:     recentKnowledges,
	}
	t := template.Must(template.ParseFiles("template/user_top.html", "template/_header.html", "template/_footer.html"))
	if err := t.Execute(w, struct {
		Header      structs.Header
		UserTopPage structs.UserTopPage
	}{
		Header:      header,
		UserTopPage: userTopPage,
	}); err != nil {
		log.Print(err)
		StatusInternalServerError(w, r, auth)
		return
	}
}
