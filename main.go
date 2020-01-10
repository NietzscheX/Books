package main

import (
	"books/model"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", IndexView)
	http.HandleFunc("/list", ListView)
	http.HandleFunc("/api/index/data", IndexData)
	http.HandleFunc("/api/list/data", ListData)
	http.ListenAndServe(":80", nil)
}

func IndexView(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./templates/index.html")
	io.Copy(w, f)
	f.Close()
}

func ListView(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./templates/list.html")
	io.Copy(w, f)
	f.Close()
}

//单条数据
func IndexData(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idstr := r.Form.Get("id")
	id, _ := strconv.ParseInt(idstr, 10, 64)
	mod, err := model.ArticleGet(id)
	if err != nil {
		Fail(w, err.Error())
		//w.Write([]byte(err.Error()))
		return
	}
	buf, _ := json.Marshal(mod)
	w.Write(buf)
}

func Fail(w http.ResponseWriter, msg string) {
	w.Write([]byte(msg))
}

//集合数据
func ListData(w http.ResponseWriter, r *http.Request) {
	mods, err := model.ArticleList()
	if err != nil {
		Fail(w, err.Error())
		return
	}

	buf, _ := json.Marshal(mods)
	w.Write(buf)

}
