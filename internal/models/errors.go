package models

import "errors"

var ErrNoRecord = errors.New("models: no matching record found")

var ErrInvalidCreds = errors.New("models: invalid creds")

var ErrDuplicateEmail = errors.New("models: duplicate email")
