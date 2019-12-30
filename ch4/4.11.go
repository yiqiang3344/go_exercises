package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sidneyyi.com/helper"
	"time"
)

func init() {
	helper.TimeLocal, _ = time.LoadLocation("Asia/Shanghai")
	helper.InitLog()
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/tool", tool)
	log.Fatal(http.ListenAndServe("go.sidney.yi:8000", nil))
}

func printItem(w http.ResponseWriter, list []helper.Issue) {
	fmt.Fprintf(w, "%-7s %20s %20s %35s %35s %25s %25s %25s\n ", "id", "user", "state", "title", "body", "create time", "update time", "close time")
	for _, item := range list {
		fmt.Fprintf(w, "#%-5d %20s %20s %35s %35s %25s %25s %25s\n ", item.Number, item.User.Login, item.State, item.Title, item.Body, item.CreatedAt.In(helper.TimeLocal).String()[0:19], item.UpdatedAt.In(helper.TimeLocal).String()[0:19], item.ClosedAt.In(helper.TimeLocal).String()[0:19])
	}
}

func tool(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	code := query.Get("code")
	state := query.Get("state")
	var err error
	err = helper.GetToken(w, r, code, state, "http://go.sidney.yi:8000/tool")
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	helper.WriteLog("code:"+code+";state:"+state+";token:"+helper.Token, "tool")

	switch state {
	case "get":
		result, err := helper.GetIssues()
		if err != nil {
			fmt.Fprintf(w, "%s", err)
			return
		}
		printItem(w, result)
	case "insert":
		params := helper.CreateIssueParams{Title: "测试一下 " + time.Now().In(helper.TimeLocal).String()[0:19], Body: "只是测试", Assignees: []string{"yiqiang3344"}, Labels: []string{"invalid"}}
		//fmt.Fprintf(w, "%#v", params)
		//return
		data, err := json.Marshal(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		item, err := helper.CreateIssues(data)
		if err != nil {
			fmt.Fprintf(w, "%s", err)
			return
		}
		printItem(w, []helper.Issue{*item})
	case "update":
		params := helper.CreateIssueParams{Body: "测试更新"}
		data, err := json.Marshal(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		//fmt.Fprintf(w, "%#v", string(data))
		//return
		id := query.Get("id")
		item, err := helper.UpdateIssues(id, data)
		if err != nil {
			fmt.Fprintf(w, "%s", err)
			return
		}
		printItem(w, []helper.Issue{*item})
	case "close":
		params := helper.CreateIssueParams{State: "close"}
		data, err := json.Marshal(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		//fmt.Fprintf(w, "%#v", string(data))
		//return
		id := query.Get("id")
		item, err := helper.UpdateIssues(id, data)
		if err != nil {
			fmt.Fprintf(w, "%s", err)
			return
		}
		printItem(w, []helper.Issue{*item})
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	code := query.Get("code")
	fmt.Fprintf(w, "首页%s\n", code)
}
