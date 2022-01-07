package Factory

import "testing"

func TestNewCompany(t *testing.T) {
	NewCompany("b").Service()
	NewCompany("t").Service()
}