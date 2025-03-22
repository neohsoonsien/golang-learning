package channel

import "log"

func Close() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// worker goroutine
	go func() {
		for {
			// 'more' value will be false if `jobs`` has been closed and all values in the channel have been received
			j, more := <-jobs
			if more {
				log.Println("received job", j)
			} else {
				log.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// sends 3 jobs to the worker over the jobs channel
	for j := 1; j <= 3; j++ {
		jobs <- j
		log.Println("sent job", j)
	}
	close(jobs)
	log.Println("sent all jobs")

	// await the worker using the synchronization approach
	<-done

	// reading from the closed channel
	j, more := <-jobs
	log.Println("received job", j)
	log.Println("received more jobs:", more)
}
