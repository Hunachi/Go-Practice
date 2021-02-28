// Implemented by Hunachi

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	pSize, err := rot.r.Read(p)
	if err != nil {
		return pSize, err
	}
	for i, v := range p {
		if (v - 0x41) < 26 && (v - 0x41) >= 0 {
			p[i] = ((v - 0x41 + 13) % 26 + 0x41)
		} else if (v - 0x61) < 26 && (v - 0x61) >= 0 {
			p[i] = ((v - 0x61 + 13) % 26 + 0x61)
		}
	}
	return pSize, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}