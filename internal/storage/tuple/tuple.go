package tuple

import "github.com/shunta0213/sprDB/internal/types"

type RID struct {
	PageId types.PageID
	SlotNo int
}

type Tuple struct {
	rid  RID
	data []byte
}

func NewTuple(rid RID, data []byte) *Tuple {
	return &Tuple{rid, data}
}

func (t *Tuple) GetRID() RID {
	return t.rid
}

func (t *Tuple) GetData() []byte {
	return t.data
}
