package internal

import "errors"

var ErrNotSaved = errors.New("The Content Is not Saved!")
var ErrEmptyFile = errors.New("Empty file!")
var ErrReadInvalidWords = errors.New("Invalid data exist in file!")
var ErrEmptyPendingWords = errors.New("Nothing to save!")
