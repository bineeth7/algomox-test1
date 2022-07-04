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
	fmt.Println("Hello") //log.Println("Hello")
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
// func Createfile() error {
// 	_, err := os.Create("license.yaml")
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	return nil
// }

func ReadData() []Licenses {
	http.HandleFunc("/upload", uploadFile)
    http.ListenAndServe(":8080", nil)
	// f, err := ioutil.ReadFile("license.json")
	// var data []Licenses
	// if err != nil {
	// 	log.Println(err)
	// }
	// err = json.Unmarshal(f, &data)
	// if err != nil {
	// }
	// return data
}



func uploadFile(w http.ResponseWriter, r *http.Request) {
    fmt.Println("File Upload Endpoint Hit")

    // Parse our multipart form, 10 << 20 specifies a maximum
    // upload of 10 MB files.
    r.ParseMultipartForm(10 << 20)
    // FormFile returns the first file for the given key `myFile`
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file
    file, handler, err := r.FormFile("myFile")
    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }
    defer file.Close()
    fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    // Create a temporary file within our temp-images directory that follows
    // a particular naming pattern
    tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
    if err != nil {
        fmt.Println(err)
    }
    defer tempFile.Close()

    // read all of the contents of our uploaded file into a
    // byte array
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }
    // write this byte array to our temporary file
    tempFile.Write(fileBytes)
    // return that we have successfully uploaded our file!
    fmt.Fprintf(w, "Successfully Uploaded File\n")
}



// func WritetoYaml(l []Licenses) {
// 	t, err := template.ParseFiles("yml.tmpl")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	f, err := os.Create("licenses.yaml")
// 	f.Close()
// 	b := &bytes.Buffer{}
// 	err = t.Execute(b, l)
// 	ioutil.WriteFile("licenses.yaml", []byte(b.String()), 0644)

// }
