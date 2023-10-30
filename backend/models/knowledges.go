package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"code-database/config"
	"code-database/structs"
)

// DeleteKnowledge 指定されたidのknowledgeを削除する
func DeleteKnowledge(id int) error {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	rows, err := db.Query("DELETE FROM knowledges WHERE id = ?", id)
	defer rows.Close()
	return err
}

// GetKnowledge 指定されたidのknowledgeを取得する
func GetKnowledge(id int) (structs.Knowledge, error) {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	var knowledge structs.Knowledge
	err = db.QueryRow("SELECT id, title, content, updated_at, likes, eyecatch_src, category FROM knowledges WHERE id = ?", id).Scan(&knowledge.ID, &knowledge.Title, &knowledge.Content, &knowledge.UpdatedAt, &knowledge.Likes, &knowledge.EyecatchSrc, &knowledge.Category)
	return knowledge, err
}

// GetKnowledgePublished 指定されたidで公開されているknowledgeを取得する
func GetKnowledgePublished(id int) (structs.Knowledge, error) {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	var knowledge structs.Knowledge
	err = db.QueryRow("SELECT id, title, content, updated_at, likes, eyecatch_src, category FROM knowledges WHERE id = ? AND is_published = true", id).Scan(&knowledge.ID, &knowledge.Title, &knowledge.Content, &knowledge.UpdatedAt, &knowledge.Likes, &knowledge.EyecatchSrc, &knowledge.Category)
	return knowledge, err
}

// GetAllKnowledges 全てのknowledgeを取得する
func GetAllKnowledges() ([]structs.Knowledge, error) {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	var knowledges []structs.Knowledge
	rows, err := db.Query("SELECT id, title, created_at, updated_at, is_published FROM knowledges ORDER BY id DESC")
	defer rows.Close()
	for rows.Next() {
		var knowledge structs.Knowledge
		err = rows.Scan(&knowledge.ID, &knowledge.Title, &knowledge.CreatedAt, &knowledge.UpdatedAt, &knowledge.IsPublished)
		knowledges = append(knowledges, knowledge)
	}
	return knowledges, err
}

// GetAllPublishedKnowledges 公開されているknowledgeを全て取得する
func GetAllPublishedKnowledges() ([]structs.Knowledge, error) {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	var knowledges []structs.Knowledge
	rows, err := db.Query("SELECT id, title, created_at, updated_at, is_published FROM knowledges WHERE is_published = true ORDER BY id")
	defer rows.Close()
	for rows.Next() {
		var knowledge structs.Knowledge
		err = rows.Scan(&knowledge.ID, &knowledge.Title, &knowledge.CreatedAt, &knowledge.UpdatedAt, &knowledge.IsPublished)
		knowledges = append(knowledges, knowledge)
	}
	return knowledges, err
}

// PostKnowledge knowledgeを新規作成して作成したknowledgeのIDを取得する
func PostKnowledge(title string, content string, rowContent string, createdAt time.Time, updatedAt time.Time, eyecatchSrc string, category string) (int64, error) {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	stmtInsert, err := db.Prepare("INSERT INTO knowledges(title, content, row_content, created_at, updated_at, eyecatch_src, category) VALUES(?, ?, ?, ?, ?, ?, ?)")
	defer stmtInsert.Close()
	result, err := stmtInsert.Exec(title, content, rowContent, createdAt, updatedAt, eyecatchSrc, category)
	knowledgeID, err := result.LastInsertId()
	return knowledgeID, err
}

// UpdateKnowledge 指定したidのknowledgeを更新する
func UpdateKnowledge(knowledgeID int, title string, content string, rowContent string, updatedAt time.Time, eyecatchSrc string, category string) error {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	rows, err := db.Query("UPDATE knowledges SET title = ?, content = ?, row_content = ?, updated_at = ?, eyecatch_src = ?, category = ? WHERE id = ?", title, content, rowContent, updatedAt, eyecatchSrc, category, knowledgeID)
	defer rows.Close()
	return err
}

