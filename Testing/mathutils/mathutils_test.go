package mathutils

import "testing"

func TestFactorial(t *testing.T) {
	t.Run("factorial of 0", func(t *testing.T) {
		want := 1
		got, err := Factorial(0)
		if err != nil {
			t.Errorf("test failed withr reason:%v", err.Error())
		}
		if want != got {
			t.Errorf("wanted %v, got %v", want, got)
		}
	})

	t.Run("factorial of negative number", func(t *testing.T) {
		_,err:=Factorial(-5)
		if err==nil{
			t.Error("expected error but got none")
		}else if err.Error()!="cannot calculate factorial of -ve number"{
			t.Error("got error but not the required one")
		}
	})

	t.Run("factorial of a positive number", func(t *testing.T) {
		want := 120
		got, err := Factorial(5)
		if err != nil {
			t.Errorf("test failed withr reason:%v", err.Error())
		}
		if want != got {
			t.Errorf("wanted %v, got %v", want, got)
		}
	})

	
	t.Run("factorial of a number>10", func(t *testing.T) {
		want := 1307674368000
		got, err := Factorial(15)
		if err != nil {
			t.Errorf("test failed withr reason:%v", err.Error())
		}
		if want != got {
			t.Errorf("wanted %v, got %v", want, got)
		}
	})

}
