package handlers

import (
	"code-database/models"
	"code-database/structs"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type ApiTagsQuery struct {
	Size    int
	Offset  int
	sortKey string
}

func parseQuery(query url.Values) (ApiTagsQuery, error) {
	var errMsg, sortKey string
	var size, offset int
	var err error
	if size, err = strconv.Atoi(query.Get("size")); err != nil {
		errMsg += "invalid size,"
	}
	if offset, err = strconv.Atoi(query.Get("offset")); err != nil {
		errMsg += "invalid offset,"
	}
	if sortKey = query.Get("sort"); sortKey != "reference" {
		errMsg += "invalid sortKey,"
	} else {
		sortKey = "count"
	}
	if errMsg == "" {
		return ApiTagsQuery{Size: size, Offset: offset, sortKey: sortKey}, nil
	} else {
		return ApiTagsQuery{}, errors.New(errMsg)
	}
}

func ApiTagsHandler(w http.ResponseWriter, r *http.Request, auth bool) {
	if r.Method == "GET" {
		query, err := parseQuery(r.URL.Query())
		if err != nil {
			fmt.Println("error: " + err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tags, err := models.GetTagsAsJson(query.Offset, query.Size, query.sortKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tagsResp := structs.TagsJson{
			Tags: tags,
		}
		respJson, err := json.Marshal(tagsResp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJson)
	}
}
