package main

import (
	"fmt"
	"time"
)

type Meeting struct {
	name     string
	date     string
	time     string
	duration int
}

func main() {
	// TODO: try read this in via JSON file
	calendar_data := []Meeting{
		{
			name:     "meet1",
			date:     "2021-01-01",
			time:     "09:00",
			duration: 60,
		},
		{
			name:     "meet2",
			date:     "2021-01-01",
			time:     "09:30",
			duration: 30,
		},
		{
			name:     "meet3",
			date:     "2021-01-01",
			time:     "12:30",
			duration: 30,
		},
	}

	datetime_format := "2006-01-0215:04"
	end_time := time.Time{}

	// TODO: need to order list before iterating

	for _, meeting := range calendar_data {
		datetime_str := meeting.date + meeting.time
		start_time, _ := time.Parse(datetime_format, datetime_str)

		if start_time.Before(end_time) {
			fmt.Println(meeting.name)
		}

		end_time = start_time.Add(time.Minute * time.Duration(meeting.duration))
	}
}
