package helpers

import (
	"encoding/base64"
	"time"
)

func DecodeCursor(cursor string) time.Time {
	ds, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return time.Now()
	}

	timestamp, err := time.Parse("2006-01-02 15:04:05 -0700 MST", string(ds))
	if err != nil {
		return time.Now()
	}

	return timestamp
}
