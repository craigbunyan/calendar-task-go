package main

import (
	"fmt"
	"sort"
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
		{
			name:     "meet1",
			date:     "2021-01-01",
			time:     "09:00",
			duration: 60,
		},
	}

	time_format := "15:04"
	datetime_format := "2006-01-02" + time_format
	end_time := time.Time{}

	// TODO: refine this
	sort.SliceStable(calendar_data, func(i, j int) bool {
		x, _ := time.Parse(time_format, calendar_data[i].time)
		y, _ := time.Parse(time_format, calendar_data[j].time)
		return x.Before(y)
	})

	for _, meeting := range calendar_data {
		datetime_str := meeting.date + meeting.time
		start_time, _ := time.Parse(datetime_format, datetime_str)

		if start_time.Before(end_time) {
			fmt.Println(meeting.name)
		}

		end_time = start_time.Add(time.Minute * time.Duration(meeting.duration))
	}
}
