package main

import (
	"fmt"
	"log"
	"time"
)

type person struct {
	id       int
	duration time.Duration
}

func worker(id int, loket chan person, done chan bool) {
	for l := range loket {
		log.Println("loket", id, "started job for person ", l.id)
		time.Sleep(l.duration)
		log.Println("loket", id, "finished job for person", l.id, "time process:",l.duration.Seconds(),"second")
		done <- true
	}
}

func main()  {
	duration := []int{1, 1, 5, 1, 1}
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

	numJobs := len(persons)
	loket := make(chan person, numJobs)
	done := make(chan bool, numJobs)

	// jumlah loket yg beroperasi
	for w := 1; w <= jmlLoket; w++ {
		go worker(w, loket, done)
	}

	for _, o := range persons {
		loket <- o
	}
	close(loket)

	for j := 1; j <= numJobs; j++ {
		<-done
	}
}