package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Websites struct {
	Backend       map[string]string
	Documentation map[string]string
	Frontend      map[string]string
	Job           map[string]string
	Social        map[string]string
	Useful        map[string]string
}

func readFile(filePath string) []string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("File reading error -> ", err)
	}

	return strings.Split(string(content), "\n")
}

func readTemplate(filePath string) *template.Template {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	template, err := template.ParseFiles(filePath)
	if err != nil {
		panic(err)
	}

	return template
}

func prepUrlsForTemplate(urls []string) map[string]string {
	var mappedWebsites = make(map[string]string)

	for index := 0; index < len(urls); index += 3 {
		mappedWebsites[urls[index]] = urls[index+1]
	}

	return mappedWebsites
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	path := "data"
	backendUrls := readFile(path + "/backend.txt")
	documentationUrls := readFile(path + "/documentation.txt")
	frontendUrls := readFile(path + "/frontend.txt")
	jobUrls := readFile(path + "/job.txt")
	socialUrls := readFile(path + "/social.txt")
	usefulUrls := readFile(path + "/useful.txt")

	indexTemplate := readTemplate("templates/index.html")

	dataForTemplate := Websites{
		Backend:       prepUrlsForTemplate(backendUrls),
		Documentation: prepUrlsForTemplate(documentationUrls),
		Frontend:      prepUrlsForTemplate(frontendUrls),
		Job:           prepUrlsForTemplate(jobUrls),
		Social:        prepUrlsForTemplate(socialUrls),
		Useful:        prepUrlsForTemplate(usefulUrls),
	}

	err := indexTemplate.Execute(w, dataForTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func robotsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/robots.txt")
}

func humansHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/humans.txt")
}

func main() {
	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/robots.txt", robotsHandler)
	router.HandleFunc("/humans.txt", humansHandler)

	fmt.Println("Server is running on port 55555...")
	err := http.ListenAndServe(":55555", router)
	if err != nil {
		log.Fatal("Server error -> ", err)
	}
}
