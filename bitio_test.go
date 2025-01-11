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
	bw := bitio.NewWriter(file)
	// 0101 0000 = 80
	bw.WriteZero()
	bw.WriteOne()
	bw.WriteZero()
	bw.WriteOne()
	err = bw.Close()
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
	bw := bitio.NewWriter(file)
	// 010 11010 = 0x5a = 90 = 'Z'
	err = bw.WriteBits(0x2, 3)
	err = bw.WriteBits(0x1a, 5)
	if err != nil {
		t.Fatal("could not write bits")
	}
	err = bw.Close()
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
	if buffer[0] != byte(90) {
		t.Fatalf("incorrect value read from file, %d\n", buffer[0])
	}
}

func testSmallMulti(t *testing.T) {
	file, err := os.Create("/tmp/test")
	if err != nil {
		t.Fatal("could not open testing file")
	}
	bw := bitio.NewWriter(file)
	// 0101 1010 = 0x5a = 90 = 'Z'
	err = bw.WriteBits(0x02, 3)
	err = bw.WriteBits(0x1a, 5)
	// 1000 0001 = 0x41 = 65 = 'A'
	err = bw.WriteBits(0x41, 8)
	if err != nil {
		t.Fatal("could not write bits")
	}
	err = bw.Close()
	if err != nil {
		t.Fatal("could not close bitio")
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		t.Fatal("could not seek to start of file")
	}
	buffer := make([]byte, 2)
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}
	if string(buffer) != "ZA" {
		t.Fatalf("incorrect value read from file, %d\n", buffer[0])
	}
}

func testLarge(t *testing.T) {
	file, err := os.Create("/tmp/test")
	if err != nil {
		t.Fatal("could not open testing file")
	}
	b := bitio.NewWriter(file)
	// 'KAZUYA'
	// 4B415A555941
	err = b.WriteBits(0x9682B4AAB28, 45)
	err = b.WriteBits(0x1, 3)
	if err != nil {
		t.Fatal("could not write bits")
	}
	err = b.Close()
	if err != nil {
		t.Fatal("could not close bitio")
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		t.Fatal("could not seek to start of file")
	}
	buffer := make([]byte, 6)
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}
	if string(buffer) != "KAZUYA" {
		t.Fatalf("incorrect value read from file, %d\n", buffer[0])
	}
}
