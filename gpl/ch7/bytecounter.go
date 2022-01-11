package main

type ByteCounter int

type ReadWrite interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}
