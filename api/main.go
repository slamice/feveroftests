package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
    r "github.com/astaxie/beedb"
)

var sessionArray []*t.Session

type post struct {
	Title string
	body  []byte
}

func initDb()
{
}

func main()
{
    initDb()
    http.HandleFunc("/", handleIndex)
    http.HandleFunc("/user/add", handleNewUser)

    err := 

}