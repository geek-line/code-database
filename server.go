package main

import (
	"net/http"
	"os"

	"code-database/config"
	"code-database/handlers"
	"code-database/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, config.UserKnowledgesPath, http.StatusFound)
}

func main() {
	dir, _ := os.Getwd()
	http.HandleFunc(config.RootPath, redirectHandler)
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
	http.HandleFunc(config.UserKnowledgePath, middleware.UserAuth(handlers.KnowledgeHandler))
	http.HandleFunc(config.UserKnowledgesLikePath, middleware.UserAuth(handlers.KnowledgeLikeHandler))
	http.HandleFunc(config.UserTagsPath, middleware.UserAuth(handlers.TagsHandler))
	http.HandleFunc(config.UserTagPath, middleware.UserAuth(handlers.TagHandler))
	http.HandleFunc(config.UserCategoriesPath, middleware.UserAuth(handlers.CategoriesHandler))
	http.HandleFunc(config.UserCategoryPath, middleware.UserAuth(handlers.CategoryHandler))
	http.Handle(config.StaticPath, http.StripPrefix(config.StaticPath, http.FileServer(http.Dir(dir+config.StaticPath))))
	http.Handle(config.NodeModulesPath, http.StripPrefix(config.NodeModulesPath, http.FileServer(http.Dir(dir+config.NodeModulesPath))))
	http.Handle(config.GoogleSitemapPath, http.StripPrefix(config.GoogleSitemapPath, http.FileServer(http.Dir(dir+config.GoogleSitemapPath))))
	http.ListenAndServe(":3000", nil)
}