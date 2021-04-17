package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/billettc/helium-data-logger/db"
	"github.com/billettc/helium-data-logger/models"
)

func main() {
	file, err := os.Open("/Users/cbillett/t/helium-data.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	database, err := db.NewMongoDB("mongodb://localhost:27017")
	check(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jsonData := scanner.Text()
		fmt.Println("data:", jsonData)
		fmt.Println("--------------------")

		if jsonData == "" || !strings.Contains(jsonData, "app_eui") {
			continue
		}

		le := &models.LogEvent{}
		err = json.Unmarshal([]byte(jsonData), &le)
		check(err)

		err = database.SaveLogEvent(le)
		check(err)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
