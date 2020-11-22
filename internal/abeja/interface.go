package abeja

import (
	"context"
	"encoding/json"
)

// Database gives the abeja the data to work on.
type Database interface {
	NewID(ctx context.Context, model string) (int, error)
	Objects(ctx context.Context, model string) (map[int]json.RawMessage, error)
	Update(ctx context.Context, model string, id int, data json.RawMessage) error
}
