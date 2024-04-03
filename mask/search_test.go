package mask

import "testing"

func TestSearch(t *testing.T) {
	expected := 30
	actual := Search([]byte("Here's my spammy page: http://hehefouls.netHAHAHA see you."), []byte("http://"))
	if expected != actual {
		t.Errorf("Expected %d do not match actual %d", expected, actual)
	}
}
