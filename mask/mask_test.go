package mask

import "testing"

func TestMask(t *testing.T) {
	expected := "Here's my spammy page: http://******************* see you."
	actual := Mask(30, []byte("Here's my spammy page: http://hehefouls.netHAHAHA see you."), '*')
	if expected != actual {
		t.Errorf("Expected %s do not match actual %s", expected, actual)
	}
}
