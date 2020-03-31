package bitcoindat

import (
	"bytes"
	"testing"
)

func TestProperCopy(t *testing.T) {
	// There was a bug where the passed byte slice was reversed when creating a
	// Hash256 datatype.

	b := []byte{0x20, 0x83, 0x59, 0x15, 0xf1, 0xdf, 0x04, 0xf8, 0x0e, 0x5f, 0xe1, 0x62, 0xc3, 0x8d, 0x6d, 0x89, 0x94, 0xd3, 0xee, 0x0b, 0x6b, 0x7d, 0x32, 0xa4, 0x49, 0xf9, 0x92, 0x5e, 0x4b, 0x0b, 0x72, 0xe2}
	h1 := NewHash256ByteSlice(b)
	h1Reversed := h1.ReversedCopy()
	h2 := h1Reversed.ReversedCopy()
	h1h2Equal := bytes.Compare(h1[:], h2[:]) == 0
	if !h1h2Equal {
		t.Errorf("Calling ReversedCopy() twice should ")
	}

}
