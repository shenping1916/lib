package cache

import "errors"

var (
	// current number of cache tables close to maxtable
	ErrorNewTable = errors.New("The number of cache tables has reached the limit!")

	// key isn't exist
	ErrorKey = errors.New("key can't be found!")
)
