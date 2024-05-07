package main

import (
	"fmt"
	"log"
	"sync"
)

type Job struct {
	number int
}

func allocateJobs(noOfJobs int, jobs chan<- Job) {
	/*
		writes to chanel like 1,2,3.. 100,
		since it is buffered channel,
		it cant store more than 100 items,
		it will wait for items to be released then will add more
	*/
	for i := 0; i < noOfJobs; i++ {
		jobs <- Job{i + 1}
	}
	close(jobs)
}

func worker(wg *sync.WaitGroup, jobs <-chan Job, results chan<- jsonRecieved) {
	for job := range jobs {
		result, err := Fetch(job.number)
		if err != nil {
			log.Printf("error in fetching: %v\n", err)
		}
		results <- *result
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int, jobs <-chan Job, results chan<- jsonRecieved) {
	var wg sync.WaitGroup
	for i := 0; i <= noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg, jobs, results)
	}
	wg.Wait()
	close(results)
}

func getResults(done chan bool, results <-chan jsonRecieved, resultCollection *[]jsonRecieved) {
	for result := range results {
		if result.Num != 0 {
			fmt.Printf("Retrieving issue #%d\n", result.Num)
			*resultCollection = append(*resultCollection, result)
			// fmt.Println(resultCollection)
		}
	}
	done <- true
}
