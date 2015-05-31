package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func lookupBookByISBN(isbn string) BookResults {
	res, err := http.Get("https://openlibrary.org/api/books?bibkeys=ISBN:9780765312815&format=json&jscmd=data")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	// parse arbitrary JSON
	var results BookResults
	err = json.Unmarshal(body, &results)

	if err != nil {
		log.Fatal(err)
	}

	return results
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		lookupBookByISBN(line)
		fmt.Println(line)
	}
}
