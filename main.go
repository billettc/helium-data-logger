package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/billettc/helium-data-logger/db"
	"github.com/billettc/helium-data-logger/models"
)

var f *os.File
var database *db.MongoDB

func main() {
	var err error
	f, err = os.Create("/Users/cbillett/t/helium-data.log")
	check(err)
	defer f.Close()

	database, err = db.NewMongoDB("mongodb://localhost:27017")
	check(err)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":915", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering from panic", r)
		}
	}()

	data, err := ioutil.ReadAll(r.Body)
	if handleError(err, w) {
		return
	}

	jsonData := string(data)
	fmt.Println("json:", jsonData)
	if jsonData == "" || !strings.Contains(jsonData, "app_eui") {
		return
	}

	_, err = f.Write([]byte(jsonData + "\n"))
	if handleError(err, w) {
		return
	}
	err = f.Sync()
	if handleError(err, w) {
		return
	}

	le := &models.LogEvent{}
	err = json.Unmarshal(data, &le)
	if handleError(err, w) {
		return
	}

	err = database.SaveLogEvent(le)
	if handleError(err, w) {
		return
	}

}

func handleError(e error, w http.ResponseWriter) bool {
	if e != nil {
		fmt.Println("err :=>", e)
		_, err := w.Write([]byte(e.Error()))
		if err != nil {
			fmt.Println("ERROR: ", err)
		}
		w.WriteHeader(500)
		return true
	}

	return false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
