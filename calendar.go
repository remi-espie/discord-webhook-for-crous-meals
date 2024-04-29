package main

import (
	ics "github.com/arran4/golang-ical"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func isEventToday(icsUrl string) bool {
	cal := getIcs(icsUrl)
	if cal == nil {
		return false
	}

	// Get the events of today
	today := time.Now()

	for _, event := range cal.Events() {
		// Print the event as JSON
		at, err := event.GetStartAt()
		if err != nil {
			return false
		}

		if DateEqual(at, today) {
			return true
		}
	}

	return false
}

func getIcs(icsUrl string) *ics.Calendar {
	resp, err := http.Get(icsUrl)
	if err != nil {
		log.Fatal("Error fetching .ics file:", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("Error closing response body:", err)
		}
	}(resp.Body)

	// Read the .ics file contents
	icsData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading .ics file:", err)
	}

	// Parse the .ics data
	cal, err := ics.ParseCalendar(strings.NewReader(string(icsData)))
	if err != nil {
		log.Fatal("Error parsing .ics data:", err)
	}

	return cal
}
