package main

import (
	"html/template"
	"net/http"
	"os"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ascii-art", processor)
	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
	err := tpl.ExecuteTemplate(w, "layout.html", res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

var res string

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if r.URL.Path != "/ascii-art" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	text := r.FormValue("box")
	banner := r.FormValue("font-files")
	buffer, err := os.ReadFile(banner + ".txt")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var ascii []string = strings.Split(strings.ReplaceAll(string(buffer), "\r", ""), "\n")
	res = AsciiArt(text, ascii)
	if text != "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} /* else {
		}
	}*/
}
func RowFinder(b byte) int {
	return (int(b)-32)*9 + 1
}
func AsciiArt(input string, font []string) string {
	var result string
	// split the input by new line whenever it finds \n
	temp := strings.ReplaceAll(input, "\r\n", "\\n")
	lines := strings.Split(temp, "\\n")
	// looping through the slice that contains the arguments
	for i := 0; i < len(lines); i++ {
		switch lines[i] {
		case "":
			result += "\n"
		default:
			for j := 0; j < 8; j++ {
				for k := 0; k < len(lines[i]); k++ {
					//fmt.Println(j,k)
					result += (font[RowFinder(lines[i][k])+j])
				}
				result += "\n"
			}
		}
	}
	return result
}
