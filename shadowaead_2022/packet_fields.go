package shadowaead_2022

import (
	"encoding/binary"
	"io"

	"github.com/metacubex/sing/common/buf"
)

// The packet path reads and writes its header fields with these instead of
// binary.Read/Write, which go through reflection and an interface and allocate
// on every datagram (fork rityak/sing-shadowsocks, branch miray). Same errors: io.EOF on an
// empty buffer, io.ErrUnexpectedEOF on a short one, io.ErrShortBuffer when full.

func readUint64(b *buf.Buffer) (uint64, error) {
	if err := short(b, 8); err != nil {
		return 0, err
	}
	v := binary.BigEndian.Uint64(b.Bytes())
	b.Advance(8)
	return v, nil
}

func readUint16(b *buf.Buffer) (uint16, error) {
	if err := short(b, 2); err != nil {
		return 0, err
	}
	v := binary.BigEndian.Uint16(b.Bytes())
	b.Advance(2)
	return v, nil
}

func short(b *buf.Buffer, n int) error {
	switch {
	case b.Len() == 0:
		return io.EOF
	case b.Len() < n:
		return io.ErrUnexpectedEOF
	}
	return nil
}

func writeUint64(b *buf.Buffer, v uint64) error {
	if b.FreeLen() < 8 {
		return io.ErrShortBuffer
	}
	binary.BigEndian.PutUint64(b.Extend(8), v)
	return nil
}

func writeUint16(b *buf.Buffer, v uint16) error {
	if b.FreeLen() < 2 {
		return io.ErrShortBuffer
	}
	binary.BigEndian.PutUint16(b.Extend(2), v)
	return nil
}
