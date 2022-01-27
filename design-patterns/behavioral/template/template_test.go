package template

import "testing"

func TestCoder(t *testing.T) {
	xss := Coder{}
	w := NewWorker(&xss)
	w.Daily()
	w.GetUp()
	w.Work()
	w.Sleep()
}
