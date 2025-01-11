package bitio_test

import (
	"io"
	"os"
	"testing"

	"github.com/kberdos/go-bitio"
)

func TestBitio(t *testing.T) {
	t.Run("basic write test", testBasicWrite)
}

func testBasicWrite(t *testing.T) {
	file, err := os.Create("/tmp/test")
	if err != nil {
		t.Fatal("could not open testing file")
	}
	b := bitio.New(file)
	// 0101 0000 = 80
	b.WriteZero()
	b.WriteOne()
	b.WriteZero()
	b.WriteOne()
	err = b.Close()
	if err != nil {
		t.Fatal("could not close bitio")
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		t.Fatal("could not seek to start of file")
	}
	buffer := make([]byte, 1)
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}
	if buffer[0] != byte(80) {
		t.Fatalf("incorrect value read from file, %d\n", buffer[0])
	}
}
