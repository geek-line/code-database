package handlers

import (
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"code-database/models"
	"code-database/structs"
)

// SearchHandler /searchに対するハンドラ
func SearchHandler(w http.ResponseWriter, r *http.Request, auth bool) {
	header := newHeader(false)
	if auth {
		header.IsLogin = true
	}
	pageNum := 1
	isHit := true
	var indexPage structs.UserIndexPage
	var err error
	query := r.URL.Query()
	if query["page"] != nil {
		if pageNum, err = strconv.Atoi(query.Get("page")); err != nil {
			StatusNotFoundHandler(w, r, auth)
			return
		}
	}
	var queryKeys string
	var currentQuery template.URL
	if query["q"] != nil {
		queryKeys = query.Get("q")
		currentQuery = template.URL(strings.Split(r.URL.RawQuery, "&")[0])
	} else {
		StatusNotFoundHandler(w, r, auth)
		return
	}
	sortKey := "created_at"
	var currentSort string
	if query["sort"] != nil {
		switch {
		case query.Get("sort") == "create":
			sortKey = "created_at"
			currentSort = "create"
		case query.Get("sort") == "update":
			sortKey = "updated_at"
			currentSort = "update"
		case query.Get("sort") == "like":
			sortKey = "likes"
			currentSort = "like"
		default:
			StatusNotFoundHandler(w, r, auth)
			return
		}
	} else {
		currentSort = "create"
	}
	tagRankingElem, err := models.GetTop10ReferencedTags()
	if err != nil {
		log.Print(err.Error())
		StatusInternalServerError(w, r, auth)
		return
	}
	NumOfKnowledges, err := models.GetNumOfKnowledgesHitByQuery(queryKeys)
	if err != nil {
		log.Print(err.Error())
		StatusNotFoundHandler(w, r, auth)
		return
	}
	if NumOfKnowledges == 0 {
		isHit = false
		indexPage = structs.UserIndexPage{
			TagRanking: tagRankingElem,
		}
	} else {
		isHit = true
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
		indexElems, err = models.Get20SortedElemHitByQuery(sortKey, queryKeys, (pageNum-1)*20, 20)
		if err != nil {
			log.Print(err.Error())
			StatusNotFoundHandler(w, r, auth)
			return
		}
		indexPage = structs.UserIndexPage{
			PageNation:  pageNation,
			IndexElems:  indexElems,
			CurrentSort: currentSort,
			TagRanking:  tagRankingElem,
		}
	}
	t := template.Must(getTemplate("dist/template/user_search.html", "dist/template/_header.html", "dist/template/_footer.html"))
	if err = t.Execute(w, struct {
		Header       structs.Header
		IndexPage    structs.UserIndexPage
		IsHit        bool
		QueryKeys    string
		CurrentQuery template.URL
	}{
		Header:       header,
		IndexPage:    indexPage,
		IsHit:        isHit,
		QueryKeys:    queryKeys,
		CurrentQuery: currentQuery,
	}); err != nil {
		StatusInternalServerError(w, r, auth)
		return
	}
}
