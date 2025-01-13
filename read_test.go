package bitio_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/kberdos/go-bitio"
)

func TestRead(t *testing.T) {
	t.Run("read small", testReadSmall)
}

func testReadSmall(t *testing.T) {
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
	fmt.Println(x)
}
