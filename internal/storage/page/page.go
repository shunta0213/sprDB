// Page is the basic unit of storage in the database.
// The size of a page is 8KiB.
package page

// page size 8KiB
const PageSize = 8192

type Page struct {
	id       int32           // identifier
	pinCount uint32          // ?: counts how many tx are accessing
	isDirty  bool            // the page is modified but not persistent
	data     *[PageSize]byte // bytes stored on disk
}

// Create a new Page
// pinCount is 0, isDirty is false for default
func NewPage(id int32, data *[PageSize]byte) *Page {
	return &Page{id: id, pinCount: 0, isDirty: false, data: data}
}

func (p *Page) Id() int32 {
	return p.id
}

func (p *Page) PinCount() uint32 {
	return p.pinCount
}

func (p *Page) IsDirty() bool {
	return p.isDirty
}

func (p *Page) Data() *[PageSize]byte {
	return p.data
}

// Increase the pin count of the Page
func (p *Page) IncPinCount() {
	p.pinCount++
}

// Decrease the pin count of the Page
func (p *Page) DecPinCount() {
	p.pinCount--
}

// Set the page as dirty
func (p *Page) SetDirty() {
	p.isDirty = true
}

// Set the page as clean
func (p *Page) SetClean() {
	p.isDirty = false
}
