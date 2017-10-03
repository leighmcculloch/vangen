package main

import (
	"bytes"
	"testing"
)

func TestGenerate(t *testing.T) {
	testCases := []struct {
		domain      string
		pkg         string
		r           repository
		expectedOut string
		expectedErr error
	}{
		{
			domain: "example.com",
			pkg:    "pkg1",
			r: repository{
				Prefix: "pkg1",
				Subs:   []string{"subpkg1", "subpkg2"},
				Type:   "git",
				URL:    "https://repositoryhost.com/example/go-pkg1",
			},
			expectedOut: `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="go-import" content="example.com/pkg1 git https://repositoryhost.com/example/go-pkg1">
<meta name="go-source" content="example.com/pkg1 _ _ _">
<meta http-equiv="refresh" content="0; url=https://godoc.org/example.com/pkg1">
</head>
<body>
If you are not redirected, <a href="https://godoc.org/example.com/pkg1">click here</a>.
</body>
</html>`,
			expectedErr: nil,
		},
		{
			domain: "example.com",
			pkg:    "pkg1",
			r: repository{
				Prefix: "pkg1",
				Subs:   []string{"subpkg1", "subpkg2"},
				Type:   "git",
				URL:    "https://repositoryhost.com/example/go-pkg1",
				SourceURLs: sourceURLs{
					Home: "https://repositoryhost.com/example/go-pkg1/home",
					Dir:  "https://repositoryhost.com/example/go-pkg1/browser{/dir}",
					File: "https://repositoryhost.com/example/go-pkg1/view{/dir}{/file}",
				},
				Website: website{
					URL: "https://www.example.com",
				},
			},
			expectedOut: `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="go-import" content="example.com/pkg1 git https://repositoryhost.com/example/go-pkg1">
<meta name="go-source" content="example.com/pkg1 https://repositoryhost.com/example/go-pkg1/home https://repositoryhost.com/example/go-pkg1/browser{/dir} https://repositoryhost.com/example/go-pkg1/view{/dir}{/file}">
<meta http-equiv="refresh" content="0; url=https://www.example.com">
</head>
<body>
If you are not redirected, <a href="https://www.example.com">click here</a>.
</body>
</html>`,
			expectedErr: nil,
		},
		{
			domain: "example.com",
			pkg:    "pkg1/subpkg1",
			r: repository{
				Prefix: "pkg1",
				Subs:   []string{"subpkg1", "subpkg2"},
				Type:   "git",
				URL:    "https://repositoryhost.com/example/go-pkg1",
				SourceURLs: sourceURLs{
					Home: "https://repositoryhost.com/example/go-pkg1/home",
					Dir:  "https://repositoryhost.com/example/go-pkg1/browser{/dir}",
					File: "https://repositoryhost.com/example/go-pkg1/view{/dir}{/file}",
				},
				Website: website{
					URL: "https://www.example.com",
				},
			},
			expectedOut: `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="go-import" content="example.com/pkg1 git https://repositoryhost.com/example/go-pkg1">
<meta name="go-source" content="example.com/pkg1 https://repositoryhost.com/example/go-pkg1/home https://repositoryhost.com/example/go-pkg1/browser{/dir} https://repositoryhost.com/example/go-pkg1/view{/dir}{/file}">
<meta http-equiv="refresh" content="0; url=https://www.example.com">
</head>
<body>
If you are not redirected, <a href="https://www.example.com">click here</a>.
</body>
</html>`,
			expectedErr: nil,
		},
		{
			domain: "example.com",
			pkg:    "pkg1",
			r: repository{
				Prefix: "pkg1",
				Subs:   []string{"subpkg1", "subpkg2"},
				URL:    "https://github.com/example/go-pkg1",
			},
			expectedOut: `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="go-import" content="example.com/pkg1 git https://github.com/example/go-pkg1">
<meta name="go-source" content="example.com/pkg1 https://github.com/example/go-pkg1 https://github.com/example/go-pkg1/tree/master{/dir} https://github.com/example/go-pkg1/blob/master{/dir}/{file}#L{line}">
<meta http-equiv="refresh" content="0; url=https://godoc.org/example.com/pkg1">
</head>
<body>
If you are not redirected, <a href="https://godoc.org/example.com/pkg1">click here</a>.
</body>
</html>`,
			expectedErr: nil,
		},
		{
			domain: "example.com",
			pkg:    "pkg1/subpkg1",
			r: repository{
				Prefix: "pkg1",
				Subs:   []string{"subpkg1", "subpkg2"},
				URL:    "https://github.com/example/go-pkg1",
			},
			expectedOut: `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="go-import" content="example.com/pkg1 git https://github.com/example/go-pkg1">
<meta name="go-source" content="example.com/pkg1 https://github.com/example/go-pkg1 https://github.com/example/go-pkg1/tree/master{/dir} https://github.com/example/go-pkg1/blob/master{/dir}/{file}#L{line}">
<meta http-equiv="refresh" content="0; url=https://godoc.org/example.com/pkg1/subpkg1">
</head>
<body>
If you are not redirected, <a href="https://godoc.org/example.com/pkg1/subpkg1">click here</a>.
</body>
</html>`,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		var out bytes.Buffer
		err := generate_package(&out, tc.domain, tc.pkg, tc.r)
		if err != tc.expectedErr {
			t.Errorf("Test case %#v got err %#v, want %#v", tc, err, tc.expectedErr)
		} else if out.String() != tc.expectedOut {
			t.Logf("Out: %s", out.String())
			t.Errorf("Test case %#v got %#v, want %#v", tc, out.String(), tc.expectedOut)
		}
	}
}
