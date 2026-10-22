package editor

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"text/template"
	"valterbarros/blog-tools/internal/utils"
)

var f, _ = utils.GetFromRoot("cli/web/editor.html")
var templates = template.Must(template.ParseFiles(f))

type Page struct {
	Title string
	Body  []byte
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func editorHandler(w http.ResponseWriter, r *http.Request) {
	post, err := utils.GetFromRoot("articles/_posts/2026-3-13-mudancas.md")
	if err != nil {
		http.NotFound(w, r)
		return
	}

	d := utils.GetFileData(post)

	p := &Page{
		Title: "Editor",
		Body:  []byte(d),
	}
	renderTemplate(w, "editor", p)
}

func RunServer() error {
	const PORT = "12769"

	signals := make(chan os.Signal, 1)
	// closes server on siging -9
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signals
		log.Println("Shutting down server...")
		os.Exit(0)
	}()

	p, _ := utils.GetFromRoot("cli/web/assets")
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(p))))
	http.HandleFunc("/editor", editorHandler)

	err := http.ListenAndServe(":"+PORT, nil)

	if err != nil {
		log.Fatal("Error to start server: ", err)
		return err
	}

	return nil
}
