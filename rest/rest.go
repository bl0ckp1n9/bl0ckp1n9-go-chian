package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bl0ckp1n9/bl0ckp1n9chain/blockchain"
	"github.com/bl0ckp1n9/bl0ckp1n9chain/blockchain/utils"
)

var port string 
type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s",port,u)
	return []byte(url),nil
}

type urlDescription struct {
	url url `json:"url"` 
	Method string `json:"method"` 
	Description string `json:"description"`
	Payload string `json:"payload,omitempty"`
}

type AddBlockBody struct {
	Message string 
}



func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			url:url('/'),
			Method: "GET",
			Description: "See Documentation",
		},
		{
			url:url("/blocks"),
			Method: "POST" ,
			Description: "Add Block",
			Payload: "data:string",
		},
		{
			url:url("/blocks/{id}"),
			Method: "POST" ,
			Description: "See a Block",
			Payload: "data:string",
		},
	}
	fmt.Println(data)
	rw.Header().Add("Content-Type","application/json")
	
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
		rw.Header().Add("Content-Type","application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockBody AddBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}
func Start(aPort int) {
	port = fmt.Sprintf(":%d",aPort)
	http.HandleFunc("/", documentation)
	http.HandleFunc("/blocks", blocks)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port,nil))

}