package simplelist

import "testing"

func Test_Append(t *testing.T) {
	a1 := List{1, 2, 3}
	a2 := List{}
	a2.Append(1)
	a2.Append(2, 3)
	for i, v := range a1 {
		if a2[i] != v {
			t.Errorf("o valor de índice %d não corresponde a %f", i, v)
		}
	}
}

func Test_Add(t *testing.T) {
	a1 := List{1, 2, 3}
	a2 := List{4, 5, 6}

	a3, _ := a1.Add(a2)
	if !(a3[0] == 5 && a3[1] == 7 && a3[2] == 9) {
		t.Errorf("valor esperado era List[5,7,9] retornou %v", a3)
	}

	a4 := List{4, 5}
	_, err := a1.Add(a4)
	if err == nil {
		t.Error("esperado err != nil")
	}

}
