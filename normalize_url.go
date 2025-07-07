package main

import (
	"errors"
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	// Check for empty URL
	if inputURL == "" {
		return "", errors.New("empty URL")
	}

	// Parse the URL
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	// Check for invalid URLs (no host)
	if parsedURL.Host == "" {
		return "", errors.New("invalid URL: missing host")
	}

	// Convert host to lowercase
	host := strings.ToLower(parsedURL.Host)

	// Remove default ports
	if parsedURL.Port() == "80" && parsedURL.Scheme == "http" {
		host = parsedURL.Hostname()
	} else if parsedURL.Port() == "443" && parsedURL.Scheme == "https" {
		host = parsedURL.Hostname()
	}

	// Get the path - use EscapedPath to preserve URL encoding
	path := parsedURL.EscapedPath()

	// Remove trailing slash from path (including when path is just "/")
	if strings.HasSuffix(path, "/") && len(path) > 0 {
		path = strings.TrimSuffix(path, "/")
	}

	// Convert path to lowercase after removing trailing slash
	path = strings.ToLower(path)

	// Build the normalized URL: host + path
	// NOTE: We're removing scheme, query params, fragment, and userinfo
	normalized := host + path

	return normalized, nil
}
