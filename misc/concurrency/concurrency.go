package main

import (
	"fmt"
	"sync"
	"time"
)

// Job represents the structure of the task to be processed
type Job struct {
	ID      int
	Payload string
}

// worker is the function that processes jobs from the channel
func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	// Signal to WaitGroup that this worker is done when the function returns
	defer wg.Done()

	// Range over the channel will pull jobs until the channel is closed
	for job := range jobs {
		fmt.Printf("Worker %d started job %d: %s\n", id, job.ID, job.Payload)

		// Simulate a time-consuming task
		time.Sleep(time.Millisecond * 500)

		fmt.Printf("Worker %d finished job %d\n", id, job.ID)
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 3

	// 1. Create the channel for jobs
	// We use a buffered channel so the sender doesn't block immediately
	jobs := make(chan Job, numJobs)

	// 2. Use a WaitGroup to keep track of active workers
	var wg sync.WaitGroup

	// 3. Start the workers (Parallelizing)
	fmt.Printf("Starting %d workers to process %d jobs...\n", numWorkers, numJobs)
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// 4. Send jobs to the channel
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Payload: "Data Task"}
	}

	// 5. CRITICAL: Close the channel to tell workers no more jobs are coming
	close(jobs)

	// 6. Wait for all workers to finish
	wg.Wait()
	fmt.Println("All jobs processed successfully.")
}
