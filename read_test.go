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
	t.Run("read bytes", testReadBytes)
	// TODO: reading past the end errors
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

func testReadBytes(t *testing.T) {
	file, err := os.Open("./testfiles/readtest_big")
	if err != nil {
		t.Fatal("error opening file", err)
	}
	bw, err := bitio.NewReader(file)
	if err != nil {
		t.Fatal("error creating reader", err)
	}
	// 'HELLO' = 72 69 76 76 79
	// = 0100 1000 0100 0101 0100 1100 0100 1100 0100 1111
	x, err := bw.ReadBits(15) // 0100 1000 0100 010
	if err != nil {
		t.Fatal("error reading from file", err)
	}
	if x != 9250 {
		t.Fatalf("incorrect value read, received %d\n", x)
	}
}
