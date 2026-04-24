package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type DorkResult struct {
	Category string
	Dork     string
	Link     string
}

func generateDorks(keyword string) []DorkResult {
	templates := map[string][]string{
		"Sqli":   {"inurl:index.php?id=", "intext:\"sql syntax near\""},
		"Önemli": {"extension:config", "extension:bak"},
		"Osint":  {"site:linkedin.com \"+90\"", "site:github.com \"@gmail.com\""},
	}

	var results []DorkResult
	for cat, dorks := range templates {
		for _, d := range dorks {
			fullDork := d + " " + keyword
			results = append(results, DorkResult{
				Category: cat,
				Dork:     fullDork,
				Link:     "https://www.google.com/search?q=" + strings.ReplaceAll(fullDork, " ", "+"),
			})
		}
	}
	return results
}

func handler(w http.ResponseWriter, r *http.Request) {
	var results []DorkResult
	if r.Method == http.MethodPost {
		results = generateDorks(r.FormValue("keyword"))
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Fprintf(w, "HATA: index.html dosyasını bulamadım :(")
		return
	}
	tmpl.Execute(w, results)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Dorkful hazır. Tarayıcıya şunu yaz: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
