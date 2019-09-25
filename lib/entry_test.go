package pages

import "testing"

func TestEntry_Validate(t *testing.T) {
	type tc struct {
		e   Entry
		err error
	}

	tcs := []tc{
		{
			e:   newEntry("test_name", "test-name", nil),
			err: nil,
		},
		{
			e:   newEntry("", "test-name", nil),
			err: ErrEmptyName,
		},
		{
			e:   newEntry("test_name", "", nil),
			err: ErrEmptyKey,
		},
	}

	for _, tc := range tcs {
		if err := tc.e.Validate(); err != tc.err {
			t.Fatalf("invalid error, expected %v and received %v", tc.err, err)
		}
	}
}
