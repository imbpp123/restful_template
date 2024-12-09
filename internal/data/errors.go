package data

import "errors"

var (
	ErrParameterNotFound = errors.New("parameter not found")

	ErrArticleNotFound = errors.New("article not found")

	ErrLocationNotFound = errors.New("location not found")
)