// IncrementLikes 指定されたidのlikesを増やす
func IncrementLikes(id int) error {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	rows, err := db.Query("UPDATE knowledges SET likes = likes + 1 WHERE id = ?", id)
	defer rows.Close()
	return err
}

// DecrementLikes 指定されたidのlikesを減らす
func DecrementLikes(id int) error {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	rows, err := db.Query("UPDATE knowledges SET likes = likes - 1 WHERE id = ?", id)
	defer rows.Close()
	return err
}

// GetNumOfKnowledges knowledgeの数を取得する
func GetNumOfKnowledges() (float64, error) {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	var numOfKnowledges float64
	err = db.QueryRow("SELECT count(id) FROM knowledges WHERE is_published = true").Scan(&numOfKnowledges)
	return numOfKnowledges, err
}

// GetNumOfKnowledgesFilteredByTagID 指定されたtag_idに該当するknowledgeの数を取得する
func GetNumOfKnowledgesFilteredByTagID(id int) (float64, error) {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	var numOfKnowledges float64
	err = db.QueryRow("SELECT count(knowledges.id) FROM knowledges INNER JOIN knowledges_tags ON knowledges_tags.knowledge_id = knowledges.id WHERE tag_id = ? AND is_published = true", id).Scan(&numOfKnowledges)
	return numOfKnowledges, err
}

// GetNumOfCategoriesFilteredByCategoryID 指定されたidのカテゴリに含まれる記事の数を取得する
func GetNumOfCategoriesFilteredByCategoryID(id int) (float64, error) {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	var numOfKnowledges float64
	err = db.QueryRow("SELECT count(knowledges.id) FROM knowledges INNER JOIN categories ON categories.name = knowledges.category WHERE categories.id = ?", id).Scan(&numOfKnowledges)
	return numOfKnowledges, err
}

// GetNumOfKnowledgesHitByQuery 指定されたクエリの配列がコンテンツに含まれるナレッジの数を返す
func GetNumOfKnowledgesHitByQuery(queryKeys string) (float64, error) {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	conditionText := "%" + strings.ReplaceAll(queryKeys, " ", "%") + "%"
	var numOfKnowledges float64
	_ = db.QueryRow("SELECT count(id) FROM knowledges WHERE row_content LIKE ? AND is_published = true", conditionText).Scan(&numOfKnowledges)
	return numOfKnowledges, err
}

// Get20SortedElems 指定のsortKeyでソートされた20のknowledgeの要素を取得する
func Get20SortedElems(sortKey string, startIndex int, length int) ([]structs.IndexElem, error) {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	qtext := fmt.Sprintf("SELECT id, title, updated_at, likes, eyecatch_src FROM knowledges WHERE is_published = true ORDER BY %s DESC LIMIT ?, ?", sortKey)
	rows, err := db.Query(qtext, startIndex, length)
	defer rows.Close()
	var indexElems []structs.IndexElem
	for rows.Next() {
		var indexElem structs.IndexElem
		err = rows.Scan(&indexElem.Knowledge.ID, &indexElem.Knowledge.Title, &indexElem.Knowledge.UpdatedAt, &indexElem.Knowledge.Likes, &indexElem.Knowledge.EyecatchSrc)
		var selectedTags []structs.Tag
		var tagsRows *sql.Rows
		tagsRows, err = db.Query("SELECT tag_id FROM knowledges_tags WHERE knowledge_id = ?", indexElem.Knowledge.ID)
		defer tagsRows.Close()
		for tagsRows.Next() {
			var selectedTag structs.Tag
			err = tagsRows.Scan(&selectedTag.ID)
			db.QueryRow("SELECT name FROM tags WHERE id = ?", selectedTag.ID).Scan(&selectedTag.Name)
			selectedTags = append(selectedTags, selectedTag)
		}
		indexElem.SelectedTags = selectedTags
		indexElems = append(indexElems, indexElem)
	}
	return indexElems, err
}

