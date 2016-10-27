package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"time"
	"io/ioutil"
	"encoding/json"
)

// Volume folder content
type VolumeItem struct {
	Name       string
	Size	   int64
	Modified   time.Time
	IsFolder   bool
	FolderTree VolumeItems
}

type VolumeItems []VolumeItem

func readFolders(currentFolder string, volumeItems VolumeItems) VolumeItems {
	files, _ := ioutil.ReadDir(currentFolder)
	for _, f := range files {
		//fmt.Println(f.Name())
		subFolderItems := VolumeItems{}
		if f.IsDir() == true {
			subFolderItems = readFolders(currentFolder + "/" + f.Name(), VolumeItems{})
		}
		volumeItems = append(volumeItems, VolumeItem{Name: f.Name(), Size: f.Size(), Modified: f.ModTime(), IsFolder:f.IsDir(), FolderTree: subFolderItems})
	}
	return volumeItems
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Welcome!")

	emptyItems := VolumeItems{}
	volumeItems := readFolders("/Users/megalex/dirlist", emptyItems)
	json.NewEncoder(w).Encode(volumeItems)
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}