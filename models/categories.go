package models

import (
	"code-database/config"
	"code-database/structs"
	"database/sql"
	"time"
)

//GetAllCategories 全てのカテゴリを取得する
func GetAllCategories() ([]structs.Category, error) {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	rows, err := db.Query("SELECT id, name, eyecatch_src, summary FROM categories")
	defer rows.Close()
	var categories []structs.Category
	for rows.Next() {
		var category structs.Category
		err = rows.Scan(&category.ID, &category.Name, &category.EyecatchSrc, &category.Summary)
		categories = append(categories, category)
	}
	return categories, err
}

//GetNumOfCategories 全てのカテゴリの数を取得する
func GetNumOfCategories() (float64, error) {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	var numOfCategories float64
	err = db.QueryRow("SELECT count(id) FROM categories").Scan(&numOfCategories)
	return numOfCategories, err
}

//Get20CategoryElems 記事数の多い上位20カテゴリを取得する
func Get20CategoryElems(startIndex int, length int) ([]structs.CategoryElem, error) {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	rows, err := db.Query("SELECT categories.id, categories.name, categories.eyecatch_src, categories.summary, count(categories.id) AS num FROM categories INNER JOIN knowledges ON knowledges.category = categories.name WHERE is_published = true GROUP BY categories.id ORDER BY num DESC LIMIT ?, ?", startIndex, length)
	defer rows.Close()
	var categoryElems []structs.CategoryElem
	for rows.Next() {
		var categoryElem structs.CategoryElem
		err = rows.Scan(&categoryElem.Category.ID, &categoryElem.Category.Name, &categoryElem.Category.EyecatchSrc, &categoryElem.Category.Summary, &categoryElem.NumOfArticles)
		categoryElems = append(categoryElems, categoryElem)
	}
	return categoryElems, err
}

//PostCategory categoriesを新規作成する
func PostCategory(name string, eyecatchSrc string, summary string, createdAt time.Time, updatedAt time.Time) error {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	rows, err := db.Query("INSERT INTO categories(name, eyecatch_src, summary, created_at, updated_at) VALUES(?, ?, ?, ?, ?)", name, eyecatchSrc, summary, createdAt, updatedAt)
	defer rows.Close()
	return err
}

//UpdateCategory 指定されたidのcategoriesを変更する
func UpdateCategory(id int, name string, eyecatchSrc string, summary string, updatedAt time.Time) error {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	rows, err := db.Query("UPDATE categories SET name = ?, eyecatch_src = ?, summary = ?, updated_at = ? WHERE id = ?", name, eyecatchSrc, summary, updatedAt, id)
	defer rows.Close()
	return err
}

//DeleteCategory 指定されたidのtagを削除する
func DeleteCategory(id int) error {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	rows, err := db.Query("DELETE FROM categories WHERE id = ?", id)
	defer rows.Close()
	return err
}
