package statistics

import "testing"

func Test_Max(t *testing.T) {
	res := Max(1, 2, 3, 4, 5)
	if res != 5 {
		t.Errorf("resultado esperado erá 5, retornou %f", res)
	}
}
