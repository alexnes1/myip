package main

import (
	"html/template"
	"net/http"
)

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

func index(template *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := template.Execute(w, IndexContext{Ip: getIp(r)})
		if err != nil {
			panic(err)
		}
	}
}
