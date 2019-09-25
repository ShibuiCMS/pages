package pages

import (
	"fmt"
	"os"
	"testing"

	"github.com/Hatch1fy/errors"
	core "github.com/Hatch1fy/service-core"
)

func TestNew(t *testing.T) {
	var (
		p   *Pages
		err error
	)

	if p, err = testInit(); err != nil {
		t.Fatal(err)
	}
	defer testTeardown(p, nil)

	return
}

func TestPages_New(t *testing.T) {
	var (
		p   *Pages
		err error
	)

	if p, err = testInit(); err != nil {
		t.Fatal(err)
	}
	defer testTeardown(p, nil)

	type tc struct {
		name string
		err  error
	}

	tcs := []tc{
		{
			name: "test page name",
			err:  nil,
		},
		{
			name: "",
			err:  ErrEmptyName,
		},
		{
			name: "(>'.')>",
			err:  ErrEmptyKey,
		},
	}

	for _, tc := range tcs {
		if _, err = p.New(tc.name, Data{"foo": "bar"}); err != tc.err {
			t.Fatalf("invalid error, expected %v and received %v", tc.err, err)
		}
	}
}

func TestPages_Get(t *testing.T) {
	var (
		p   *Pages
		err error
	)

	if p, err = testInit(); err != nil {
		t.Fatal(err)
	}
	defer testTeardown(p, nil)

	type tc struct {
		name string
		key  string
		data Data
		err  error
	}

	tcs := []tc{
		{
			name: "test page name",
			key:  "test-page-name",
			data: Data{"foo": "bar"},
			err:  nil,
		},
		{
			name: "test page name 2",
			key:  "test-page-name-2",
			data: Data{"foo": "bar 2"},
			err:  nil,
		},
		{
			name: "test page name 3",
			key:  "test-page-name-3",
			data: Data{"foo": "bar 3"},
			err:  nil,
		},
		{
			name: "test page name 4",
			key:  "test-page-name-4",
			data: Data{"foo": "bar 4"},
			err:  nil,
		},
		{
			name: "",
			key:  "test-page-name-5",
			data: nil,
			err:  core.ErrEntryNotFound,
		},
	}

	for _, tc := range tcs {
		if tc.name != "" {
			var key string
			if key, err = p.New(tc.name, tc.data); err != nil {
				t.Fatal(err)
			}

			if tc.key != key {
				t.Fatalf("invalid key, expected \"%s\" and received \"%s\"", tc.key, key)
			}
		}

		var e *Entry
		if e, err = p.Get(tc.key); err != tc.err {
			t.Fatalf("invalid error encountered while getting entry with key of \"%s\", expected %v and received %v", tc.key, tc.err, err)
		}

		if e == nil {
			continue
		}

		if e.Data["foo"] != tc.data["foo"] {
			t.Fatalf("invalid value for data[%s], expected \"%s\" and received \"%s\"", "foo", tc.data["foo"], e.Data["foo"])
		}
	}
}

func TestPages_GetAll(t *testing.T) {
	var (
		p   *Pages
		err error
	)

	if p, err = testInit(); err != nil {
		t.Fatal(err)
	}
	defer testTeardown(p, nil)

	type tc struct {
		name    string
		data    Data
		newData Data
		err     error
	}

	tcs := []tc{
		{
			name:    "test page name",
			data:    Data{"foo": "bar"},
			newData: Data{"foo": "baz"},
			err:     nil,
		},
		{
			name:    "another test page name",
			data:    Data{"one": "two"},
			newData: Data{"three": "four"},
			err:     nil,
		},
	}

	for _, tc := range tcs {
		var key string
		if key, err = p.New(tc.name, tc.data); err != tc.err {
			t.Fatalf("invalid error, expected %v and received %v", tc.err, err)
		}

		if err = p.Edit(key, tc.newData); err != nil {
			t.Fatalf("error editing \"%s\": %v", key, err)
		}

		var es []*Entry
		if es, err = p.GetAll(key); err != nil {
			t.Fatalf("error getting entry with key of \"%s\": %v", key, err)
		}

		if len(es) != 2 {
			t.Fatalf("invalid number of entries, expected %d and received %d", 2, len(es))
		}

		if e := es[0]; e.Data["foo"] != tc.data["foo"] {
			t.Fatalf("invalid value for data[%s], expected \"%s\" and received \"%s\"", "foo", tc.data["foo"], e.Data["foo"])
		}

		if e := es[1]; e.Data["foo"] != tc.newData["foo"] {
			t.Fatalf("invalid value for data[%s], expected \"%s\" and received \"%s\"", "foo", tc.newData["foo"], e.Data["foo"])
		}
	}
}

