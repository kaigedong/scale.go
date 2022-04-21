package types_test

import (
	"github.com/kaigedong/scale.go/types"
	"testing"
)

func TestRegCustomTypes(t *testing.T) {
	types.RuntimeType{}.Reg()
}
