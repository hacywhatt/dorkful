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
		"Sqli":   "inurl:index.php?id=", 
			"intext:\"sql syntax near\"",
			"intext:SQL syntax & inurl:index.php?=id",
		"Önemli": "extension:config", 
			"extension:bak",
			"site:github.com \"BEGIN OPENSSH PRIVATE KEY\"",
			"site:*.*.* intitle:\"index of\" *.pcapng",
			"intitle:\"index of\" \"configuration.php\"",
			"site:com inurl:invoice",
			"intitle:index.of /logs.txt",
			"inurl:\"/wp-content/debug.log\"",
			"intitle:\"index of\" \"config.php.txt\"",
			"intitle:\"index of\" \"*robots.txt\" site:.edu",
			"filetype:log",
			"intext:\"index of\" \"phpinfo\" site:*.in",
			"inurl:/default.rdp",
		"Osint": "site:linkedin.com \"+90\"", 
			"site:github.com \"@gmail.com\"",
			"site:linkedin.com \"@gmail.com\"",
			"site:linkedin.com \"Phone:\" \"+90\"",
			"site:linkedin.com \"Cyber Security Analyst\" \"Phone * * *\"",
			"site:.edu intext:\"robotics\" inurl:/research",
		"Admin": {
			"site:login.*.* site:portal.*.*",
			"site:uat.* * inurl:login",
			"intext:\"user\" filetype:php intext:\"account\" inurl:/admin",
			"inurl:adminpanel site:*.* -site:github.com",
			"site:login.*.* | site:portal.*.*",
			"inurl: edu + site: admin",
			"inurl:/admin.php",
			"intitle:\"index of\" intext: \"login.php\"",
			"site:.co.in intitle:index of /wp-admin",
			"intitle:index.of login.js",
			"inurl:\"admin/default.aspx\"",
			"inurl:\"UserLogin/\" intitle:\"Panel\"",
		},
		"Webcam": {
			"intitle:\"Webcam\" inurl:WebCam.htm",
			"inurl:home.htm intitle:1766",
			"intitle:\"webcam\" \"login\"",
		},

		"Db": {
			"intitle:\"Index of /databases\"",
			"intitle:index.of./.database",
			"intitle:\"index of\" mysql inurl:./db/",
			"\"structure\" + ext:sql",
		},
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
