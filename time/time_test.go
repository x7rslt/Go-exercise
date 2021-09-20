package time_test

import (
	"fmt"
	"log"
	"testing"
	"time"
)

//Measure a piece of code
func TestExecutionTime(t *testing.T){
	start := time.Now()
	// Code to measure
	duration := time.Since(start)

	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Println(duration)

	// Nanoseconds as int64
	fmt.Println(duration.Nanoseconds())
}

func foo() {
	defer duration(track("foo"))
	// Code to measure
}
func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

func TestFuncTime(t *testing.T){
	foo()
}

//Format a time or date

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)
func TestBasicTimeFormat(t *testing.T){
	date := "1999-12-31"
	time1, _ := time.Parse(layoutISO, date)
	fmt.Println(time1)                  // 1999-12-31 00:00:00 +0000 UTC
	fmt.Println(time1.Format(layoutUS)) // December 31, 1999
}


// TimeIn returns the time in UTC if the name is "" or "UTC".
// It returns the local time if the name is "Local".
// Otherwise, the name is taken to be a location name in
// the IANA Time Zone database, such as "Africa/Lagos".
func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func TestLocal(t *testing.T){
	for _, name := range []string{
		"",
		"Local",
		"Asia/Shanghai",
		"America/Metropolis",
	} {
		t, err := TimeIn(time.Now(), name)
		if err == nil {
			fmt.Println(t.Location(), t.Format("15:04"))
		} else {
			fmt.Println(name, "<time unknown>")
		}
	}
}

//get current timestamp

func TestCurrentTime(t *testing.T){
	now := time.Now()      // current local time
	sec := now.Unix()      // number of seconds since January 1, 1970 UTC
	nsec := now.UnixNano() // number of nanoseconds since January 1, 1970 UTC

	fmt.Println(now)  // time.Time
	fmt.Println(sec)  // int64
	fmt.Println(nsec) // int64
}

//Find day in week
func TestWeekday(t *testing.T){
	weekday := time.Now().Weekday()
	fmt.Println(weekday)      // "Tuesday"
	fmt.Println(int(weekday)) // "2"
}