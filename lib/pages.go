package pages

import (
	"github.com/Hatch1fy/errors"
	core "github.com/Hatch1fy/service-core"
)

const (
	// ErrEmptyKey is returned when an empty key is provided
	ErrEmptyKey = errors.Error("invalid key, cannot be empty")
	// ErrEmptyName is returned when an empty name is provided
	ErrEmptyName = errors.Error("invalid name, cannot be empty")
	// ErrPageWithNameExists is returned when a page name is used and already exists
	ErrPageWithNameExists = errors.Error("cannot create a page with the provided name, one already exists")
)

const relationshipKey = "key"

var relationships = []string{
	relationshipKey,
}

// New will return a new instance of pages
func New(dir string) (pp *Pages, err error) {
	var p Pages
	if p.c, err = core.New("pages", dir, &Entry{}, relationships...); err != nil {
		return
	}

	pp = &p
	return
}

// Pages manages the pages as a whole
type Pages struct {
	c *core.Core
}

func (p *Pages) getByKey(txn *core.Transaction, key string) (e *Entry, err error) {
	var entry Entry
	if err = txn.GetLastByRelationship(relationshipKey, key, &entry); err != nil {
		return
	}

	e = &entry
	return
}

func (p *Pages) getAllByKey(txn *core.Transaction, key string) (es []*Entry, err error) {
	err = txn.GetByRelationship(relationshipKey, key, &es)
	return
}

func (p *Pages) forEachByKey(txn *core.Transaction, key string, fn func(*Entry) error) (err error) {
	err = txn.ForEachRelationship(relationshipKey, key, func(id string, val core.Value) (err error) {
		return fn(val.(*Entry))
	})

	return
}

func (p *Pages) new(txn *core.Transaction, e *Entry) (err error) {
	if _, err = p.getByKey(txn, e.Key); err == nil {
		return ErrPageWithNameExists
	}

	_, err = txn.New(e)
	return
}

func (p *Pages) edit(txn *core.Transaction, key string, d Data) (err error) {
	var latest *Entry
	if latest, err = p.getByKey(txn, key); err != nil {
		return
	}

	e := newEntry(latest.Name, latest.Key, d)

	// Create a new entry using the provided
	_, err = txn.New(&e)
	return
}

func (p *Pages) remove(txn *core.Transaction, key string) (err error) {
	var count int
	if err = p.forEachByKey(txn, key, func(e *Entry) (err error) {
		count++
		return txn.Remove(e.ID)
	}); err != nil {
		return
	}

	if count == 0 {
		return core.ErrEntryNotFound
	}

	return
}

// New will insert a new page with a given page name
func (p *Pages) New(pagename string, d Data) (key string, err error) {
	e := newEntry(pagename, santizeName(pagename), d)
	if err = e.Validate(); err != nil {
		return
	}

	if err = p.c.Transaction(func(txn *core.Transaction) (err error) {
		return p.new(txn, &e)
	}); err != nil {
		return
	}

	key = e.Key
	return
}

// Get will retrieve an entry by key
func (p *Pages) Get(key string) (e *Entry, err error) {
	err = p.c.Transaction(func(txn *core.Transaction) (err error) {
		e, err = p.getByKey(txn, key)
		return
	})

	return
}

// GetAll will retrieve all entries by key
func (p *Pages) GetAll(key string) (es []*Entry, err error) {
	err = p.c.Transaction(func(txn *core.Transaction) (err error) {
		es, err = p.getAllByKey(txn, key)
		return
	})

	return
}

// Edit will update an existing page with a given key
func (p *Pages) Edit(key string, d Data) (err error) {
	err = p.c.Transaction(func(txn *core.Transaction) (err error) {
		return p.edit(txn, key, d)
	})

	return
}

// Remove will delete an existing page with a given key
func (p *Pages) Remove(key string) (err error) {
	err = p.c.Transaction(func(txn *core.Transaction) (err error) {
		return p.remove(txn, key)
	})

	return
}

// Close will close an instance of pages
func (p *Pages) Close() error {
	return p.c.Close()
}
