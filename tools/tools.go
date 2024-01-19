//go:build tools

package tools

// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

//go:generate go install github.com/golangci/golangci-lint/cmd/golangci-lint
//go:generate go install mvdan.cc/gofumpt
//go:generate go install github.com/daixiang0/gci
//go:generate go install gotest.tools/gotestsum
//go:generate go install golang.org/x/tools/cmd/goimports
//go:generate go install golang.org/x/lint/golint
//go:generate go install github.com/go-critic/go-critic/cmd/gocritic

// nolint
import (
	// gci
	_ "github.com/daixiang0/gci"
	// gocritic
	_ "github.com/go-critic/go-critic/cmd/gocritic"
	// golangci-lint
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	// golint
	_ "golang.org/x/lint/golint"
	// goimports
	_ "golang.org/x/tools/cmd/goimports"
	// gotestsum
	_ "gotest.tools/gotestsum"
	// gofumpt
	_ "mvdan.cc/gofumpt"
)
