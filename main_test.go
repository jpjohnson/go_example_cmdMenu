package main

import (
	"testing"
)

func TestMenu(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Fail()
		}
	}()
	menu()
}
func TestShowDate(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Fail()
		}
	}()
	showDate()
}
func TestRunCmdFail(t *testing.T) {
	t.Log(runCmd("foo"))
}
