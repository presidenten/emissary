// +build pin

// Package ignore is mostly ignored; it just serves to pin the
// packages of external commands (i.e. things that we don't use as
// libraries) in to the module, so that `go mod tidy` won't make the
// `go.mod` file forget which version we want.
package ignore

import (
	// protoc-gen-validate
	_ "github.com/envoyproxy/protoc-gen-validate"

	// protoc-gen-gogofast
	_ "github.com/gogo/protobuf/protoc-gen-gogofast"

	// protoc-gen-go-json
	_ "github.com/mitchellh/protoc-gen-go-json"

	// controller-gen
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"
)
