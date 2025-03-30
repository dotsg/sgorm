package testdata

import "embed"

//go:embed *.sql */*.go *.go
var Fixtures embed.FS
