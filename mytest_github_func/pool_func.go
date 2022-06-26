package main

import (
	"fmt"
	"time"
)

type Pool struct {
	Entory  chan *Task
	Jobs    chan *Task
	WorkNum int
}

func NewPool(worknum int) *Pool {
	return &Pool{
		make(chan *Task),
		make(chan *Task),
		worknum,
	}
}

func (p *Pool) Workers() {
	for {
		select {
		case task := <-p.Jobs:
			task.ExcuteFunc()
		}
	}
}

func (p *Pool) Run() {
	for i := 0; i < p.WorkNum; i++ {
		go p.Workers()
	}
	for {
		select {
		case task := <-p.Entory:
			p.Jobs <- task
		}
	}
}

type Task struct {
	f func()
}

func (t *Task) ExcuteFunc() {
	t.f()
}

func main() {

	t := Task{
		func() {
			fmt.Println(time.Now())
		},
	}
	pool := NewPool(3)
	go func() {
		for {
			time.Sleep(time.Second)
			select {
			case pool.Entory <- &t:
			}
		}
	}()
	pool.Run()

}