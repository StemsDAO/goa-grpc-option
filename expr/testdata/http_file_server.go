package testdata

import (
	. "github.com/StemsDAO/goa-grpc-option/v3/dsl"
)

var FilesValidDSL = func() {
	Service("files-dsl", func() {
		Files("path", "filename")
	})
}

var FilesIncompatibleDSL = func() {
	API("files-incompatile", func() {
		Files("path", "filename")
	})
}
