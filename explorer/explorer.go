package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/bl0ckp1n9/bl0ckp1n9chain/blockchain"
	"github.com/gorilla/mux"
)
var templates *template.Template

const (
	templateDir string = "explorer/templates/"
)


type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(w http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	templates.ExecuteTemplate(w, "home", data)
}


func add(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(w, "add" , nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(w,r,"/",http.StatusPermanentRedirect)
	}
}

func Start(aPort int) {
	// handler := http.NewServeMux()
	router := mux.NewRouter()
	// all of templates files update
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	port := fmt.Sprintf(":%d",aPort)
	router.HandleFunc("/", home)
	router.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}