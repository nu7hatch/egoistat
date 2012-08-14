package main

import (
	"encoding/json"
	"github.com/nu7hatch/egoistat/backend"
	"log"
	"net/http"
	"net/url"
	"strings"
	"text/template"
	"flag"
)

var Addr string

func init() {
	flag.StringVar(&Addr, "addr", ":8080", "The address to serve on")
	flag.Parse()
}

var scriptTpl = "var egoistat={c:{{.Data}}, count: function(sn){return this.c[sn] || 0;}};\n{{.Callback}};"

func transformParams(form url.Values) (res map[string]string) {
	res = make(map[string]string)
	for k, v := range form {
		res[k] = strings.Join(v, "\n")
	}
	return
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var networks = strings.Split(r.FormValue("n"), ",")
	var url = r.FormValue("url")
	var params = transformParams(r.Form)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	request := egoist.NewRequest(url, params)
	counts := request.Count(networks...)

	enc := json.NewEncoder(w)
	enc.Encode(counts)
}

type countScriptData struct {
	Data     string
	Callback string
}

func countScriptHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var networks = strings.Split(r.FormValue("n"), ",")
	var url = r.FormValue("url")
	var params = transformParams(r.Form)
	var callback = r.FormValue("cb")

	if len(strings.TrimSpace(callback)) > 0 {
		callback = callback + "()"
	}

	w.Header().Set("Content-Type", "text/javascript")
	w.WriteHeader(200)

	request := egoist.NewRequest(url, params)
	counts := request.Count(networks...)
	data, _ := json.Marshal(counts)

	tmpl, _ := template.New("script").Parse(scriptTpl)
	tmpl.Execute(w, countScriptData{string(data), callback})
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/index.html")
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/stat/", indexHandler)
	http.HandleFunc("/api/v1/count.json", countHandler)
	http.HandleFunc("/api/v1/count.js", countScriptHandler)

	log.Fatal(http.ListenAndServe(Addr, nil))
}
