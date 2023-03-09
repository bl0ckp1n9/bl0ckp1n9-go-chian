package main

import (
	"github.com/bl0ckp1n9/bl0ckp1n9chain/explorer"
	"github.com/bl0ckp1n9/bl0ckp1n9chain/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)

}
