package main

import "testing"

func TestBoardInfoStruct(t *testing.T) {
	board := BoardInfo{
		Port: "/dev/ttyUSB0",
		FQBN: "arduino:avr:uno",
	}
	
	if board.Port != "/dev/ttyUSB0" {
		t.Errorf("Expected port /dev/ttyUSB0, got %s", board.Port)
	}
	
	if board.FQBN != "arduino:avr:uno" {
		t.Errorf("Expected FQBN arduino:avr:uno, got %s", board.FQBN)
	}
}

func TestDummy(t *testing.T) {
	if 1+1 != 2 {
		t.Error("math is broken")
	}
}
