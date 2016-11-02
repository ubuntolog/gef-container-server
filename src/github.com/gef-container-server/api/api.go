package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"time"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"io"
	"os"
)

// Volume folder content
type VolumeItem struct {
	Name       string `json:"name"`
	Size	   int64 `json:"size"`
	Modified   time.Time `json:"modified"`
	IsFolder   bool `json:"isFolder"`
	FolderTree VolumeItems `json:"folderTree"`
}

type VolumeItems []VolumeItem

type JReply struct {
	Message string `json:"message"`
}

type JPost struct {
	FolderPath string `json:"folderPath"`
}

func readFolders(currentFolder string, volumeItems VolumeItems) VolumeItems {
	files, _ := ioutil.ReadDir(currentFolder)
	for _, f := range files {
		subFolderItems := VolumeItems{}
		if f.IsDir() == true {
			subFolderItems = readFolders(currentFolder + "/" + f.Name(), VolumeItems{})
		}
		volumeItems = append(volumeItems, VolumeItem{Name: f.Name(), Size: f.Size(), Modified: f.ModTime(), IsFolder:f.IsDir(), FolderTree: subFolderItems})
	}
	return volumeItems
}

func Handlers() *mux.Router {
	//router := mux.NewRouter().StrictSlash(true)
	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	router.HandleFunc("/ls", doLsRecursively).Methods("POST")
	router.HandleFunc("/post", doExamplePost)
	return router
	//log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	emptyItems := VolumeItems{}
	volumeItems := readFolders("/Users/megalex/dirlist", emptyItems)
	json.NewEncoder(w).Encode(volumeItems)
}

func doLsRecursively(w http.ResponseWriter, r *http.Request) {
	folderPath := ""
	// Form was POSTed
	if r.FormValue("folderPath") != "" {
		folderPath = r.FormValue("folderPath")
	} else { // JSON was POSTed
		var incomingData JPost
		if r.Body == nil {
			json.NewEncoder(w).Encode(JReply{Message: "Please send a request body"})
			return
		}
		err := json.NewDecoder(r.Body).Decode(&incomingData)
		if err != nil {
			json.NewEncoder(w).Encode(JReply{Message: err.Error()})
			return
		}
		folderPath = incomingData.FolderPath
	}
	fmt.Println(folderPath)
	if folderPath == "" {
		json.NewEncoder(w).Encode(JReply{Message: "The path has not been specified"})
		return
	} else {
		JFolderList := readFolders(folderPath, VolumeItems{})
		json.NewEncoder(w).Encode(JFolderList)
	}
}

func doExamplePost(w http.ResponseWriter, r *http.Request) {
	u := JPost{FolderPath: "/Users/megalex/m"}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	res, _ := http.Post("http://localhost:8080/ls", "application/json; charset=utf-8", b)
	io.Copy(os.Stdout, res.Body)
}