package mainpage

import (
	"html/template"
	"net/http"
)

type indexContext struct {
	Ip string
}

type IndexPage struct {
	templateString string
	template       *template.Template
}

var indexTemplate string = `
<!DOCTYPE html>
<html>
    <head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>Your IP address</title>
	<style type="text/css">
	body{
	    color: #444;
	    background-color: #EEE;
	    margin: 40px auto;
	    max-width: 650px;
	    line-height: 1.6em;
	    font-size: 18px;
	    padding: 0;
	}
	</style>
    </head>
    <body>
	<header><h1>Your IP: {{ .Ip }}</h1></header>
    </body>
</html>
`

func (i *IndexPage) prepareTemplate() {
	i.template = template.Must(template.New("index.tmpl").Parse(i.templateString))
}

func (i *IndexPage) GetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := i.template.Execute(w, indexContext{Ip: GetIp(r)})
		if err != nil {
			panic(err)
		}
	}
}

func New() IndexPage {
	index := IndexPage{templateString: indexTemplate}
	index.prepareTemplate()
	return index
}