// GetRecomendedElems 任意の配列の要素の該当する記事を恣意的に取得する関数
func GetRecomendedElems(indexes []string) ([]structs.IndexElem, error) {
	intext := strings.Join(indexes, ",")
	qtext := fmt.Sprintf("SELECT id, title, updated_at, likes, eyecatch_src FROM knowledges WHERE is_published = true AND id IN(%s)", intext)
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	rows, err := db.Query(qtext)
	defer rows.Close()
	var indexElems []structs.IndexElem
	for rows.Next() {
		var indexElem structs.IndexElem
		err = rows.Scan(&indexElem.Knowledge.ID, &indexElem.Knowledge.Title, &indexElem.Knowledge.UpdatedAt, &indexElem.Knowledge.Likes, &indexElem.Knowledge.EyecatchSrc)
		var selectedTags []structs.Tag
		var tagsRows *sql.Rows
		tagsRows, err = db.Query("SELECT tag_id FROM knowledges_tags WHERE knowledge_id = ?", indexElem.Knowledge.ID)
		defer tagsRows.Close()
		for tagsRows.Next() {
			var selectedTag structs.Tag
			err = tagsRows.Scan(&selectedTag.ID)
			db.QueryRow("SELECT name FROM tags WHERE id = ?", selectedTag.ID).Scan(&selectedTag.Name)
			selectedTags = append(selectedTags, selectedTag)
		}
		indexElem.SelectedTags = selectedTags
		indexElems = append(indexElems, indexElem)
	}
	return indexElems, err
}

// Get20SortedElemFilteredTagID 指定のsortKeyでソートされ、指定のtagIdでフィルターされた20のknowledgeの要素を取得する
func Get20SortedElemFilteredTagID(sortKey string, tagID int, startIndex int, length int) ([]structs.IndexElem, string, error) {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	var tagName string
	if err = db.QueryRow("SELECT name FROM tags WHERE id = ?", tagID).Scan(&tagName); err != nil {
		return nil, "", err
	}
	qtext := fmt.Sprintf("SELECT knowledges.id, title, knowledges.updated_at, likes, eyecatch_src FROM knowledges INNER JOIN knowledges_tags ON knowledges_tags.knowledge_id = knowledges.id WHERE tag_id = ? AND is_published = true ORDER BY knowledges.%s DESC LIMIT ?, ?", sortKey)
	rows, err := db.Query(qtext, tagID, startIndex, length)
	defer rows.Close()
	var indexElems []structs.IndexElem
	for rows.Next() {
		var indexElem structs.IndexElem
		err = rows.Scan(&indexElem.Knowledge.ID, &indexElem.Knowledge.Title, &indexElem.Knowledge.UpdatedAt, &indexElem.Knowledge.Likes, &indexElem.Knowledge.EyecatchSrc)
		var selectedTags []structs.Tag
		var tagsRows *sql.Rows
		tagsRows, err = db.Query("SELECT tags.id, tags.name FROM tags INNER JOIN knowledges_tags ON knowledges_tags.tag_id = tags.id WHERE knowledge_id = ?", indexElem.Knowledge.ID)
		defer tagsRows.Close()
		for tagsRows.Next() {
			var selectedTag structs.Tag
			err = tagsRows.Scan(&selectedTag.ID, &selectedTag.Name)
			selectedTags = append(selectedTags, selectedTag)
		}
		indexElem.SelectedTags = selectedTags
		indexElems = append(indexElems, indexElem)
	}
	return indexElems, tagName, err
}

