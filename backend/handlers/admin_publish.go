package handlers

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"code-database/config"
	"code-database/models"
	"code-database/structs"
)

const lenPathAdminPublish = len(config.AdminPublishPath)

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

func makeXMLSitemap(publishedKnowledges []structs.Knowledge) URLSet {
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

func (urlSet URLSet) updateXMLSitemap() error {
	filename := config.ExecuteDir + "/google_sitemap/knowledges.xml"
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

//AdminPublishHandler /admin/publish/に対するハンドラ
func AdminPublishHandler(w http.ResponseWriter, r *http.Request) {
	suffix := r.URL.Path[lenPathAdminPublish:]
	id, _ := strconv.Atoi(suffix)
	var message string
	if r.Method == "POST" {
		if err := models.SetPublicKnowledge(id); err != nil {
			AdminKnowledgesHandler(w, r)
		}
		message = fmt.Sprintf("{\"text\":\"https://code-database.com/knowledges/%d が公開されました!\nサイトマップにも反映済みです。\"}", id)

	} else {
		if err := models.SetPrivateKnowledge(id); err != nil {
			AdminKnowledgesHandler(w, r)
		}
		message = fmt.Sprintf("{\"text\":\"https://code-database.com/knowledges/%d が非公開になりました!\nサイトマップにも反映済みです。\"}", id)

	}
	if config.BuildMode == "prod" {
		publishedKnowledges, err := models.GetAllPublishedKnowledges()
<<<<<<< Updated upstream
		urlSet := makeXMLSitemap(publishedKnowledges)
		if err = urlSet.updateXMLSitemap(); err != nil {
=======
		urlSet := helpers.MakeKnowledgesXMLSitemap(publishedKnowledges)
		if err = urlSet.UpdateXMLSitemap("knowledges.xml"); err != nil {
>>>>>>> Stashed changes
			AdminKnowledgesHandler(w, r)
		}
		resp, err := http.Post("https://hooks.slack.com/services/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", "applicotion/json", strings.NewReader(message))
		if err != nil {
			AdminKnowledgesHandler(w, r)
		}
		defer resp.Body.Close()
	}
}
