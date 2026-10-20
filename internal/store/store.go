package store

import (
	"errors"
	"fmt"
)

// ErrUnknownDriver is returned when an unknown driver name is requested.
var ErrUnknownDriver = errors.New("unknown store driver")

// Open opens a store implementation for the driver name.
func Open(driver, dsn string) (Store, error) {
	switch driver {
	case "sqlite":
		return &sqliteStore{dsn: dsn}, nil
	case "postgres":
		return &postgresStore{dsn: dsn}, nil
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnknownDriver, driver)
	}
}
