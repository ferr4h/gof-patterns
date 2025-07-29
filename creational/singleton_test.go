package creational

import "testing"

func TestSingleton(t *testing.T) {
	s1 := GetInstance()
	s2 := GetInstance()

	if s1 != s2 {
		t.Error("Expected s1 and s2 to be the same")
	}
}
