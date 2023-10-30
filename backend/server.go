package main

import (
	"fmt"
	"net/http"
	"os"

	"code-database/config"
	"code-database/development"
	"code-database/handlers"
	"code-database/middleware"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/csrf"
)

func main() {
	dir, _ := os.Getwd()
	key := make([]byte, 32)
	CSRF := csrf.Protect([]byte(key))
	http.HandleFunc(config.RootPath, middleware.UserAuth(handlers.TopHandler))
	http.HandleFunc(config.AdminLoginPath, middleware.UserAuth(handlers.AdminLoginHandler))
	http.HandleFunc(config.AdminLogoutPath, middleware.AdminAuth(handlers.AdminLogoutHandler))
	http.HandleFunc(config.AdminKnowledgesPath, middleware.AdminAuth(handlers.AdminKnowledgesHandler))
	http.HandleFunc(config.AdminTagsPath, middleware.AdminAuth(handlers.AdminTagsHandler))
	http.HandleFunc(config.AdminNewPath, middleware.AdminAuth(handlers.AdminNewHandler))
	http.HandleFunc(config.AdminSavePath, middleware.AdminAuth(handlers.AdminSaveHandler))
	http.HandleFunc(config.AdminDeletePath, middleware.AdminAuth(handlers.AdminDeleteHandler))
	http.HandleFunc(config.AdminPublishPath, middleware.AdminAuth(handlers.AdminPublishHandler))
	http.HandleFunc(config.AdminEyecatchesPath, middleware.AdminAuth(handlers.AdminEyecatchesHandler))
	http.HandleFunc(config.AdminCategoriesPath, middleware.AdminAuth(handlers.AdminCategoriesHandler))
	http.HandleFunc(config.UserKnowledgesPath, middleware.UserAuth(handlers.KnowledgesHandler))
	http.HandleFunc(config.UserSearchPath, middleware.UserAuth(handlers.SearchHandler))
	http.Handle(config.UserKnowledgePath, CSRF(middleware.UserAuth(handlers.KnowledgeHandler)))
	http.Handle(config.UserKnowledgesLikePath, CSRF(middleware.UserAuth(handlers.KnowledgeLikeHandler)))
	http.HandleFunc(config.UserTagsPath, middleware.UserAuth(handlers.TagsHandler))
	http.HandleFunc(config.UserTagPath, middleware.UserAuth(handlers.TagHandler))
	http.HandleFunc(config.UserCategoriesPath, middleware.UserAuth(handlers.CategoriesHandler))
	http.HandleFunc(config.UserCategoryPath, middleware.UserAuth(handlers.CategoryHandler))
	http.HandleFunc(config.UserAboutPath, middleware.UserAuth(handlers.AboutHandler))
	http.HandleFunc(config.UserPrivacyPath, middleware.UserAuth(handlers.PrivacyHandler))
	http.HandleFunc(config.RssPath, handlers.RssHandler)
	if config.BuildMode == "prod" {
		http.Handle(config.StaticPath, http.StripPrefix(config.StaticPath, http.FileServer(http.Dir(dir+config.StaticPath))))
		http.Handle(config.PublicPath, http.StripPrefix(config.PublicPath, http.FileServer(http.Dir(dir+config.PublicPath))))
	} else {
		http.HandleFunc(config.StaticPath, development.GetStaticFileFromDevServer)
		http.HandleFunc(config.PublicPath, development.GetStaticFileFromDevServer)
	}

	http.Handle(config.GoogleSitemapPath, http.StripPrefix(config.GoogleSitemapPath, http.FileServer(http.Dir(dir+config.GoogleSitemapPath))))
	fmt.Printf("Success starting backend server (%s build)\n", config.BuildMode)
	http.ListenAndServe(":8080", nil)
}
