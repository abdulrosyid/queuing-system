package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

type person struct {
	id       int
	duration time.Duration
}

func worker(id int, loket chan person, done chan bool) {
	for l := range loket {
		log.Println("loket", id, "started  job for person ", l.id)
		time.Sleep(l.duration)
		log.Println("loket", id, "finished job for person", l.id, "time process:", l.duration.Seconds(), "second")
		done <- true
	}
	// wg.Done()
}

func main() {
	now := time.Now()
	duration := []int{1, 2, 4, 2, 3, 5, 2, 3, 1, 3}
	var persons []person
	var jmlLoket int

	log.Println("Masukkan jumlah loket:")
	_, err := fmt.Scanf("%d", &jmlLoket)
	if err != nil {
		log.Println("Error Input :", err.Error())
	}

	for i, d := range duration {
		persons = append(persons, person{
			id: i + 1, duration: time.Duration(d) * time.Second,
		})
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	numJobs := len(persons)
	loketJob := make(chan person, numJobs)
	done := make(chan bool, numJobs)

	// jumlah loket yg beroperasi
	for w := 1; w <= jmlLoket; w++ {
		go worker(w, loketJob, done)
	}

	for _, o := range persons {
		loketJob <- o
	}

	close(loketJob)

	for range persons {
		select {
		case <-done:
		case <-quit:
			log.Println("gracefully shutdown goroutines...")
			// wait
			<-done
			goto finish
		}
	}

finish:
	log.Println(time.Since(now))
}
