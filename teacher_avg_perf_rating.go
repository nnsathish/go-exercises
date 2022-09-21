/*
Suppose you are a class monitor and you have to take the performance rating of the
class teacher from your classmates. Write a program to take the ratings from all of your
classmates and then print the average rating.

Assumptions
  * There are 200 students in the class.
  * Every student will take a random amount of time to respond.
  * You can simulate the random response time of your classmates by using the
    math/rand package.
*/

package main

import (
  "fmt"
  "math/rand"
  //"runtime"
  "time"
)

type WorkerPool interface {
	Run()
	AddTask(task func())
}

type workerPool struct {
	maxWorker   int
	queuedTaskC chan func()
}

// WorkerPool constructor
func NewWorkerPool(maxWorker int) WorkerPool {
	wp := &workerPool{
		maxWorker:   maxWorker,
		queuedTaskC: make(chan func()),
	}
	return wp
}

func (wp *workerPool) Run() {
	wp.run()
}

func (wp *workerPool) AddTask(task func()) {
	wp.queuedTaskC <- task
}

func (wp *workerPool) GetTotalQueuedTask() int {
	return len(wp.queuedTaskC)
}

func (wp *workerPool) run() {
	for i := 0; i < wp.maxWorker; i++ {
		wID := i + 1
		fmt.Printf("[WorkerPool] Worker %d has been spawned\n", wID)

		go func(workerID int) {
			for task := range wp.queuedTaskC {
				task()
			}
		}(wID)
	}
}

func main() {
  rand.Seed(time.Now().UnixNano())

  go func() {
    fmt.Print("Computing")
		for {
			fmt.Print(".")
			time.Sleep(1 * time.Second)
		}
	}()

  totalWorker := 10
  wp := NewWorkerPool(totalWorker)
  wp.Run()

  type rating struct {
    studentId int
    value     int
  }

  studentsCnt := 200
  ratings := make(chan rating, studentsCnt)

  for i := 0; i < studentsCnt; i++ {
    id := i + 1
    wp.AddTask(func() {
      time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
      ratings <- rating{id, rand.Intn(10)} // 0 is valid rating for now!
    })
  }

  var ratingsSum int
  // try - range over ratings
  for i := 0; i < studentsCnt; i++ {
    rating := <-ratings
    ratingsSum += rating.value
  }
  ratingAvg := (float32(ratingsSum) / float32(studentsCnt))
  fmt.Println("\nAvg Rating:", ratingAvg)
}
