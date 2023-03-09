package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/bl0ckp1n9/bl0ckp1n9chain/blockchain"
	"github.com/bl0ckp1n9/bl0ckp1n9chain/blockchain/utils"
	"github.com/gorilla/mux"
)

var port string 
type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s",port,u)
	return []byte(url),nil
}

type urlDescription struct {
	URL url `json:"url"` 
	Method string `json:"method"` 
	Description string `json:"description"`
	Payload string `json:"payload,omitempty"`
}

type addBlockBody struct {
	Message string 
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}


func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:url('/'),
			Method: "GET",
			Description: "See Documentation",
		},
		{
			URL:url("/blocks"),
			Method: "POST" ,
			Description: "Add Block",
			Payload: "data:string",
		},
		{
			URL:url("/blocks/{height}"),
			Method: "POST" ,
			Description: "See a Block",
			Payload: "data:string",
		},
	}
	// Marshal -> json
	// b, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw,"%s",b);
	
	// 위의 3줄 코드와 동일
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockBody addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	height,err := strconv.Atoi(vars["height"])
	utils.HandleErr(err)
	block,err := blockchain.GetBlockchain().GetBlock(height)

	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	}else {
		encoder.Encode(block)
	}
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter,r *http.Request){
		rw.Header().Add("Content-Type","application/json")
		next.ServeHTTP(rw,r)
	}) 
}

func Start(aPort int) {
	port = fmt.Sprintf(":%d",aPort)
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET","POST")
	router.HandleFunc("/blocks/{height:[0-9]+}",block).Methods("GET")
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port,router))
}