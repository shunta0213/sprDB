// Disk package is responsible for managing the disk file
package disk

import (
	"errors"
	"fmt"
	"os"

	"github.com/shunta0213/sprDB/internal/storage/page"
	"github.com/shunta0213/sprDB/internal/types"
)

type Manager struct {
	DB         *os.File
	PageSize   int64
	NextPageId int64
	NumPages   int64
}

func NewManager(fileName string) (*Manager, error) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileSize := fileInfo.Size()
	pageSize := page.PageSize
	numPages := fileSize / pageSize

	nextPageId := numPages + 1

	return &Manager{
		DB:         file,
		PageSize:   pageSize,
		NextPageId: nextPageId,
		NumPages:   numPages,
	}, nil
}

func (m *Manager) Close() error {
	return m.DB.Close()
}

func (m *Manager) WritePage(pageId types.PageID, data []byte) error {
	if len(data) != int(page.PageSize) {
		return errors.New(fmt.Sprintf("len(data): %d, pageSize: %d", len(data), page.PageSize))
	}

	offset := int64(pageId) * m.PageSize
	bytesWritten, err := m.DB.WriteAt(data, offset)
	if err != nil {
		return err
	}

	if bytesWritten != int(page.PageSize) {
		return errors.New(fmt.Sprintf("bytesWritten: %d, pageSize: %d", bytesWritten, page.PageSize))
	}

	err = m.DB.Sync()
	return err
}

func (m *Manager) ReadPage(pageId types.PageID, data []byte) error {
	offset := int64(pageId) * m.PageSize
	_, err := m.DB.ReadAt(data, offset)

	return err
}
