/*
Package tools is used to track binary dependencies with go modules
https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
*/
package tools

// +build tools

import (
	// linter(s)
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"

	// kubernetes code generators
	_ "k8s.io/code-generator/cmd/deepcopy-gen"

	// test runner
	_ "gotest.tools/gotestsum"

	// standard go proto compiler
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
