package service

import "testing"

func TestRead(t *testing.T) {
	m := new(mockString)
	m.On("stringRead", "data.txt").Return("Here's my spammy page: http://hehefouls.netHAHAHA see you.")
	str := Read("data.txt", m)
	if str != "Here's my spammy page: http://hehefouls.netHAHAHA see you." {
		t.Errorf("%v not match", str)
	}
}

func TestWrite(t *testing.T) {
	m := new(mockString)
	m.On("stringWrite", "data.txt,").Return("Here's my spammy page: http://******************* see you.")
	str := Write("data.txt", "Here's my spammy page: http://******************* see you.", m)
	if str != "Here's my spammy page: http://hehefouls.netHAHAHA see you.Here's my spammy page: http://******************* see you." {
		t.Errorf("%v not match", str)
	}
}
