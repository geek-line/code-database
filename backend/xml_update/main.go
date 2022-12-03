package main

import (
	"code-database/config"
	"code-database/models"
	"code-database/xml_update/helpers"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println(config.SQLEnv)
	db, err := sql.Open("mysql", config.SQLEnv)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	publishedKnowledges, err := models.GetAllPublishedKnowledges()
	urlSet := helpers.MakeKnowledgesXMLSitemap(publishedKnowledges)
	if err = urlSet.UpdateXMLSitemap("knowledges.xml"); err != nil {
		fmt.Println(err)
	}
	tags, err := models.GetAllTags()
	urlSet = helpers.MakeTagsXMLSitemap(tags)
	if err = urlSet.UpdateXMLSitemap("tags.xml"); err != nil {
		fmt.Println(err)
	}
	categories, err := models.GetAllCategories()
	urlSet = helpers.MakeCategoriesXMLSitemap(categories)
	if err = urlSet.UpdateXMLSitemap("categories.xml"); err != nil {
		fmt.Println(err)
	}
}
