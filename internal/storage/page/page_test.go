package page_test

import (
	"testing"

	"github.com/shunta0213/sprDB/internal/storage/page"
	"github.com/shunta0213/sprDB/internal/types"
	"github.com/stretchr/testify/assert"
)

// Test NewPage
func TestNewPage(t *testing.T) {
	t.Run("Test NewPage", func(t *testing.T) {
		p := page.NewPage(10, &[page.PageSize]byte{'a', 'b', 'c'})

		a := assert.New(t)
		// type check
		a.IsType(p.Id(), types.PageID(0))
		a.IsType(p.PinCount(), uint32(0))
		a.IsType(p.IsDirty(), bool(false))
		a.IsType(p.Data(), &[page.PageSize]byte{})

		// default value check
		a.Equal(p.Id(), types.PageID(10))
		a.Equal(p.IsDirty(), false)
		a.Equal(p.PinCount(), uint32(0))
		a.Equal(p.Data(), &[page.PageSize]byte{'a', 'b', 'c'})
	})

	t.Run("Test Inc/DecPinCount", func(t *testing.T) {
		p := page.NewPage(10, &[page.PageSize]byte{'a', 'b', 'c'})

		a := assert.New(t)

		a.Equal(p.PinCount(), uint32(0))

		p.IncPinCount()
		a.Equal(p.PinCount(), uint32(1))

		p.DecPinCount()
		a.Equal(p.PinCount(), uint32(0))
	})

	t.Run("Test SetDirty/SetClean", func(t *testing.T) {
		p := page.NewPage(10, &[page.PageSize]byte{'a', 'b', 'c'})

		a := assert.New(t)

		a.Equal(p.IsDirty(), false)

		p.SetDirty()
		a.Equal(p.IsDirty(), true)

		p.SetClean()
		a.Equal(p.IsDirty(), false)
	})

}
