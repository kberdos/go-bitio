package bitio

const (
	READ_BUF_SIZE = 8  // 8 bytes in a uint64
	MAX_READ_SIZE = 64 // read a uint64 at msot
	W_START_POS   = 7  // use 8-bit cache for writing
	R_START_POS   = 63 // use 64-bit cache for reading
)
