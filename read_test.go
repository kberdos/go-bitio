package bitio_test

import (
	"os"
	"testing"

	"github.com/kberdos/go-bitio"
)

func TestRead(t *testing.T) {
	t.Run("read bit", testReadBit)
	t.Run("read bits", testReadBits)
	t.Run("read byte", testReadByte)
}

func testReadBits(t *testing.T) {
	file, err := os.Open("./testfiles/readtest")
	if err != nil {
		t.Fatal("error opening file", err)
	}
	bw, err := bitio.NewReader(file)
	if err != nil {
		t.Fatal("error creating reader", err)
	}
	// 'H' = 01001000
	x, err := bw.ReadBits(2) // 01
	if err != nil {
		t.Fatal("error reading from file", err)
	}
	if x != 1 {
		t.Fatalf("incorrect value read, received %d\n", x)
	}

	x, err = bw.ReadBits(4) // 0010
	if err != nil {
		t.Fatal("error reading from file", err)
	}
	if x != 2 {
		t.Fatalf("incorrect value read, received %d\n", x)
	}
}

func testReadBit(t *testing.T) {
	file, err := os.Open("./testfiles/readtest")
	if err != nil {
		t.Fatal("error opening file", err)
	}
	bw, err := bitio.NewReader(file)
	if err != nil {
		t.Fatal("error creating reader", err)
	}
	// 'H' = 01001000
	x, err := bw.ReadBits(1) // 0
	if err != nil {
		t.Fatal("error reading from file", err)
	}
	if x != 0 {
		t.Fatalf("incorrect value read, received %d\n", x)
	}
}

func testReadByte(t *testing.T) {
	file, err := os.Open("./testfiles/readtest")
	if err != nil {
		t.Fatal("error opening file", err)
	}
	bw, err := bitio.NewReader(file)
	if err != nil {
		t.Fatal("error creating reader", err)
	}
	x, err := bw.ReadBits(8)
	if err != nil {
		t.Fatal("error reading from file", err)
	}
	if x != 72 {
		t.Fatalf("incorrect value read, received %d\n", x)
	}
}
