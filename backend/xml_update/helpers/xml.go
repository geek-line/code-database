package helpers

import (
	"code-database/structs"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"time"
)

//URLSet sitemapのurlsetを表現したstruct
type URLSet struct {
	Xmlns   string   `xml:"xmlns,attr"`
	XMLName xml.Name `xml:"urlset"`
	Urls    []URL    `xml:"url"`
}

//URL sitemapのurlを表現したstruct
type URL struct {
	XMLName    xml.Name `xml:"url"`
	Loc        string   `xml:"loc"`
	Priority   string   `xml:"priority"`
	ChangeFreq string   `xml:"changefreq"`
	LastMod    string   `xml:"lastmod"`
}

func MakeKnowledgesXMLSitemap(publishedKnowledges []structs.Knowledge) URLSet {
	urlSet := URLSet{}
	t := time.Now()
	tstr := t.Format("2006-01-02")
	urls := make([]URL, len(publishedKnowledges))
	for index, knowledge := range publishedKnowledges {
		id := strconv.Itoa(knowledge.ID)
		url := URL{}
		url.Loc = "https://code-database.com/knowledges/" + id
		url.Priority = "1.0"
		url.LastMod = tstr
		url.ChangeFreq = "daily"
		urls[index] = url
	}
	urlSet.Xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"
	urlSet.Urls = urls
	return urlSet
}

func MakeTagsXMLSitemap(tags []structs.Tag) URLSet {
	urlSet := URLSet{}
	t := time.Now()
	tstr := t.Format("2006-01-02")
	urls := make([]URL, len(tags))
	for index, tag := range tags {
		id := strconv.Itoa(tag.ID)
		url := URL{}
		url.Loc = "https://code-database.com/tags/" + id
		url.Priority = "1.0"
		url.LastMod = tstr
		url.ChangeFreq = "daily"
		urls[index] = url
	}
	urlSet.Xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"
	urlSet.Urls = urls
	return urlSet
}

func MakeCategoriesXMLSitemap(categories []structs.Category) URLSet {
	urlSet := URLSet{}
	t := time.Now()
	tstr := t.Format("2006-01-02")
	urls := make([]URL, len(categories))
	for index, category := range categories {
		id := strconv.Itoa(category.ID)
		url := URL{}
		url.Loc = "https://code-database.com/categories/" + id
		url.Priority = "1.0"
		url.LastMod = tstr
		url.ChangeFreq = "daily"
		urls[index] = url
	}
	urlSet.Xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"
	urlSet.Urls = urls
	return urlSet
}

func (urlSet URLSet) UpdateXMLSitemap(xmlFilename string) error {
	filename := "google_sitemap/" + xmlFilename
	output, err := xml.MarshalIndent(urlSet, "  ", "	")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return err
	}
	outputfile, err := os.Create(filename)
	if err != nil {
		fmt.Printf("error: %v\n", err) //ファイルが開けなかったときエラー出力
		return err
	}
	defer outputfile.Close()

	outputfile.Write([]byte(xml.Header))
	outputfile.Write(output)
	return nil
}
