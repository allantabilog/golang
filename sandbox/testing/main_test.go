package main

import (
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		desc      string
		name      string
		want      string
		expectErr bool
	}{
		{
			desc:      "Error - empty name",
			expectErr: true,
			name:      "",
		},
		{
			desc:      "Success",
			name:      "John",
			want:      "Hello John",
			expectErr: false,
		},
	}

	for _, test := range tests {
		got, err := Greet(test.name)
		switch {
		case err == nil && test.expectErr:
			t.Errorf("TestGreet(%s) got err == nil, want err != nil", test.desc)
			continue
		case err != nil && !test.expectErr:
			t.Errorf("TestGreet(%s) got err == %s, want err == nil", test.desc, err)
			continue
		case err != nil:
			continue
		}

		if got != test.want {
			t.Errorf("TestGreet(%s): got result %q, want %q", test.desc, got, test.want)
		}
	}
}
