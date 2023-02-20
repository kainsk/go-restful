package helpers

import (
	"encoding/base64"
	"sqlc-rest-api/responses"
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

func NewPageInfo(startCursor, endCursor string, hasNextPage bool) *responses.PageInfo {
	return &responses.PageInfo{
		StartCursor: startCursor,
		EndCursor:   endCursor,
		HasNextPage: hasNextPage,
	}
}
