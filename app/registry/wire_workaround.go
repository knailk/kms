//go:build tools
// +build tools

package registry

// Issue: https://github.com/google/wire/issues/299

import (
	_ "github.com/google/wire/cmd/wire"
)
