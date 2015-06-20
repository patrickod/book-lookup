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
	"strings"
)

func lookupBookByISBN(isbn string) BookResults {
	res, err := http.Get(fmt.Sprintf("https://openlibrary.org/api/books?bibkeys=ISBN:%s&format=json&jscmd=data", isbn))
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	var results BookResults
	err = json.Unmarshal(body, &results)

	if err != nil {
		log.Fatal(err)
	}

	return results
}

func lookupAndRecordBook(isbn string) {
	searchResults := lookupBookByISBN(isbn)
	if len(searchResults) > 0 {
		for k, v := range searchResults {
			recordBook(k, v)
		}
	} else {
		fmt.Printf("Nothing found for %s\n", isbn)
	}
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
		lookupAndRecordBook(strings.Replace(line, "\n", "", -1))
	}
}
