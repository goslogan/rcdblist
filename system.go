package rcutils

import (
	"encoding/csv"
	"io"
	"time"

	"github.com/jszwec/csvutil"
)

// LogEvent represents a single line from the system log downloaded from redis Cloud
type LogEvent struct {
	Id        int64     `csv:"id"`
	User      string    `csv:"user name"`
	Email     string    `csv:"email"`
	Activity  string    `csv:"activity"`
	TimeStamp time.Time `csv:"date"`
	Database  string    `csv:"database name"`
	Change    string    `csv:"description"`
}

// SystemLog reads the system log downloaded from Redis Cloud and parses it into a useful format.
// If not all the events in the log are required, the filter function parameter can be used to exclude undesired events.
// If it returns false, the event will not be returned.
func SystemLog(source io.Reader, filter func(*LogEvent) bool) ([]*LogEvent, error) {

	events := []*LogEvent{}

	reader := csv.NewReader(source)

	if decoder, err := csvutil.NewDecoder(reader); err != nil {
		return nil, err
	} else {
		for {
			var event LogEvent
			if err := decoder.Decode(&event); err == io.EOF {
				break
			} else if err != nil {
				return nil, err
			} else {
				if filter(&event) {
					events = append(events, &event)
				}
			}
		}
	}

	return events, nil
}
