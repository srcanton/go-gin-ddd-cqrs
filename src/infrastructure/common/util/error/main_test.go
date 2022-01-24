package error_test

import (
	"github.com/srcanton/go-gin-ddd-cqrs/src/infrastructure/common/util/error"
	"testing"
)

// TestNew test new method in error package
func TestNew(t *testing.T) {
	instance := &error.Error{}
	if instance == nil {
		t.Error("Can not create error instance")
	}
}
