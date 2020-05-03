package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	ascii "./ascii"
)

type fillerText struct {
	Text string
}

var temp *template.Template
var file = "file.txt"

func main() {
	http.HandleFunc("/", AsciiServer)

	http.HandleFunc("/download", ForceDownload)
	//http.Handle("/"+file, http.FileServer(http.Dir("./")))
	//use - style, js file
	static := http.FileServer(http.Dir("public"))
	//secure, not access another files
	http.Handle("/public/", http.StripPrefix("/public/", static))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)

	}
}

func AsciiServer(w http.ResponseWriter, r *http.Request) {

	temp, err := template.ParseFiles("template/index.html")
	r.ParseForm()
	if r.Method == "GET" {
		if r.URL.Path != "/" {
			errorHandler(w, r, http.StatusNotFound)
			return
		}

		if err != nil {
			errorHandler(w, r, http.StatusInternalServerError)
			return
		}
		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		input := r.FormValue("inputFromHtml")
		fmt.Println(input)
		font := r.FormValue("fontsNameHtml")
		inputNL := strings.Replace(input, "\r\n", "\\n", -1)
		resultAscii := ascii.FontAscii(inputNL, font)

		final := fillerText{Text: resultAscii}
		temp.ExecuteTemplate(w, "index.html", final)
	}
}

func ForceDownload(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "POST" {
		input := r.FormValue("inputFromHtml")
		font := r.FormValue("fontsNameHtml")
		inputNL := strings.Replace(input, "\r\n", "\\n", -1)
		resFs := ascii.FontAscii(inputNL, font)
		b := []byte(resFs)
		fileSize := len(resFs)
		w.Header().Set("Content-Type", "multipart/form-data; charset=utf-8")
		w.Header().Set("Content-Disposition", "attachment; filename=file.txt")
		w.Header().Set("Expires", "0")
		w.Header().Set("Content-Transfer-Encoding", "binary")
		w.Header().Set("Content-Length", strconv.Itoa(fileSize))
		w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

		http.ServeContent(w, r, "file.txt", time.Now(), bytes.NewReader(b))

	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	temp = template.Must(template.ParseGlob("template/*.html"))
	if status == 500 {
		temp.ExecuteTemplate(w, "error1.html", nil)
		return
	}
	if status == 404 {
		temp.ExecuteTemplate(w, "error.html", nil)
		return
	}
}
