package main

import "time"

type Task struct {
	ID      int
	Title   string
	Done    bool
	Created time.Time
}
