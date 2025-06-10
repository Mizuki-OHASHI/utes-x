package model

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type ID = string

func NewID() ID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}

// ID がパース可能かわからない場合のパース関数
func ParseID(idStr string) (ID, error) {
	id, err := ulid.Parse(idStr)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

// ID がパース可能であることがわかっている場合のパース関数
func MustParseID(idStr string) ID {
	return ID(idStr)
}
