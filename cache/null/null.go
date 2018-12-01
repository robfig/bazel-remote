package null

import (
	"errors"
	"io"
	"log"

	humanize "github.com/dustin/go-humanize"
)

type Cache struct {
	Max, Size int64
	numItems  int
	Log       *log.Logger
}

func (c *Cache) Put(key string, size int64, expectedSha256 string, r io.Reader) error {
	c.Size += size
	c.numItems++
	c.Log.Println(humanize.Bytes(uint64(c.Size)), humanize.Bytes(uint64(size)))
	return nil
}

func (*Cache) Get(key string, actionCache bool) (data io.ReadCloser, sizeBytes int64, err error) {
	return nil, 0, errors.New("nope")
}

func (*Cache) Contains(key string, actionCache bool) (ok bool) {
	return false
}

func (c *Cache) MaxSize() int64 {
	return c.Max
}

func (c *Cache) CurrentSize() int64 {
	return c.Size
}

func (c *Cache) NumItems() int {
	return c.numItems
}
