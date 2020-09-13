// Test file for even package
package even

import "testing"

// Even returns true if i is even, else falise is returned
func TestEven(t *testing.T) {
	if !Even(2) {
		t.Log("2 should be even!")
		t.Fail()
	}
	if !Even(3) {
		t.Log("3 should be odd!")
		t.Fail()
	}
}
