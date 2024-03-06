package disk_test

import (
	"os"
	"testing"

	"github.com/shunta0213/sprDB/internal/storage/disk"
	"github.com/shunta0213/sprDB/internal/storage/page"
	"github.com/shunta0213/sprDB/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	m.Run()

	// remove test.db after all tests
	os.Remove("test.db")
}

func TestNewManager(t *testing.T) {
	t.Run("Test NewManager", func(t *testing.T) {
		m, err := disk.NewManager("test.db")

		a := assert.New(t)
		a.Nil(err)
		a.Equal(int64(page.PageSize), m.PageSize)
		a.Equal(int64(1), m.NextPageId)
		a.Equal(int64(0), m.NumPages)
	})
}

func TestClose(t *testing.T) {
	t.Run("Test Close", func(t *testing.T) {
		m, _ := disk.NewManager("test.db")
		err := m.Close()

		a := assert.New(t)
		a.Nil(err)
	})
}

func TestWritePage(t *testing.T) {
	t.Run("Test WritePage", func(t *testing.T) {
		m, _ := disk.NewManager("test.db")
		// write page
		contents := make([]byte, page.PageSize)
		err := m.WritePage(types.PageID(0), contents)

		a := assert.New(t)
		a.Nil(err)
	})

	t.Run("Test Invalid Page Size", func(t *testing.T) {
		m, _ := disk.NewManager("test.db")
		// write page
		contents := make([]byte, page.PageSize-1)
		err := m.WritePage(types.PageID(0), contents)

		a := assert.New(t)
		a.NotNil(err)
		t.Log(err)
	})
}

func TestReadPage(t *testing.T) {
	// Test fails if WritePage fails
	t.Run("Test ReadPage", func(t *testing.T) {
		m, _ := disk.NewManager("test.db")
		// write page
		contents := make([]byte, page.PageSize)
		err := m.WritePage(types.PageID(10), contents)
		if err != nil {
			t.Fatal(err)
		}

		buffer := make([]byte, page.PageSize)
		err = m.ReadPage(types.PageID(10), buffer)

		a := assert.New(t)
		a.Nil(err)
	})

	// Test multiple pages
	t.Run("Test ReadPage Multiple Pages", func(t *testing.T) {
		m, _ := disk.NewManager("test.db")
		// write multiple pages
		for i := 0; i < 10; i++ {
			contents := make([]byte, page.PageSize, page.PageSize)
			contents[0] = byte(i)
			err := m.WritePage(types.PageID(i), contents)
			if err != nil {
				t.Fatal(err)
			}
		}

		a := assert.New(t)

		for i := 0; i < 10; i++ {
			buffer := make([]byte, page.PageSize)
			err := m.ReadPage(types.PageID(i), buffer)
			a.Nil(err)
		}
	})
}
