package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sidneyyi.com/helper"
	"time"
)

func init() {
	helper.TimeLocal, _ = time.LoadLocation("Asia/Shanghai")
	helper.InitLog()
}

func main() {
	http.HandleFunc("/", home14)
	http.HandleFunc("/search", search)
	log.Fatal(http.ListenAndServe("go.sidney.yi:8000", nil))
}

func home14(w http.ResponseWriter, r *http.Request) {
	helper.WriteLog(fmt.Sprintf("%#v", r.PostForm), "gopl.io/ch4/4.14.go::home14")
	fmt.Fprintf(w, "done")
	return
}

func search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	code := query.Get("code")
	helper.WriteLog("debug1 code:"+code+";token:"+helper.Token, "gopl.io/ch4/4.14.go::search")
	var err error
	err = helper.GetToken(w, r, code, "", "http://go.sidney.yi:8000/search")
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	helper.WriteLog("debug2 code:"+code+";token:"+helper.Token, "gopl.io/ch4/4.14.go::search")

	result, err := helper.GetIssues()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}
	show(w, result)
}

func show(w http.ResponseWriter, list []helper.Issue) {
	t := template.Must(template.New("issues").Parse(`<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>`))
	var data struct {
		TotalCount int
		Items      []helper.Issue
	}
	data.TotalCount = len(list)
	data.Items = list
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}
