package bitio_test

import (
	"io"
	"os"
	"testing"

	"github.com/kberdos/go-bitio"
)

func TestBitio(t *testing.T) {
	t.Run("basic write test", testBasicWrite)
	t.Run("write < 8 bits at once", testSmallAtOnce)
	t.Run("write < 8 bits, multi-byte ", testSmallMulti)
	t.Run("write > 8 bits", testLarge)
	// TODO: ignoring higher set bits
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

func testSmallAtOnce(t *testing.T) {
	file, err := os.Create("/tmp/test")
	if err != nil {
		t.Fatal("could not open testing file")
	}
	b := bitio.New(file)
	// 0101 1010 = 0x5a = 90 = 'Z'
	err = b.WriteBits(0x02, 3)
	err = b.WriteBits(0x1a, 5)
	if err != nil {
		t.Fatal("could not write bits")
	}
	err = b.Close()
	if err != nil {
		t.Fatal("could not close bitio")
	}
}

func testSmallMulti(t *testing.T) {
	file, err := os.Create("/tmp/test")
	if err != nil {
		t.Fatal("could not open testing file")
	}
	b := bitio.New(file)
	// 0101 1010 = 0x5a = 90 = 'Z'
	err = b.WriteBits(0x02, 3)
	err = b.WriteBits(0x1a, 5)
	// 1000 0001 = 0x41 = 65 = 'A'
	err = b.WriteBits(0x41, 8)
	if err != nil {
		t.Fatal("could not write bits")
	}
	err = b.Close()
	if err != nil {
		t.Fatal("could not close bitio")
	}
}

func testLarge(t *testing.T) {
	file, err := os.Create("/tmp/test")
	if err != nil {
		t.Fatal("could not open testing file")
	}
	b := bitio.New(file)
	// 4B 41 5A 55 59 41 = 'KAZUYA'
	err = b.WriteBits(0x4b415a555941, 48)
	if err != nil {
		t.Fatal("could not write bits")
	}
	err = b.Close()
	if err != nil {
		t.Fatal("could not close bitio")
	}
}
