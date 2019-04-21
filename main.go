package main

import (
	"encoding/json"
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
		err = writeToFile(data)
            	if err != nil {
                	fmt.Println("There was an error: ", err)
            }
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

func writeToFile(jsonByteSlice []byte) error {

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
        return err
    }
	
	defer f.Close()
	
	//write html and body tags to file 
    _, err = f.WriteString("<html>\n\n<body>\n\n")
    if err != nil {
        fmt.Println(err)
        return err
    }

	
	//struct for json; further code to extract info from content variable
	type Post struct {
			UserId int
			Id  int
			Title  string
			Body  string
         }
	var posts []Post

	err = json.Unmarshal(jsonByteSlice, &posts)

	if err != nil {
		fmt.Println(err)
		return err
	}

	 //write data to the file; using title and body as task requests
	for _, postData := range posts{
	
		_, err = f.WriteString("<h1>" + postData.Title + "</h1>\n\n" +
                               "<p>" + postData.Body + "</p>\n\n")
		if err != nil {
			fmt.Println(err)
			return err
		} 

	}
	    _, err = f.WriteString("</html>\n\n</body>\n\n")
    if err != nil {
        fmt.Println(err)
        return err
    }
    
    return nil

}


//function to check if file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
