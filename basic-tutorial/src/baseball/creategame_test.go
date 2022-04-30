package baseball

import (
	"fmt"
	"testing"
)

type Game interface {
}

type gameNumbers struct {
	numbers interface{}
}

func CreateGame(nums string) (Game, error) {
	//return &gameNumbers{numbers: nums}, nil
	return nil, fmt.Errorf("invalid nums: %s", nums)
}

func TestInvalidNums(t *testing.T) {
	_, err := CreateGame("01")
	if err == nil {
		t.Fatalf("error must be returnd: %s", "01")
	}

	_, err2 := CreateGame("0112")
	if err2 == nil {
		t.Fatalf("error must be returnd: %s", "0112")
	}

	_, err3 := CreateGame("01417")
	if err3 == nil {
		t.Fatalf("error must be returnd: %s", "01417")
	}
}
