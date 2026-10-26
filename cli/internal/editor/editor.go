package editor

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"text/template"
	"valterbarros/blog-tools/internal/content"
	"valterbarros/blog-tools/internal/utils"
)

var f, _ = utils.GetFromRoot("cli/web/editor.html")
var f2, _ = utils.GetFromRoot("cli/web/index.html")
var templates = template.Must(template.ParseFiles(f, f2))

type Page struct {
	Title string
	Body  []byte
}

func renderTemplate[T any](w http.ResponseWriter, tmpl string, p *T) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func editorHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	post, _ := utils.GetFromRoot("articles/_posts/" + name)

	d, err := utils.GetFileData(post)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	p := &Page{
		Title: "Editor",
		Body:  []byte(d),
	}
	renderTemplate(w, "editor", p)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	ps, err := content.ListPosts()

	if err != nil {
		http.NotFound(w, r)
		return
	}

	list := strings.Split(ps, "\n")

	renderTemplate(w, "index", &list)
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

	http.HandleFunc("/ws", handleWebSocket)

	p, _ := utils.GetFromRoot("cli/web/assets")
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(p))))
	http.HandleFunc("/editor", editorHandler)
	http.HandleFunc("/", listHandler)

	err := http.ListenAndServe(":"+PORT, nil)

	if err != nil {
		log.Fatal("Error to start server: ", err)
		return err
	}

	return nil
}
