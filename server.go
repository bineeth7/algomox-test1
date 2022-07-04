package main

import (
	//"net/http"
	"fmt"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"log"
	"os"
	"text/template"
)

type Licenses struct {
	ID           string        `json:"id"`
	Identifiers  []Identifiers `json:"identifiers"`
	Links        []Links       `json:"links"`
	Name         string        `json:"name"`
	OtherNames   []interface{} `json:"other_names"`
	SupersededBy interface{}   `json:"superseded_by"`
	Keywords     []string      `json:"keywords"`
	Text         []Text        `json:"text"`
}
type Identifiers struct {
	Identifier string `json:"identifier"`
	Scheme     string `json:"scheme"`
}
type Links struct {
	Note string `json:"note"`
	URL  string `json:"url"`
}
type Text struct {
	MediaType string `json:"media_type"`
	Title     string `json:"title"`
	URL       string `json:"url"`
}

func main() {
	fmt.Println("Hello") 
	//To create server
    // fmt.Printf("Starting server at port 8080\n")
    // http.ListenAndServe(":8080", nil)
	//setupRoutes()
	l := ReadData()
	WritetoYaml(l)

}

// func setupRoutes(){
// 	http.HandleFunc("/upload",Createfile)
// 	http.ListenAndServe(":8080",nil)
// }

// Createfile reads the filename from the UI
func Createfile() error {
	_, err := os.Create("license.yaml")
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func ReadData() []Licenses {
	f, err := ioutil.ReadFile("license.json")
	var data []Licenses
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(f, &data)
	if err != nil {
	}
	return data
}

func WritetoYaml(l []Licenses) {
	t, err := template.ParseFiles("yml.tmpl")
	if err != nil {
		log.Println(err)
	}
	f, err := os.Create("licenses.yaml")
	f.Close()
	b := &bytes.Buffer{}
	err = t.Execute(b, l)
	ioutil.WriteFile("licenses.yaml", []byte(b.String()), 0644)

}
