package simplematrix

import (
	"testing"
)

func Test_NewMatrix(t *testing.T) {
	_, err := NewMatrix(2, 2, 1, 2)
	if err == nil {
		t.Errorf("esperado erro por falta de elementos")
	}
}
