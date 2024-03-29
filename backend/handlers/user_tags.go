package handlers

import (
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"

	"code-database/models"
	"code-database/structs"
)

//TagsHandler /tagsに対するハンドラ
func TagsHandler(w http.ResponseWriter, r *http.Request, auth bool) {
	header := newHeader(false)
	if auth {
		header.IsLogin = true
	}
	pageNum := 1
	var err error
	query := r.URL.Query()
	if query["page"] != nil {
		if pageNum, err = strconv.Atoi(query.Get("page")); err != nil {
			StatusNotFoundHandler(w, r, auth)
			return
		}
	}
	NumOfTags, err := models.GetNumOfTags()
	if err != nil {
		log.Print(err.Error())
		StatusNotFoundHandler(w, r, auth)
		return
	}
	pageNums := int(math.Ceil(NumOfTags / 100))
	if pageNums < pageNum {
		StatusNotFoundHandler(w, r, auth)
		return
	}
	var pageNationElems = make([]structs.Page, pageNums)
	for i := 0; i < pageNums; i++ {
		pageNationElems[i].PageNum = i + 1
		pageNationElems[i].IsSelect = false
	}
	pageNationElems[pageNum-1].IsSelect = true
	var pageNation structs.PageNation
	pageNation.PageElems = pageNationElems
	pageNation.PageNum = pageNum
	pageNation.NextPageNum = pageNum + 1
	pageNation.PrevPageNum = pageNum - 1
	tagRanking, err := models.GetTop10ReferencedTags()
	if err != nil {
		log.Print(err.Error())
		StatusInternalServerError(w, r, auth)
		return
	}
	tagElems, err := models.Get50TagElems((pageNum-1)*100, 100)
	if err != nil {
		log.Print(err.Error())
		StatusInternalServerError(w, r, auth)
		return
	}
	userTagsPage := structs.UserTagsPage{
		Tags:       tagElems,
		TagRanking: tagRanking,
		PageNation: pageNation,
	}
	t := template.Must(getTemplate("dist/template/user_tags.html", "dist/template/_header.html", "dist/template/_footer.html"))
	if err = t.Execute(w, struct {
		Header   structs.Header
		TagsPage structs.UserTagsPage
	}{
		Header:   header,
		TagsPage: userTagsPage,
	}); err != nil {
		log.Print(err)
		StatusInternalServerError(w, r, auth)
		return
	}
}
