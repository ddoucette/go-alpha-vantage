package av

import (
	"strconv"
	"time"

	"github.com/pkg/errors"
)

var location, _ = time.LoadLocation("America/New_York")

// parseFloat parses a float value.
// An error is returned if the value is not a float value.
func parseFloat(val string) (float64, error) {
	return strconv.ParseFloat(val, 64)
}

// parseInt parses an int value.
// An error is returned if the value is not an int value.
func parseInt(val string) (int, error) {
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

// parseDate parses a date value from a string.
// An error is returned if the value is not in one of the dateFormat formats.
func parseDate(v string, dateFormat ...string) (time.Time, error) {
	for _, format := range dateFormat {
		t, err := time.ParseInLocation(format, v, location)
		if err != nil {
			continue
		}
		return t, nil
	}
	return time.Time{}, errors.Errorf("applicable date format not found for date %s", v)
}
