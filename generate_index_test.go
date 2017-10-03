package main

import (
	"bytes"
	"testing"
)

func TestGenerateIndex(t *testing.T) {
	testCases := []struct {
		domain      string
		r           []repository
		expectedOut string
		expectedErr error
	}{
		{
			domain: "example.com",
			r: []repository{
				{
					Prefix: "pkg1",
					Subs:   []string{"subpkg1", "subpkg2"},
				},
				{
					Prefix: "pkg2",
					Subs:   []string{"subpkg1", "subpkg2/subsubpkg1"},
				},
			},
			expectedOut: `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<style>
* { font-family: sans-serif; }
</style>
</head>
<body>
<ul>

<li><a href="//example.com/pkg1">example.com/pkg1</a></li>

<li><a href="//example.com/pkg1/subpkg1">example.com/pkg1/subpkg1</a></li>

<li><a href="//example.com/pkg1/subpkg2">example.com/pkg1/subpkg2</a></li>


<li><a href="//example.com/pkg2">example.com/pkg2</a></li>

<li><a href="//example.com/pkg2/subpkg1">example.com/pkg2/subpkg1</a></li>

<li><a href="//example.com/pkg2/subpkg2/subsubpkg1">example.com/pkg2/subpkg2/subsubpkg1</a></li>


</ul>
</body>
</html>`,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		var out bytes.Buffer
		err := generate_index(&out, tc.domain, tc.r)
		if err != tc.expectedErr {
			t.Errorf("Test case %#v got err %#v, want %#v", tc, err, tc.expectedErr)
		} else if out.String() != tc.expectedOut {
			t.Logf("Expect: %s", tc.expectedOut)
			t.Logf("Out: %s", out.String())
			t.Errorf("Test case %#v got %#v, want %#v", tc, out.String(), tc.expectedOut)
		}
	}
}
