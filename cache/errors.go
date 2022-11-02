package cache

import (
	"errors"
)

var (
	ErrorKeyFound = errors.New("Key not found in cache")
	ErrorReset    = errors.New("cache reset time error")
)
