package postgresql

import (
	"encoding/base64"
	"time"
)

func encodeCursor(cursor time.Time) (string, error) {
	byt, err := cursor.MarshalText()
	if err != nil {
		return "", err
	}

	stringEncode := base64.StdEncoding.EncodeToString(byt)

	return stringEncode, nil
}

func decodeCursor(cursor string) (time.Time, error) {
	byt, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return time.Time{}, err
	}

	cursorTime, err := time.Parse(time.RFC3339, string(byt))
	if err != nil {
		return time.Time{}, err
	}

	return cursorTime, nil
}
