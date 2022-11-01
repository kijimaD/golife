package main

import (
	"fmt"
	"testing"
)

func worldSeed() World {
	return World{
		Cell{1, false},
		Cell{2, true},
		Cell{3, false},

		Cell{4, true},
		Cell{5, true},
		Cell{6, true},

		Cell{7, false},
		Cell{8, true},
		Cell{9, false},
	}
}

func TestUpPos(t *testing.T) {
	// fmt.Println(worldSeed().UpPos(4))
	// fmt.Println(worldSeed()[1])

	// 通常
	if worldSeed().UpPos(4) != worldSeed()[1] {
		t.Errorf("is not match")
	}

	// 乗り越え
	if worldSeed().UpPos(1) != worldSeed()[7] {
		t.Errorf("is not match")
	}
}

func TestDownPos(t *testing.T) {
	// 通常
	if worldSeed().DownPos(1) != worldSeed()[4] {
		t.Errorf("is not match")
	}

	// 乗り越え
	if worldSeed().DownPos(7) != worldSeed()[1] {
		t.Errorf("is not match")
	}
}

// func TestLeftPos(t *testing.T) {
// 	// 通常
// 	fmt.Println(worldSeed().LeftPos(3))
// 	fmt.Println(worldSeed()[2])
// 	if worldSeed().LeftPos(3) == worldSeed()[2] {
// 		t.Errorf("is not match")
// 	}

// 	// 乗り越え
// 	// if worldSeed().LeftPos(4) == worldSeed()[6] {
// 	// 	t.Errorf("is not match")
// 	// }
// }

func TestIntToCord(t *testing.T) {
	w := worldSeed()
	fmt.Println(1, w.IntToCord(1))
	fmt.Println(2, w.IntToCord(2))
	fmt.Println(3, w.IntToCord(3))
	fmt.Println(4, w.IntToCord(4))
	fmt.Println(5, w.IntToCord(5))
	fmt.Println(6, w.IntToCord(6))
}
