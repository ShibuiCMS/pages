package pages

import "testing"

func TestEntry_Validate(t *testing.T) {
	type tc struct {
		e   Entry
		err error
	}

	tcs := []tc{
		{
			e:   newEntry("test_name", "test-name", "test-template", nil),
			err: nil,
		},
		{
			e:   newEntry("", "test-name", "test-template", nil),
			err: ErrEmptyName,
		},
		{
			e:   newEntry("test_name", "", "test-template", nil),
			err: ErrEmptyKey,
		},
		{
			e:   newEntry("test_name", "test-name", "", nil),
			err: ErrEmptyTemplate,
		},
	}

	for _, tc := range tcs {
		if err := tc.e.Validate(); err != tc.err {
			t.Fatalf("invalid error, expected %v and received %v", tc.err, err)
		}
	}
}
