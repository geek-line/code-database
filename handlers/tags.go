package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
)

const lenPathTags = len("/tags/")

//TagsHandler /tags/に対するハンドラ
func TagsHandler(w http.ResponseWriter, r *http.Request, env map[string]string) {
	session, _ := store.Get(r, "cookie-name")
	header := newHeader(false)
	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		header.IsLogin = true
	}

	suffix := r.URL.Path[lenPathTags:]
	db, err := sql.Open("mysql", env["SQL_ENV"])
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if suffix != "" {
		pageNum := 1
		query := r.URL.Query()
		if query["page"] != nil {
			if pageNum, err = strconv.Atoi(query.Get("page")); err != nil {
				StatusNotFoundHandler(w, r)
				return
			}
		}
		var filteredTag Tag
		filteredTag.ID, err = strconv.Atoi(suffix)
		if err != nil {
			StatusNotFoundHandler(w, r)
			return
		}
		rows, err := db.Query("SELECT id, name FROM tags")
		if err != nil {
			panic(err.Error())
		}
		defer rows.Close()
		var tags []Tag
		for rows.Next() {
			var tag Tag
			err := rows.Scan(&tag.ID, &tag.Name)
			if err != nil {
				panic(err.Error())
			}
			tags = append(tags, tag)
		}

		var indexPage IndexPage
		var knowledgeNums float64
		db.QueryRow("SELECT count(knowledges.id) FROM knowledges INNER JOIN knowledges_tags ON knowledges_tags.knowledge_id = knowledges.id WHERE tag_id = ?", filteredTag.ID).Scan(&knowledgeNums)
		pageNums := int(math.Ceil(knowledgeNums / 20))
		if pageNums < pageNum {
			StatusNotFoundHandler(w, r)
			return
		}
		var pageNationElems = make([]Page, pageNums)
		for i := 0; i < pageNums; i++ {
			pageNationElems[i].PageNum = i + 1
			pageNationElems[i].IsSelect = false
		}
		pageNationElems[pageNum-1].IsSelect = true
		indexPage.PageNation.PageElems = pageNationElems
		indexPage.PageNation.PageNum = pageNum
		indexPage.PageNation.NextPageNum = pageNum + 1
		indexPage.PageNation.PrevPageNum = pageNum - 1

		db.QueryRow("SELECT name FROM tags WHERE id = ?", filteredTag.ID).Scan(&filteredTag.Name)
		rows, err = db.Query("SELECT knowledges.id, title, knowledges.updated_at, likes, eyecatch_src FROM knowledges INNER JOIN knowledges_tags ON knowledges_tags.knowledge_id = knowledges.id WHERE tag_id = ? LIMIT ?, ?", filteredTag.ID, (pageNum-1)*20, 20)
		if err != nil {
			// panic(err.Error())
			StatusNotFoundHandler(w, r)
			log.Println("クエリのエラー")
			return
		}
		defer rows.Close()
		for rows.Next() {
			var indexElem IndexElem
			err := rows.Scan(&indexElem.ID, &indexElem.Title, &indexElem.UpdatedAt, &indexElem.Likes, &indexElem.EyeCatchSrc)
			if err != nil {
				panic(err.Error())
			}
			var selectedTags []Tag
			tagsRows, err := db.Query("SELECT tags.id, tags.name FROM tags INNER JOIN knowledges_tags ON knowledges_tags.tag_id = tags.id WHERE knowledge_id = ?", indexElem.ID)
			if err != nil {
				panic(err.Error())
			}
			defer tagsRows.Close()
			for tagsRows.Next() {
				var selectedTag Tag
				err := tagsRows.Scan(&selectedTag.ID, &selectedTag.Name)
				if err != nil {
					panic(err.Error())
				}
				selectedTags = append(selectedTags, selectedTag)
			}
			indexElem.SelectedTags = selectedTags
			indexPage.IndexElems = append(indexPage.IndexElems, indexElem)
		}
		t := template.Must(template.ParseFiles("template/user_tags.html", "template/_header.html", "template/_footer.html"))
		if err = t.Execute(w, struct {
			Header      Header
			Tags        []Tag
			IndexPage   IndexPage
			FilteredTag Tag
		}{
			Header:      header,
			Tags:        tags,
			IndexPage:   indexPage,
			FilteredTag: filteredTag,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		StatusNotFoundHandler(w, r)
	}
}