package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	_map := map[string][]*Issue{}
	now := time.Now().Unix()
	for _, item := range result.Items {
		_tDiff := now - item.CreatedAt.Unix()
		if _tDiff < 30*86400 {
			_map["不到一个月的"] = append(_map["不到一个月的"], item)
		} else if _tDiff < 360*86400 {
			_map["不到一年的"] = append(_map["不到一年的"], item)
		} else
		{
			_map["超过一年的"] = append(_map["超过一年的"], item)
		}
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for k, _list := range _map {
		fmt.Printf("%s:\n", k)
		for _, item := range _list {
			fmt.Printf("#%-5d %9.9s %.55s %s\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt.Format("2006-01-02 15:04:05"))
		}
	}
}

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
