package handlers

import (
	"code-database/config"
	"code-database/models"
	"code-database/structs"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
)

const lenPathTag = len(config.UserTagPath)

//TagHandler /tags/に対するハンドラ
func TagHandler(w http.ResponseWriter, r *http.Request, auth bool) {
	header := newHeader(false)
	if auth {
		header.IsLogin = true
	}
	suffix := r.URL.Path[lenPathTag:]
	if suffix != "" {
		pageNum := 1
		var err error
		query := r.URL.Query()
		if query["page"] != nil {
			if pageNum, err = strconv.Atoi(query.Get("page")); err != nil {
				StatusNotFoundHandler(w, r, auth)
				return
			}
		}
		sortKey := "created_at"
		var currentSort string
		if query["sort"] != nil {
			switch {
			case query.Get("sort") == "create":
				sortKey = "created_at"
				currentSort = "create"
				break
			case query.Get("sort") == "update":
				sortKey = "updated_at"
				currentSort = "update"
				break
			case query.Get("sort") == "like":
				sortKey = "likes"
				currentSort = "like"
				break
			default:
				StatusNotFoundHandler(w, r, auth)
				return
			}
		} else {
			currentSort = "create"
		}
		var filteredTag structs.Tag
		filteredTag.ID, err = strconv.Atoi(suffix)
		if err != nil {
			StatusNotFoundHandler(w, r, auth)
			return
		}
		tagRankingElem, err := models.GetTop10ReferencedTags()
		if err != nil {
			log.Print(err.Error())
			StatusInternalServerError(w, r, auth)
			return
		}
		NumOfKnowledges, err := models.GetNumOfKnowledgesFilteredByTagID(filteredTag.ID)
		if err != nil {
			log.Print(err.Error())
			StatusNotFoundHandler(w, r, auth)
			return
		}
		pageNums := int(math.Ceil(NumOfKnowledges / 20))
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
		var indexElems []structs.IndexElem
		indexElems, filteredTag.Name, err = models.Get20SortedElemFilteredTagID(sortKey, filteredTag.ID, (pageNum-1)*20, 20)
		if err != nil {
			log.Print(err.Error())
			StatusNotFoundHandler(w, r, auth)
			return
		}
		indexPage := structs.UserIndexPage{
			PageNation:  pageNation,
			IndexElems:  indexElems,
			CurrentSort: currentSort,
			TagRanking:  tagRankingElem,
		}
		t := template.Must(template.ParseFiles("template/user_tag.html", "template/_header.html", "template/_footer.html"))
		if err = t.Execute(w, struct {
			Header      structs.Header
			IndexPage   structs.UserIndexPage
			FilteredTag structs.Tag
		}{
			Header:      header,
			IndexPage:   indexPage,
			FilteredTag: filteredTag,
		}); err != nil {
			StatusInternalServerError(w, r, auth)
			return
		}
	} else {
		TagsHandler(w, r, auth)
		return
	}
}
