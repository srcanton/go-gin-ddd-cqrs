package util_test

import (
	"github.com/srcanton/go-gin-ddd-cqrs/src/infrastructure/common/util"
	"testing"
)

// TestInitialize test InitizlizeUtil method
func TestInitializeUtil(t *testing.T) {
	instance := util.Initialize()
	if instance == nil {
		t.Error("Can not create util instance")
	}
}
