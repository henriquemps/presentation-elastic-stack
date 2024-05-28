package main

import (
	"bytes"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"net/http"
)

var connEs *elasticsearch.Client
var cfgEs = elasticsearch.Config{
	CloudID: "da2790b693d04a789afea747756a8834:dXMtY2VudHJhbDEuZ2NwLmNsb3VkLmVzLmlvJDgwY2MzMjIxZjkzYTQ0ZDJiZjk5ZWQyYmU5ODM3ZmM1JGQ4ZjcxZDVlZmQ0MDRlNzU5ZDlmMjczYzIxMmJiYjgw",
	APIKey:  "a19PU3ZJOEIyaklzV3JCRlF3OHY6QmFSZWM5ajZTbWlucllWSmhXVnBuZw==",
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func handlerElasticsearch(w http.ResponseWriter, r *http.Request) {

	buf := bytes.NewBufferString(`
{ "index": { "_index": "authors" } }
{"name":"Snow Crash","author":"Neal Stephenson","release_date":"1992-06-01","page_count": 470}
{ "index": { "_index": "authors" } }
{"name": "Revelation Space", "author": "Alastair Reynolds", "release_date": "2000-03-15", "page_count": 585}
{ "index": { "_index": "authors" } }
{"name": "1984", "author": "George Orwell", "release_date": "1985-06-01", "page_count": 328}
{ "index": { "_index": "authors" } }
{"name": "Fahrenheit 451", "author": "Ray Bradbury", "release_date": "1953-10-15", "page_count": 227}
{ "index": { "_index": "authors" } }
{"name": "Brave New World", "author": "Aldous Huxley", "release_date": "1932-06-01", "page_count": 268}
{ "index": { "_index": "authors" } }
{"name": "The Handmaid's Tale", "author": "Margaret Atwood", "release_date": "1985-06-01", "page_count": 311}
`)

	ingestResult, err := connEs.Bulk(
		bytes.NewReader(buf.Bytes()),
		connEs.Bulk.WithIndex("authors"),
	)

	fmt.Println(ingestResult, err)

	fmt.Fprintf(w, "Gravando dados no elastic")
}

func init() {
	var err error

	connEs, err = elasticsearch.NewClient(cfgEs)

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
}

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/elasticsearch", handlerElasticsearch)

	http.ListenAndServe(":8080", nil)
}
