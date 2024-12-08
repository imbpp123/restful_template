package data

import "errors"

var (
	ErrUnsupportedAPIVersion = errors.New("unsupported api version")
	ErrParameterNotFound     = errors.New("parameter not found")
	ErrUnsupportedBinder     = errors.New("biinder not supported")
	ErrUnsupportedTransform  = errors.New("unsupported transform")

	ErrArticleNotFound = errors.New("article not found")
)
