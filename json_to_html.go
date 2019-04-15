package main

import (
	"encoding/json"
	"strconv"
	"fmt"
	"os"
	"time"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://jsonplaceholder.typicode.com/posts"

	data := getFromUrl(url)
	writeToFile(data)
}


func getFromUrl(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	result, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	return result
}


func writeToFile(content []byte) {
	//program writes to dir\posts; dir being the directory where
	//the program is placed
	
	//get current dir
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	//check if posts folder exists in current dir;
	//otherwise create it
	postsDir := dir + "\\posts\\"
	if _, err := os.Stat(postsDir); os.IsNotExist(err) {
		os.Mkdir(postsDir, os.ModePerm)
	}
	//generate filename
	t := time.Now() 
	n:=1
	outputFile := fmt.Sprintf("%sp_%d%02d%02d_%d.html",
		postsDir, t.Year(), t.Month(), t.Day(),n)
	
	//loop to check if filename exists
	for {
		if _, err := os.Stat(outputFile); err == nil {
			//if file exists, increment n and rename current file
			n+=1
			outputFile = fmt.Sprintf("%sp_%d%02d%02d_%d.html",
			postsDir, t.Year(), t.Month(), t.Day(),n)
		} else { break }
		
	}

	//create the output file
	f, err := os.Create(outputFile)
    if err != nil {
        fmt.Println(err)
        return
    }
	//write html and body tags to file 
    _, err = f.WriteString("<html>")
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
	_, err = f.WriteString("\n\n")
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
	_, err = f.WriteString("<body>")
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
 	_, err = f.WriteString("\n\n")
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    } 
	
	//struct for json; further code to extract info from content variable
	jsonByteSlice := content
	type Person struct {
			userId int
			id  int
			title  string
			body  string
         }
	var people []Person

	var personMap []map[string]interface{}

	err = json.Unmarshal(jsonByteSlice, &personMap)

	if err != nil {
		panic(err)
	}

	for _, personData := range personMap {

	// convert map to array of Person struct
		var p Person
		
		p.userId, _ = strconv.Atoi(fmt.Sprintf("%v", personData["userId"]))
		p.id, _ = strconv.Atoi(fmt.Sprintf("%v", personData["id"]))
		p.title = fmt.Sprintf("%s", personData["title"])
		p.body = fmt.Sprintf("%s", personData["body"])
		people = append(people, p)

	 }
	 //write data to the file; using title and body as task requests
	for _, person := range people{
	
		_, err = f.WriteString("<h1>" + person.title + "</h1>")
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		} 
		_, err = f.WriteString("\n\n")
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		} 
		_, err = f.WriteString("<p>" + person.body + "</p>")
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		} 
		_, err = f.WriteString("\n\n")
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		} 
	}
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
   }
	
	
} 