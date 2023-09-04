package codegen

import (
	"bytes"
	"testing"

	"github.com/StemsDAO/goa-grpc-option/v3/codegen"
	"github.com/StemsDAO/goa-grpc-option/v3/expr"
	"github.com/StemsDAO/goa-grpc-option/v3/grpc/codegen/testdata"
)

func TestServerTypeFiles(t *testing.T) {
	cases := []struct {
		Name string
		DSL  func()
		Code string
	}{
		{"payload-with-nested-types", testdata.PayloadWithNestedTypesDSL, testdata.PayloadWithNestedTypesServerTypeCode},
		{"payload-with-duplicate-use", testdata.PayloadWithMultipleUseTypesDSL, testdata.PayloadWithMultipleUseTypesServerTypeCode},
		{"payload-with-alias-type", testdata.PayloadWithAliasTypeDSL, testdata.PayloadWithAliasTypeServerTypeCode},
		{"payload-with-mixed-attributes", testdata.PayloadWithMixedAttributesDSL, testdata.PayloadWithMixedAttributesServerTypeCode},
		{"result-list", testdata.ResultWithListDSL, testdata.ResultWithListServerTypeCode},
		{"with-errors", testdata.UnaryRPCWithErrorsDSL, testdata.WithErrorsServerTypeCode},
		{"elem-validation", testdata.ElemValidationDSL, testdata.ElemValidationServerTypesFile},
		{"alias-validation", testdata.AliasValidationDSL, testdata.AliasValidationServerTypesFile},
		{"struct-meta-type", testdata.StructMetaTypeDSL, testdata.StructMetaTypeServerTypeCode},
		{"default-fields", testdata.DefaultFieldsDSL, testdata.DefaultFieldsServerTypeCode},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			RunGRPCDSL(t, c.DSL)
			fs := ServerTypeFiles("", expr.Root)
			if len(fs) != 1 {
				t.Fatalf("got %d files, expected one", len(fs))
			}
			var buf bytes.Buffer
			for _, s := range fs[0].SectionTemplates[1:] {
				if err := s.Write(&buf); err != nil {
					t.Fatal(err)
				}
			}
			code := codegen.FormatTestCode(t, "package foo\n"+buf.String())
			if code != c.Code {
				t.Errorf("invalid code, got:\n%s\ngot vs. expected:\n%s", code, codegen.Diff(t, code, c.Code))
			}
		})
	}
}
