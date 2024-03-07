package tuple_test

import (
	"testing"

	"github.com/shunta0213/sprDB/internal/storage/tuple"
	"github.com/stretchr/testify/assert"
)

func TestNewTuple(t *testing.T) {
	rid := tuple.RID{PageId: 1, SlotNo: 1}
	data := []byte("test")
	tuple := tuple.NewTuple(rid, data)

	a := assert.New(t)
	a.Equal(rid, tuple.GetRID())
	a.Equal(data, tuple.GetData())
}
