package modlib

import(
	"testing"
)

func TestMessage(t *testing.T)  {
	msg := Message()
	t.Log(msg)
}

func TestMaxIntegerDigits(t *testing.T){
	MaxIntegerDigits()
}