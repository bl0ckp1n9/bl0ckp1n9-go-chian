package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/bl0ckp1n9/bl0ckp1n9chain/blockchain"
)
var templates *template.Template

const (
	port        string = ":4000"
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

func Start() {
	// all of templates files update
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}