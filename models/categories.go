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
	rows, err := db.Query("SELECT id, name, eyecatch_src FROM categories")
	defer rows.Close()
	var categories []structs.Category
	for rows.Next() {
		var category structs.Category
		err = rows.Scan(&category.ID, &category.Name, &category.EyecatchSrc)
		categories = append(categories, category)
	}
	return categories, err
}

//PostCategory categoriesを新規作成する
func PostCategory(name string, eyecatchSrc string, createdAt time.Time, updatedAt time.Time) error {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	rows, err := db.Query("INSERT INTO categories(name, eyecatch_src, created_at, updated_at) VALUES(?, ?, ?, ?)", name, eyecatchSrc, createdAt, updatedAt)
	defer rows.Close()
	return err
}

//UpdateCategory 指定されたidのcategoriesを変更する
func UpdateCategory(id int, name string, eyecatchSrc string, updatedAt time.Time) error {
	db, err := sql.Open("mysql", config.SQLEnv)
	defer db.Close()
	rows, err := db.Query("UPDATE categories SET name = ?, eyecatch_src = ?, updated_at = ? WHERE id = ?", name, eyecatchSrc, updatedAt, id)
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
