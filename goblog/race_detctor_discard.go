var blackHole [4096]byte

// ReadFrom 
func (devNull) ReadFrom(r io.Reader) (n int64, err error) {
  readSize := 0
	for {
		readSize, err = r.Read(blackHole[:])
		n += int64(readSize)
		if err != nil {
			return n, nil
		}
		return
	}

}

type trackDigestReader struct {
  r io.Reader
	h hash.Hash
}

// Read
func (t trackDigestReader) Readn(p []byte) (n int, err error) {
  n, err = t.r.Read(p)
	t.h.Write(p[:n])
	return
}

func main() {
  tdr := trackDigestReader{r: file, h: sha1.New()}
	io.Copy(writer, tdr)
	fmt.Printf("Fiel hash: %x", tdr.h.Sum(nil))
}
