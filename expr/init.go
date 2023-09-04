package expr

import (
	"github.com/StemsDAO/goa-grpc-option/v3/eval"
)

// Register DSL roots.
func init() {
	if err := eval.Register(Root); err != nil {
		panic(err) // bug
	}
	if err := eval.Register(Root.GeneratedTypes); err != nil {
		panic(err) // bug
	}
}
