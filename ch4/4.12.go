package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sidneyyi.com/helper"
	"strconv"
)

var IndexPath = "./4.12.json"
var Index map[int]XkcdData

type XkcdData struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	Safe_title string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func init() {
	//检查离线索引有没有，没有则建立离线索引
	b, err := helper.PathExists(IndexPath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("是否存在离线索引 %v\n", b)
	if !b {
		file, err := os.OpenFile(IndexPath, os.O_WRONLY|os.O_CREATE, 0766)
		if err != nil {
			panic(err)
		}

		header := http.Header{}
		Index = map[int]XkcdData{}
		for i := 1; i <= 1; i++ {
			jsonStr, err := helper.Request("GET", "https://xkcd.com/"+strconv.Itoa(i)+"/info.0.json", []byte{}, header, http.StatusOK)
			if err != nil {
				panic(err)
			}
			var _d XkcdData
			if err := json.Unmarshal([]byte(jsonStr), &_d); err != nil {
				panic(err)
			}
			//fmt.Printf("%s\n%#v\n", jsonStr, _d)
			Index[_d.Num] = _d
		}

		_json, _ := json.Marshal(Index)

		//fmt.Printf("%s\n", _json)
		//os.Exit(111)

		file.WriteString(string(_json))
		file.Close()

	} else {
		file, err := os.Open(IndexPath)
		if err != nil {
			panic(err)
		}
		_json, _ := ioutil.ReadAll(file)
		if err := json.Unmarshal([]byte(_json), &Index); err != nil {
			panic(err)
		}
	}

	//fmt.Printf("%#v\n", Index)
}

func main() {
	id, _ := strconv.Atoi(os.Args[1])
	fmt.Println(Index[id].Title, Index[id].Img)
}
