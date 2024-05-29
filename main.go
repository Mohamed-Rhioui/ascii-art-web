package main

import (
	"ascii-art-fs/tools"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Inputs struct {
	Message string
	Banner  string
}

func DrawAsciiArt(elements []string, text string) string {
	var result string
	words := strings.Split(text, `\n`)
	for _, word := range words {
		// replace empty string by new line
		if word != "" {
			for j := 0; j < 8; j++ {
				for _, char := range word {
					if char < 32 || char > 126 {
						fmt.Println("Error: please provide printable characters!!\nhelp: man ascii")
					} else {
						// detect the line from where we should start reading
						start := int(char-32)*8 + j
						result += (elements[start])
					}
				}
				result += "\n"
			}
		} else {
			result += "\n"
		}
	}
	return result
}

func AsciiArtFs(text string, banner string) string {
	// Read from the file standard
	var data string
	// if have to argements we well have choice to work with any template
	data = tools.CHeckTemplate(banner)
	// Split by newline and after that delete the empty strings to organise the file
	var elements []string
	data = strings.ReplaceAll(string(data[1:]), "\r", "\n")
	elements = strings.Split(string(data[1:]), "\n")
	elements = tools.RemoveEmptyString(elements)
	// Split the argument by new line to check every one
	result := DrawAsciiArt(elements, text)
	// handling the additionnel new line if the arguiment is a bunche of new lines
	if tools.IsAllNl(result) {
		result = result[1:]
	}
	// Printing final result
	return result
}

func main() {
	red := "\033[31m"
	blue := "\033[34m"
	reset := "\033[0m"

	tmpl := template.Must(template.ParseFiles("static/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := Inputs{
			Message: r.FormValue("text"),
			Banner:  r.FormValue("Templates"),
		}
		text := details.Message
		banner := details.Banner
		// do something with details
		fmt.Println(banner)
		fmt.Print(AsciiArtFs(text, banner))

		tmpl.Execute(w, struct{ Success bool }{true})
		for _, v := range AsciiArtFs(text, banner) {
			fmt.Fprint(w, string(v))
		}
	})

	// http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println(string(blue), "Server started on :8080", string(reset))
	log.Fatal(string(red), http.ListenAndServe(":8080", nil))
}
