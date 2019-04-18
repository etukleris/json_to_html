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

	data, err := getFromUrl(url)
	if err == nil {
		writeToFile(data)
	}
}
	

func getFromUrl(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	result, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}


	//program writes to current_dir\posts; currnent_dir being the directory where
	//the program is placed

func writeToFile(content []byte) {

	postsDir := "posts/"
	if !fileExists(postsDir){
		os.Mkdir(postsDir, os.ModePerm)
	}
	
	//generate filename
	t := time.Now() 
	n:=1
	outputFile := fmt.Sprintf("%sp_%d%02d%02d_%d.html",
		postsDir, t.Year(), t.Month(), t.Day(), n)

	//loop to check if filename exists
	//if file exists, increment n and rename current file
	for {
		
		if fileExists(outputFile) {
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
	
	defer f.Close()
	
	//write html and body tags to file 
    _, err = f.WriteString("<html>\n\n<body>\n\n")
    if err != nil {
        fmt.Println(err)
        return
    }

	
	//struct for json; further code to extract info from content variable
	//>> shouldn't really be called person and people but oh well	
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
		fmt.Println(err)
		return
	}

	for _, personData := range personMap {

	// convert map to array of Person struct
		var p Person
		
		p.userId, _ = strconv.Atoi(fmt.Sprintf("%v", personData["userId"]))
		p.id, _ = strconv.Atoi(fmt.Sprintf("%v", personData["id"]))
		p.title = fmt.Sprintf("%s", personData["title"])
		p.body = fmt.Sprintf("%s", personData["body"])
		people	 = append(people, p)

	 }
	 //write data to the file; using title and body as task requests
	for _, person := range people{
	
		_, err = f.WriteString("<h1>" + person.title + "</h1>\n\n")
		if err != nil {
			fmt.Println(err)
			return
		} 

		_, err = f.WriteString("<p>" + person.body + "</p>\n\n")
		if err != nil {
			fmt.Println(err)
			return
		} 

	}
	    _, err = f.WriteString("</html>\n\n</body>\n\n")
    if err != nil {
        fmt.Println(err)
        return
    }


}


//function to check if file exists
func fileExists(filename string) bool {
	info, err := os.Stat(	filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}


