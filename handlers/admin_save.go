package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"../routes"
)

//AdminSaveHandler /admin/saveに対するハンドラ
func AdminSaveHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	title := r.FormValue("title")
	content := r.FormValue("content")
	eyecatchSrc := r.FormValue("eyecatch_src")
	switch {
	case r.Method == "POST":
		createdAt := time.Now()
		updatedAt := time.Now()
		stmtInsert, err := db.Prepare("INSERT INTO knowledges(title, content, created_at, updated_at, eyecatch_src) VALUES(?, ?, ?, ?, ?)")
		if err != nil {
			log.Print(err.Error())
			return
		}
		defer stmtInsert.Close()
		result, err := stmtInsert.Exec(title, content, createdAt, updatedAt, eyecatchSrc)
		if err != nil {
			log.Print(err.Error())
			return
		}
		knowledgeID, err := result.LastInsertId()
		if err != nil {
			log.Print(err.Error())
			return
		}
		if r.FormValue("tags") != "" {
			tags := strings.Split(r.FormValue("tags"), ",")
			for _, tag := range tags {
				tagID, _ := strconv.Atoi(tag)
				rows, err := db.Query("INSERT INTO knowledges_tags(knowledge_id, tag_id, created_at, updated_at) VALUES(?, ?, ?, ?)", knowledgeID, tagID, createdAt, updatedAt)
				if err != nil {
					log.Print(err.Error())
					return
				}
				defer rows.Close()
			}
		}
	case r.Method == "PUT":
		knowledgeID, _ := strconv.Atoi(r.FormValue("id"))
		updatedAt := time.Now()
		rows, err := db.Query("UPDATE knowledges SET title = ?, content = ?, updated_at = ?, eyecatch_src = ? WHERE id = ?", title, content, updatedAt, eyecatchSrc, knowledgeID)
		if err != nil {
			log.Print(err.Error())
			return
		}
		defer rows.Close()
		rows, err = db.Query("DELETE FROM knowledges_tags WHERE knowledge_id = ?", knowledgeID)
		if err != nil {
			log.Print(err.Error())
			return
		}
		defer rows.Close()
		if r.FormValue("tags") != "" {
			tags := strings.Split(r.FormValue("tags"), ",")
			createdAt := time.Now()
			for _, tag := range tags {
				tagID, _ := strconv.Atoi(tag)
				rows, err := db.Query("INSERT INTO knowledges_tags(knowledge_id, tag_id, created_at, updated_at) VALUES(?, ?, ?, ?)", knowledgeID, tagID, createdAt, updatedAt)
				if err != nil {
					log.Print(err.Error())
					return
				}
				defer rows.Close()
			}
		}
		return
	default:
		break
	}
	http.Redirect(w, r, routes.AdminKnowledgesPath, http.StatusFound)
}