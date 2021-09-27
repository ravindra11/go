package main

import (
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// save the text

func (p *Page) save() error {
	filename := p.Title + ".txt"
	// ioutil.writeFile always return null or error only.
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	// p1 := &Page{Title: "TestPage", Body: []byte("Hello world")}
	// p1.save()
	// p2, _ := loadPage(p1.Title)
	// fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// why ResponseWriter not a pointer, while request is pointer
// ResponseWriter is an interface, that's not visible type ResponseWriter interface { }
// https://golang.org/src/net/http/server.go
// r is a pointer to a concrete struct, hence the need to pass a reference explicitly.
// https://golang.org/src/net/http/request.go

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hi There, I Love %s", r.URL.Path[1:])
// }

// view handler
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusNotFound)
		return
	}
	renderTemplate(w, "view.html", p)
	// t, _ := template.ParseFiles("view.html")
	// t.Execute(w, p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	// edit means, when user load a particular file
	// load that page, show the existing content in edit form
	// when user clicks on save button, call the save handler

	// title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit.html", p)
}

func renderTemplate(w http.ResponseWriter, tmplt string, p *Page) {
	// template.ParseFiles will read the contents of edit.html and return a *template.Template
	// the method t.Execute executes the template, writing the generated html to the http.ResponseWritter
	// The .Title and .Body dotted identifiers refer to p.Title and p.Body
	// The printf "%s" .Body instruction is a function call that outputs .Body as a string instead of a stream of bytes
	//, the same as a call to fmt.Printf.

	err := templates.ExecuteTemplate(w, tmplt, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// read body from r.FormValue("body")
	// save file
	// title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		return "", errors.New("invalid page title")
	}
	// the title is the second expression
	return m[2], nil
}

func makeHandler(fn func(w http.ResponseWriter, r *http.Request, template string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title, err := getTitle(w, r)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, title)
	}
}