// Get20SortedElemFilteredByCategoryID 指定のsortKeyでソートされ、指定のcategoryIdでフィルターされた20のknowledgeの要素を取得する
func Get20SortedElemFilteredByCategoryID(sortKey string, categoryID int, startIndex int, length int) ([]structs.IndexElem, structs.Category, error) {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	category := structs.Category{
		ID: categoryID,
	}
	if err = db.QueryRow("SELECT name, eyecatch_src FROM categories WHERE id = ?", categoryID).Scan(&category.Name, &category.EyecatchSrc); err != nil {
		return nil, category, err
	}
	qtext := fmt.Sprintf("SELECT knowledges.id, title, knowledges.updated_at, likes, knowledges.eyecatch_src FROM knowledges INNER JOIN categories ON categories.name = knowledges.category WHERE categories.id = ? AND is_published = true ORDER BY knowledges.%s DESC LIMIT ?, ?", sortKey)
	rows, err := db.Query(qtext, categoryID, startIndex, length)
	defer rows.Close()
	var indexElems []structs.IndexElem
	for rows.Next() {
		var indexElem structs.IndexElem
		err = rows.Scan(&indexElem.Knowledge.ID, &indexElem.Knowledge.Title, &indexElem.Knowledge.UpdatedAt, &indexElem.Knowledge.Likes, &indexElem.Knowledge.EyecatchSrc)
		var selectedTags []structs.Tag
		var tagsRows *sql.Rows
		tagsRows, err = db.Query("SELECT tags.id, tags.name FROM tags INNER JOIN knowledges_tags ON knowledges_tags.tag_id = tags.id WHERE knowledge_id = ?", indexElem.Knowledge.ID)
		defer tagsRows.Close()
		for tagsRows.Next() {
			var selectedTag structs.Tag
			err = tagsRows.Scan(&selectedTag.ID, &selectedTag.Name)
			selectedTags = append(selectedTags, selectedTag)
		}
		indexElem.SelectedTags = selectedTags
		indexElems = append(indexElems, indexElem)
	}
	return indexElems, category, err
}

// Get20SortedElemHitByQuery 指定されたクエリを含むコンテンツにヒットしたナレッジのなかで上位20を返す
func Get20SortedElemHitByQuery(sortKey string, queryKeys string, startIndex int, length int) ([]structs.IndexElem, error) {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	qtext := fmt.Sprintf("SELECT id, title, updated_at, likes, eyecatch_src FROM knowledges WHERE row_content LIKE ? AND is_published = true ORDER BY %s DESC LIMIT ?, ?", sortKey)
	conditionText := "%" + strings.ReplaceAll(queryKeys, " ", "%") + "%"
	rows, _ := db.Query(qtext, conditionText, startIndex, length)
	defer rows.Close()
	var indexElems []structs.IndexElem
	if rows == nil {
		return indexElems, err
	}
	for rows.Next() {
		var indexElem structs.IndexElem
		err = rows.Scan(&indexElem.Knowledge.ID, &indexElem.Knowledge.Title, &indexElem.Knowledge.UpdatedAt, &indexElem.Knowledge.Likes, &indexElem.Knowledge.EyecatchSrc)
		var selectedTags []structs.Tag
		var tagsRows *sql.Rows
		tagsRows, err = db.Query("SELECT tags.id, tags.name FROM tags INNER JOIN knowledges_tags ON knowledges_tags.tag_id = tags.id WHERE knowledge_id = ?", indexElem.Knowledge.ID)
		defer tagsRows.Close()
		for tagsRows.Next() {
			var selectedTag structs.Tag
			err = tagsRows.Scan(&selectedTag.ID, &selectedTag.Name)
			selectedTags = append(selectedTags, selectedTag)
		}
		indexElem.SelectedTags = selectedTags
		indexElems = append(indexElems, indexElem)
	}
	return indexElems, err
}

// SetPublicKnowledge 指定されたidのナレッジを公開する
func SetPublicKnowledge(id int) error {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	rows, err := db.Query("UPDATE knowledges SET is_published = true WHERE id = ?", id)
	defer rows.Close()
	return err
}

// SetPrivateKnowledge 指定されたidのナレッジを非公開にする
func SetPrivateKnowledge(id int) error {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	rows, err := db.Query("UPDATE knowledges SET is_published = false WHERE id = ?", id)
	defer rows.Close()
	return err
}
