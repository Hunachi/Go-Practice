// Implemented by Hunachi

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(bytes []byte) (int, error) {
	for i, _ := range bytes {
		bytes[i] = byte('A')
	}
	return len(bytes), nil
}

func main() {
	reader.Validate(MyReader{})
}