func TestPages_Edit(t *testing.T) {
	var (
		p   *Pages
		err error
	)

	if p, err = testInit(); err != nil {
		t.Fatal(err)
	}
	defer testTeardown(p, nil)

	type tc struct {
		name    string
		data    Data
		newData Data
		err     error
	}

	tcs := []tc{
		{
			name:    "test page name",
			data:    Data{"foo": "bar"},
			newData: Data{"foo": "baz"},
			err:     nil,
		},
	}

	for _, tc := range tcs {
		var key string
		if key, err = p.New(tc.name, tc.data); err != tc.err {
			t.Fatalf("invalid error, expected %v and received %v", tc.err, err)
		}

		if err = p.Edit(key, tc.newData); err != nil {
			t.Fatalf("error editing \"%s\": %v", key, err)
		}

		var e *Entry
		if e, err = p.Get(key); err != nil {
			t.Fatalf("error getting entry with key of \"%s\": %v", key, err)
		}

		if e.Data["foo"] != tc.newData["foo"] {
			t.Fatalf("invalid value for data[%s], expected \"%s\" and received \"%s\"", "foo", tc.newData["foo"], e.Data["foo"])
		}
	}
}

func TestPages_Remove(t *testing.T) {
	var (
		p   *Pages
		err error
	)

	if p, err = testInit(); err != nil {
		t.Fatal(err)
	}
	defer testTeardown(p, nil)

	type tc struct {
		name string
		key  string
		err  error
	}

	tcs := []tc{
		{
			name: "test page name 1",
			err:  nil,
		},
		{
			name: "test page name 2",
			err:  nil,
		},
		{
			name: "test page name 3",
			err:  nil,
		},
		{
			name: "test page name 4",
			err:  nil,
		},
		{
			name: "",
			key:  santizeName("test page name 1"),
			err:  core.ErrEntryNotFound,
		},
	}

	for _, tc := range tcs {
		if tc.key == "" {
			if tc.key, err = p.New(tc.name, Data{"foo": "bar"}); err != nil {
				t.Fatal(err)
			}

		}

		if err = p.Remove(tc.key); err != tc.err {
			t.Fatalf("invalid error while removing \"%s\", expected %v and recieved %v", tc.key, tc.err, err)
		}

		if err = p.Remove(tc.key); err != core.ErrEntryNotFound {
			t.Fatalf("invalid error, expected %v and received %v", core.ErrEntryNotFound, err)
		}
	}
}

func TestPages_Close(t *testing.T) {
	var (
		p   *Pages
		err error
	)

	if p, err = testInit(); err != nil {
		t.Fatal(err)
	}

	if err = testTeardown(p, nil); err != nil {
		t.Fatal(err)
	}

	if err = p.Close(); err != errors.ErrIsClosed {
		t.Fatalf("invalid error, expected %v and received %v", errors.ErrIsClosed, err)
	}
}

func testInit() (p *Pages, err error) {
	if err = os.Mkdir("./test_data", 0744); err != nil {
		return
	}

	return New("./test_data")
}

func testTeardown(p *Pages, expectedErr error) (err error) {
	var errs errors.ErrorList
	errs.Push(p.Close())
	errs.Push(os.RemoveAll("./test_data"))

	if err = errs.Err(); err != expectedErr {
		err = fmt.Errorf("invalid error, expected %v and received %v", expectedErr, err)
		return
	}

	return nil
}
