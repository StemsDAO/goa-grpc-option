package codegen

import (
	"bytes"
	"testing"

	"github.com/StemsDAO/goa-grpc-option/v3/codegen"
	"github.com/StemsDAO/goa-grpc-option/v3/codegen/service"
	"github.com/StemsDAO/goa-grpc-option/v3/expr"
)

// RunGRPCDSL returns the GRPC DSL root resulting from running the given DSL.
// It is used only in tests.
func RunGRPCDSL(t *testing.T, dsl func()) *expr.RootExpr {
	// reset all roots and codegen data structures
	service.Services = make(service.ServicesData)
	GRPCServices = make(ServicesData)
	return expr.RunDSL(t, dsl)
}

func sectionCode(t *testing.T, section ...*codegen.SectionTemplate) string {
	t.Helper()
	var code bytes.Buffer
	for _, s := range section {
		if err := s.Write(&code); err != nil {
			t.Fatal(err)
		}
	}
	return code.String()
}
