package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove http scheme",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove trailing slash",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove query parameters",
			inputURL: "https://blog.boot.dev/path?id=123&name=test",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove fragment",
			inputURL: "https://blog.boot.dev/path#section",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove query and fragment",
			inputURL: "https://blog.boot.dev/path?id=123#section",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "convert to lowercase",
			inputURL: "https://BLOG.BOOT.DEV/PATH",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove port 80 for http",
			inputURL: "http://blog.boot.dev:80/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove port 443 for https",
			inputURL: "https://blog.boot.dev:443/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "keep non-standard port",
			inputURL: "https://blog.boot.dev:8080/path",
			expected: "blog.boot.dev:8080/path",
		},
		{
			name:     "remove userinfo",
			inputURL: "https://user:pass@blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "just domain",
			inputURL: "https://blog.boot.dev",
			expected: "blog.boot.dev",
		},
		{
			name:     "just domain with trailing slash",
			inputURL: "https://blog.boot.dev/",
			expected: "blog.boot.dev",
		},
		{
			name:     "with www subdomain",
			inputURL: "https://www.blog.boot.dev/path",
			expected: "www.blog.boot.dev/path",
		},
		{
			name:     "with multiple path segments",
			inputURL: "https://blog.boot.dev/path/to/resource",
			expected: "blog.boot.dev/path/to/resource",
		},
		{
			name:     "with encoded characters",
			inputURL: "https://blog.boot.dev/path%20with%20spaces",
			expected: "blog.boot.dev/path%20with%20spaces",
		},
		{
			name:     "complex URL with everything",
			inputURL: "https://user:pass@blog.boot.dev:443/path/to/resource?query=value&other=test#fragment",
			expected: "blog.boot.dev/path/to/resource",
		},
	}

	// Test invalid URLs separately
	invalidTests := []struct {
		name     string
		inputURL string
	}{
		{
			name:     "empty URL",
			inputURL: "",
		},
		{
			name:     "invalid scheme",
			inputURL: "://blog.boot.dev/path",
		},
		{
			name:     "malformed URL",
			inputURL: "ht tp://blog.boot.dev/path",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}

	// Test invalid URLs - these should return errors
	for i, tc := range invalidTests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := normalizeURL(tc.inputURL)
			if err == nil {
				t.Errorf("Test %v - '%s' FAIL: expected error but got none", i, tc.name)
			}
		})
	}
}
