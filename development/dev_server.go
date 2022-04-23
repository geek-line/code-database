package development

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

const devServerOrigin = "http://localhost:8080"

func GetStaticFileFromDevServer(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, devServerOrigin+r.RequestURI, 301)
}

func GetTemplateFromDevServer(filePaths ...string) (*template.Template, error) {
	var htmlText string
	for _, filePath := range filePaths {
		resp, err := http.Get(devServerOrigin + "/" + filePath)
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		htmlText += string(body)
	}

	return template.New("test").Parse(htmlText)
}

func GetTemplateFromDevServerWithFuncs(funcMap template.FuncMap, filePaths ...string) (*template.Template, error) {
	var htmlText string
	for _, filePath := range filePaths {
		resp, err := http.Get(devServerOrigin + "/" + filePath)
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		htmlText += string(body)
	}

	return template.New("test").Funcs(funcMap).Parse(htmlText)
}
