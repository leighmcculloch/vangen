package main

import (
	"bytes"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestGenerate(t *testing.T) {
	testCases := []struct {
		description string
		domain      string
		pkg         string
		r           repository
		expectedOut string
		expectedErr error
	}{
		{
			description: "simple",
			domain:      "example.com",
			pkg:         "pkg1",
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
<style>
* { font-family: sans-serif; }
body { margin-top: 0; }
.content { display: inline-block; }
code { display: block; font-family: monospace; font-size: 1em; background-color: #d5d5d5; padding: 1em; margin-bottom: 16px; }
ul { margin-top: 16px; margin-bottom: 16px; }
</style>
</head>
<body>
<div class="content">
<h2>Go package: example.com/pkg1</h2>
<code>go get example.com/pkg1</code>
<code>import "example.com/pkg1"</code>
Links:
<ul>
<li>Home: <a href="https://godoc.org/example.com/pkg1">https://godoc.org/example.com/pkg1</a></code>
<li>Source: <a href="https://repositoryhost.com/example/go-pkg1">https://repositoryhost.com/example/go-pkg1</a></code>
</ul>
Sub-packages:<ul><li><a href="/pkg1/subpkg1">example.com/pkg1/subpkg1</a></li><li><a href="/pkg1/subpkg2">example.com/pkg1/subpkg2</a></li></ul></div>
</body>
</html>`,
			expectedErr: nil,
		},
		{
			description: "custom source urls",
			domain:      "example.com",
			pkg:         "pkg1",
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
<style>
* { font-family: sans-serif; }
body { margin-top: 0; }
.content { display: inline-block; }
code { display: block; font-family: monospace; font-size: 1em; background-color: #d5d5d5; padding: 1em; margin-bottom: 16px; }
ul { margin-top: 16px; margin-bottom: 16px; }
</style>
</head>
<body>
<div class="content">
<h2>Go package: example.com/pkg1</h2>
<code>go get example.com/pkg1</code>
<code>import "example.com/pkg1"</code>
Links:
<ul>
<li>Home: <a href="https://www.example.com">https://www.example.com</a></code>
<li>Source: <a href="https://repositoryhost.com/example/go-pkg1">https://repositoryhost.com/example/go-pkg1</a></code>
</ul>
Sub-packages:<ul><li><a href="/pkg1/subpkg1">example.com/pkg1/subpkg1</a></li><li><a href="/pkg1/subpkg2">example.com/pkg1/subpkg2</a></li></ul></div>
</body>
</html>`,
			expectedErr: nil,
		},
		{
			description: "sub-package",
			domain:      "example.com",
			pkg:         "pkg1/subpkg1",
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
<style>
* { font-family: sans-serif; }
body { margin-top: 0; }
.content { display: inline-block; }
code { display: block; font-family: monospace; font-size: 1em; background-color: #d5d5d5; padding: 1em; margin-bottom: 16px; }
ul { margin-top: 16px; margin-bottom: 16px; }
</style>
</head>
<body>
<div class="content">
<h2>Go package: example.com/pkg1/subpkg1</h2>
<code>go get example.com/pkg1/subpkg1</code>
<code>import "example.com/pkg1/subpkg1"</code>
Links:
<ul>
<li>Home: <a href="https://www.example.com">https://www.example.com</a></code>
<li>Source: <a href="https://repositoryhost.com/example/go-pkg1">https://repositoryhost.com/example/go-pkg1</a></code>
</ul>
Sub-packages:<ul><li><a href="/pkg1/subpkg1">example.com/pkg1/subpkg1</a></li><li><a href="/pkg1/subpkg2">example.com/pkg1/subpkg2</a></li></ul></div>
</body>
</html>`,
			expectedErr: nil,
		},
		{
			description: "github defaults",
			domain:      "example.com",
			pkg:         "pkg1",
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
<style>
* { font-family: sans-serif; }
body { margin-top: 0; }
.content { display: inline-block; }
code { display: block; font-family: monospace; font-size: 1em; background-color: #d5d5d5; padding: 1em; margin-bottom: 16px; }
ul { margin-top: 16px; margin-bottom: 16px; }
</style>
</head>
<body>
<div class="content">
<h2>Go package: example.com/pkg1</h2>
<code>go get example.com/pkg1</code>
<code>import "example.com/pkg1"</code>
Links:
<ul>
<li>Home: <a href="https://godoc.org/example.com/pkg1">https://godoc.org/example.com/pkg1</a></code>
<li>Source: <a href="https://github.com/example/go-pkg1">https://github.com/example/go-pkg1</a></code>
</ul>
Sub-packages:<ul><li><a href="/pkg1/subpkg1">example.com/pkg1/subpkg1</a></li><li><a href="/pkg1/subpkg2">example.com/pkg1/subpkg2</a></li></ul></div>
</body>
</html>`,
			expectedErr: nil,
		},
		{
			description: "sub-package github defaults",
			domain:      "example.com",
			pkg:         "pkg1/subpkg1",
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
<style>
* { font-family: sans-serif; }
body { margin-top: 0; }
.content { display: inline-block; }
code { display: block; font-family: monospace; font-size: 1em; background-color: #d5d5d5; padding: 1em; margin-bottom: 16px; }
ul { margin-top: 16px; margin-bottom: 16px; }
</style>
</head>
<body>
<div class="content">
<h2>Go package: example.com/pkg1/subpkg1</h2>
<code>go get example.com/pkg1/subpkg1</code>
<code>import "example.com/pkg1/subpkg1"</code>
Links:
<ul>
<li>Home: <a href="https://godoc.org/example.com/pkg1/subpkg1">https://godoc.org/example.com/pkg1/subpkg1</a></code>
<li>Source: <a href="https://github.com/example/go-pkg1">https://github.com/example/go-pkg1</a></code>
</ul>
Sub-packages:<ul><li><a href="/pkg1/subpkg1">example.com/pkg1/subpkg1</a></li><li><a href="/pkg1/subpkg2">example.com/pkg1/subpkg2</a></li></ul></div>
</body>
</html>`,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		var out bytes.Buffer
		err := generate_package(&out, tc.domain, tc.pkg, tc.r)
		if err != tc.expectedErr {
			t.Errorf("Test case %q got err %#v, want %#v", tc.description, err, tc.expectedErr)
		} else if out.String() != tc.expectedOut {
			dmp := diffmatchpatch.New()
			diffs := dmp.DiffMain(tc.expectedOut, out.String(), false)
			t.Errorf("Test case %q got: \n%s\nAs diff:\n%s", tc.description, out.String(), dmp.DiffPrettyText(diffs))
		}
	}
}
