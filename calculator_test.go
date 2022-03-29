package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	get := add(1, 2)
	res := 3
	if get != res {
		t.Errorf("Add 1 and 2 =3; got %d", get)
	}
	get = subtract(1, 2)
	res = -1
	if get != res {
		t.Errorf("Add 1 and 2 =3; got %d", get)
	}
	get = multiply(1, 2)
	res = 2
	if get != res {
		t.Errorf("Add 1 and 2 =3; got %d", get)
	}
	get = divide(10, 2)
	res = 5
	if get != res {
		t.Errorf("Add 1 and 2 =3; got %d", get)
	}

}
