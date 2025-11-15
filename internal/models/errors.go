package models

import "errors"

// ErrNoRecord is returned when a specific record could not be found in the database.
var ErrNoRecord = errors.New("models: no matching record found")
