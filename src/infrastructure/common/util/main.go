package util

import "github.com/srcanton/go-gin-ddd-cqrs/src/infrastructure/common/util/error"

// Interface Util interface
type Interface interface {
}

// Util provide utilities
type Util struct {
	Error *error.Error
}

// Initialize initialize utilities
func Initialize() *Util {
	return &Util{Error: error.New()}
}
