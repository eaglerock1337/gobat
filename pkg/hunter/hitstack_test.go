package hunter

import (
	"testing"

	"github.com/eaglerock1337/go/battleship/pkg/board"
)

var TestHitStacks = [5][]board.Square{
	{{Letter: 0, Number: 1}, {Letter: 5, Number: 5}, {Letter: 2, Number: 3}, {Letter: 0, Number: 0}, {Letter: 4, Number: 4}},
	{{Letter: 3, Number: 4}, {Letter: 4, Number: 4}, {Letter: 5, Number: 4}, {Letter: 5, Number: 5}},
	{{Letter: 7, Number: 5}, {Letter: 3, Number: 7}},
	{{Letter: 1, Number: 1}},
	{{Letter: 0, Number: 0}, {Letter: 1, Number: 3}, {Letter: 7, Number: 3}, {Letter: 8, Number: 0}},
}

var AbsentSquares = [5]board.Square{
	{Letter: 0, Number: 3},
	{Letter: 3, Number: 3},
	{Letter: 4, Number: 9},
	{Letter: 0, Number: 8},
	{Letter: 9, Number: 1},
}

func TestPush(t *testing.T) {
	var expected = [5]HitStack{
		[]board.Square{{Letter: 0, Number: 1}, {Letter: 5, Number: 5}, {Letter: 2, Number: 3}, {Letter: 0, Number: 0}, {Letter: 4, Number: 4}, {Letter: 0, Number: 3}},
		[]board.Square{{Letter: 3, Number: 4}, {Letter: 4, Number: 4}, {Letter: 5, Number: 4}, {Letter: 5, Number: 5}, {Letter: 3, Number: 3}},
		[]board.Square{{Letter: 7, Number: 5}, {Letter: 3, Number: 7}, {Letter: 4, Number: 9}},
		[]board.Square{{Letter: 1, Number: 1}, {Letter: 0, Number: 8}},
		[]board.Square{{Letter: 0, Number: 0}, {Letter: 1, Number: 3}, {Letter: 7, Number: 3}, {Letter: 8, Number: 0}, {Letter: 9, Number: 1}},
	}

	for i, data := range TestHitStacks {
		stack := HitStack(data)
		err := stack.Push(AbsentSquares[i])
		if err != nil {
			t.Errorf("Push returned unexpected error for test %v: %v", i, err)
		}

		if len(expected[i]) != len(stack) {
			t.Errorf("Push was incorrect for test %v, got: %v, want: %v", i, stack, expected[i])
		}

		for j, square := range expected[i] {
			if square.Letter != stack[j].Letter || square.Number != stack[j].Number {
				t.Errorf("Push was incorrect for test %v, got: %v, want: %v", i, square, stack[j])
			}
		}
	}
}

var PresentSquares = [5]board.Square{
	{Letter: 2, Number: 3},
	{Letter: 3, Number: 4},
	{Letter: 3, Number: 7},
	{Letter: 1, Number: 1},
	{Letter: 7, Number: 3},
}

func TestBadPush(t *testing.T) {
	for i, data := range TestHitStacks {
		stack := HitStack(data)
		err := stack.Push(PresentSquares[i])

		if err == nil {
			t.Errorf("Push did not error as expected with %v, returned HitStack: %v", PresentSquares[i], stack)
		}
	}
}

func TestPop(t *testing.T) {
	var expected = [5]HitStack{
		[]board.Square{{Letter: 0, Number: 1}, {Letter: 5, Number: 5}, {Letter: 4, Number: 4}, {Letter: 0, Number: 0}},
		[]board.Square{{Letter: 5, Number: 5}, {Letter: 4, Number: 4}, {Letter: 5, Number: 4}},
		[]board.Square{{Letter: 7, Number: 5}},
		[]board.Square{},
		[]board.Square{{Letter: 0, Number: 0}, {Letter: 1, Number: 3}, {Letter: 8, Number: 0}},
	}

	for i, data := range TestHitStacks {
		stack := HitStack(data)
		err := stack.Pop(PresentSquares[i])
		if err != nil {
			t.Errorf("Pop returned unexpected error for test %v: %v", i, err)
		}

		if len(expected[i]) != len(stack) {
			t.Errorf("Pop was incorrect for test %v, got: %v, want: %v", i, stack, expected[i])
		}

		for j, square := range expected[i] {
			if square.Letter != stack[j].Letter || square.Number != stack[j].Number {
				t.Errorf("Pop was incorrect for test %v, got: %v, want: %v", i, square, stack[j])
			}
		}
	}
}

func TestBadPop(t *testing.T) {
	for i, data := range TestHitStacks {
		stack := HitStack(data)
		err := stack.Pop(AbsentSquares[i])

		if err == nil {
			t.Errorf("Pop did not error as expected with %v, returned HitStack: %v", AbsentSquares[i], stack)
		}
	}
}

var MoreTestHitStacks = [5][]board.Square{
	{{Letter: 0, Number: 1}, {Letter: 5, Number: 5}, {Letter: 2, Number: 3}, {Letter: 0, Number: 0}, {Letter: 4, Number: 4}},
	{{Letter: 3, Number: 4}, {Letter: 4, Number: 4}, {Letter: 5, Number: 4}, {Letter: 5, Number: 5}},
	{{Letter: 7, Number: 5}, {Letter: 3, Number: 7}},
	{{Letter: 1, Number: 1}},
	{{Letter: 0, Number: 0}, {Letter: 1, Number: 3}, {Letter: 7, Number: 3}, {Letter: 8, Number: 0}},
}

func TestInStack(t *testing.T) {
	for i, data := range MoreTestHitStacks {
		stack := HitStack(data)
		trueTest := stack.InStack(PresentSquares[i])
		falseTest := stack.InStack(AbsentSquares[i])

		if !trueTest {
			t.Errorf("InStack returned false for %v in %v", PresentSquares[i], stack)
		}
		if falseTest {
			t.Errorf("InStack returned true for %v in %v", AbsentSquares[i], stack)
		}
	}
}
