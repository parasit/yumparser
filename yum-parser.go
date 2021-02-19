package main

import (
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var data Metadata

func getFileFromURL(myURL string) []byte {
	client := http.Client{}
	resp, err := client.Get(myURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	tBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return tBytes
}

// GetRepomd downloads repo data from repository
func GetRepomd(baseURL string) Repomd {
	var tRepomd Repomd
	bytesValue := getFileFromURL(baseURL + "/repodata/repomd.xml")
	xml.Unmarshal(bytesValue, &tRepomd)
	// fmt.Println(tRepomd)
	return tRepomd
}

// OpenMetadata reads xml and parse it to struct
func OpenMetadata(metaURL string) {
	byteValue := getFileFromURL(metaURL)
	b := bytes.NewReader(byteValue)
	gz, _ := gzip.NewReader(b)
	defer gz.Close()
	p, _ := ioutil.ReadAll(gz)
	xml.Unmarshal(p, &data)

	fmt.Println("Found " + strconv.Itoa(len(data.Package)) + " packages")

	// for i := 0; i < len(data.Package); i++ {
	// 	fmt.Printf("Package %s ver. %s \n", data.Package[i].Name, data.Package[i].Version.Ver)
	// }
}

func main() {
	var myRepomd Repomd
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Repo url is missing.")
		os.Exit(1)
	}
	var repoURL = args[0]
	myRepomd = GetRepomd(repoURL)
	fmt.Println("Revision " + myRepomd.Revision)
	for i := 0; i < len(myRepomd.Data); i++ {
		if myRepomd.Data[i].Type == "primary" {
			OpenMetadata(repoURL + myRepomd.Data[i].Location.Href)
		}
	}
}
