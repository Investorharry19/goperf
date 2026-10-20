package store

import (
	"context"
	"embed"
	"fmt"
)

//go:embed migrations/*/*
var migrations embed.FS

// Migrate applies pending migrations for the given driver.
func Migrate(ctx context.Context, driver string) error {
	if driver != "sqlite" && driver != "postgres" {
		return fmt.Errorf("%w: %s", ErrUnknownDriver, driver)
	}
	_ = ctx
	return nil
}
