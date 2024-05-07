package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type jsonRecieved struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	Safe_Title string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func main() {
	router := gin.Default()
	// Use CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                                         // Allow all origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                   // Allow all HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"}, // Allow specific headers
		ExposeHeaders:    []string{"Content-Length"},                                            // Expose specific headers
		AllowCredentials: true,                                                                  // Allow credentials such as cookies
		MaxAge:           86400,                                                                 // Maximum age of the CORS options request
	}))
	router.GET("/api/load", mainFunc)

	router.Run("localhost:8000")
}

func mainFunc(c *gin.Context) {
	// allocate jobs or inother words total number of requests to be sent
	noOfJobs := 100
	var jobs = make(chan Job, 100)
	var results = make(chan jsonRecieved, 100)
	var resultCollection []jsonRecieved

	go allocateJobs(noOfJobs, jobs)

	// get results
	done := make(chan bool)
	go getResults(done, results, &resultCollection)

	// create worker pool
	noOfWorkers := 100
	createWorkerPool(noOfWorkers, jobs, results)

	// wait for all results to be collected
	<-done

	// convert result collection to JSON
	// data, err := json.MarshalIndent(resultCollection, "", "    ")
	// if err != nil {
	// 	log.Fatal("json err: ", err)
	// 	c.IndentedJSON(http.StatusExpectationFailed, err)
	// 	return
	// }

	// write json data to file
	// err = writeToFile(data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("resultCollection")
	// fmt.Println(resultCollection)

	fmt.Println(resultCollection[0])

	c.IndentedJSON(http.StatusOK, resultCollection)
	return
}

func writeToFile(data []byte) error {
	/*
		creates a file and dump the data there
	*/
	f, err := os.Create("xkcd.json")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}
