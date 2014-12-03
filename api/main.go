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
    db, err := sql.Open("mymysql", "test/xiemengjun/123456")
    if err != nil {
        panic(err)
    }
    orm := beedb.New(db, "pg")

    sessionArray = append(sessionArray, session)
}

func main()
{
    initDb()
    http.HandleFunc("/", handleIndex)
    http.HandleFunc("/new", handleNewPost)


}