package main

import (
	"fmt"
)

type Job struct {
	State string        // strings  in Go are NOT pointer-like
	done  chan struct{} // channels in Go are     pointer-like
}

func (j *Job) Wait() {
	<-j.done
}

func (j *Job) Done() {
	j.State = "done" // changing the copied string
	close(j.done)
}

func main() {
	ch := make(chan Job) // channel of Job, not channel of *Job!
	go func() {
		j := <-ch
		j.Done()
	}()

	job := Job{"ready", make(chan struct{})}
	ch <- job // you are sending a copy of job into the channel: string (State) will be copied, channel (done) will be referenced
	job.Wait()
	fmt.Println(job.State)
}
