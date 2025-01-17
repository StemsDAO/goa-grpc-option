package codegen

import (
	"testing"

	"github.com/StemsDAO/goa-grpc-option/v3/codegen"
	"github.com/StemsDAO/goa-grpc-option/v3/expr"
	"github.com/StemsDAO/goa-grpc-option/v3/http/codegen/testdata"
)

func TestEncode(t *testing.T) {
	cases := []struct {
		Name string
		DSL  func()
		Code string
	}{
		{"header-bool", testdata.ResultHeaderBoolDSL, testdata.ResultHeaderBoolEncodeCode},
		{"header-int", testdata.ResultHeaderIntDSL, testdata.ResultHeaderIntEncodeCode},
		{"header-int32", testdata.ResultHeaderInt32DSL, testdata.ResultHeaderInt32EncodeCode},
		{"header-int64", testdata.ResultHeaderInt64DSL, testdata.ResultHeaderInt64EncodeCode},
		{"header-uint", testdata.ResultHeaderUIntDSL, testdata.ResultHeaderUIntEncodeCode},
		{"header-uint32", testdata.ResultHeaderUInt32DSL, testdata.ResultHeaderUInt32EncodeCode},
		{"header-uint64", testdata.ResultHeaderUInt64DSL, testdata.ResultHeaderUInt64EncodeCode},
		{"header-float32", testdata.ResultHeaderFloat32DSL, testdata.ResultHeaderFloat32EncodeCode},
		{"header-float64", testdata.ResultHeaderFloat64DSL, testdata.ResultHeaderFloat64EncodeCode},
		{"header-string", testdata.ResultHeaderStringDSL, testdata.ResultHeaderStringEncodeCode},
		{"header-bytes", testdata.ResultHeaderBytesDSL, testdata.ResultHeaderBytesEncodeCode},
		{"header-any", testdata.ResultHeaderAnyDSL, testdata.ResultHeaderAnyEncodeCode},
		{"header-array-bool", testdata.ResultHeaderArrayBoolDSL, testdata.ResultHeaderArrayBoolEncodeCode},
		{"header-array-int", testdata.ResultHeaderArrayIntDSL, testdata.ResultHeaderArrayIntEncodeCode},
		{"header-array-int32", testdata.ResultHeaderArrayInt32DSL, testdata.ResultHeaderArrayInt32EncodeCode},
		{"header-array-int64", testdata.ResultHeaderArrayInt64DSL, testdata.ResultHeaderArrayInt64EncodeCode},
		{"header-array-uint", testdata.ResultHeaderArrayUIntDSL, testdata.ResultHeaderArrayUIntEncodeCode},
		{"header-array-uint32", testdata.ResultHeaderArrayUInt32DSL, testdata.ResultHeaderArrayUInt32EncodeCode},
		{"header-array-uint64", testdata.ResultHeaderArrayUInt64DSL, testdata.ResultHeaderArrayUInt64EncodeCode},
		{"header-array-float32", testdata.ResultHeaderArrayFloat32DSL, testdata.ResultHeaderArrayFloat32EncodeCode},
		{"header-array-float64", testdata.ResultHeaderArrayFloat64DSL, testdata.ResultHeaderArrayFloat64EncodeCode},
		{"header-array-string", testdata.ResultHeaderArrayStringDSL, testdata.ResultHeaderArrayStringEncodeCode},
		{"header-array-bytes", testdata.ResultHeaderArrayBytesDSL, testdata.ResultHeaderArrayBytesEncodeCode},
		{"header-array-any", testdata.ResultHeaderArrayAnyDSL, testdata.ResultHeaderArrayAnyEncodeCode},

		{"header-bool-default", testdata.ResultHeaderBoolDefaultDSL, testdata.ResultHeaderBoolDefaultEncodeCode},
		{"header-bool-required-default", testdata.ResultHeaderBoolRequiredDefaultDSL, testdata.ResultHeaderBoolRequiredDefaultEncodeCode},
		{"header-string-default", testdata.ResultHeaderStringDefaultDSL, testdata.ResultHeaderStringDefaultEncodeCode},
		{"header-string-required-default", testdata.ResultHeaderStringRequiredDefaultDSL, testdata.ResultHeaderStringRequiredDefaultEncodeCode},
		{"header-array-bool-default", testdata.ResultHeaderArrayBoolDefaultDSL, testdata.ResultHeaderArrayBoolDefaultEncodeCode},
		{"header-array-bool-required-default", testdata.ResultHeaderArrayBoolRequiredDefaultDSL, testdata.ResultHeaderArrayBoolRequiredDefaultEncodeCode},
		{"header-array-string-default", testdata.ResultHeaderArrayStringDefaultDSL, testdata.ResultHeaderArrayStringDefaultEncodeCode},
		{"header-array-string-required-default", testdata.ResultHeaderArrayStringRequiredDefaultDSL, testdata.ResultHeaderArrayStringRequiredDefaultEncodeCode},

		{"body-string", testdata.ResultBodyStringDSL, testdata.ResultBodyStringEncodeCode},
		{"body-object", testdata.ResultBodyObjectDSL, testdata.ResultBodyObjectEncodeCode},
		{"body-user", testdata.ResultBodyUserDSL, testdata.ResultBodyUserEncodeCode},
		{"body-union", testdata.ResultBodyUnionDSL, testdata.ResultBodyUnionEncodeCode},
		{"body-result-multiple-views", testdata.ResultBodyMultipleViewsDSL, testdata.ResultBodyMultipleViewsEncodeCode},
		{"body-result-list-multiple-views", testdata.ResultBodyListDSL, testdata.ResultBodyListMultipleViewsEncodeCode},
		{"body-result-list-explicit-view", testdata.ResultBodyListExplicitViewDSL, testdata.ResultBodyListExplicitViewEncodeCode},
		{"empty-body-result-multiple-views", testdata.EmptyBodyResultMultipleViewsDSL, testdata.EmptyBodyResultMultipleViewsEncodeCode},
		{"body-array-string", testdata.ResultBodyArrayStringDSL, testdata.ResultBodyArrayStringEncodeCode},
		{"body-array-user", testdata.ResultBodyArrayUserDSL, testdata.ResultBodyArrayUserEncodeCode},
		{"body-primitive-string", testdata.ResultBodyPrimitiveStringDSL, testdata.ResultBodyPrimitiveStringEncodeCode},
		{"body-primitive-bool", testdata.ResultBodyPrimitiveBoolDSL, testdata.ResultBodyPrimitiveBoolEncodeCode},
		{"body-primitive-any", testdata.ResultBodyPrimitiveAnyDSL, testdata.ResultBodyPrimitiveAnyEncodeCode},
		{"body-primitive-array-string", testdata.ResultBodyPrimitiveArrayStringDSL, testdata.ResultBodyPrimitiveArrayStringEncodeCode},
		{"body-primitive-array-bool", testdata.ResultBodyPrimitiveArrayBoolDSL, testdata.ResultBodyPrimitiveArrayBoolEncodeCode},
		{"body-primitive-array-user", testdata.ResultBodyPrimitiveArrayUserDSL, testdata.ResultBodyPrimitiveArrayUserEncodeCode},
		{"body-inline-object", testdata.ResultBodyInlineObjectDSL, testdata.ResultBodyInlineObjectEncodeCode},

		{"body-header-object", testdata.ResultBodyHeaderObjectDSL, testdata.ResultBodyHeaderObjectEncodeCode},
		{"body-header-user", testdata.ResultBodyHeaderUserDSL, testdata.ResultBodyHeaderUserEncodeCode},

		{"explicit-body-primitive-result-multiple-views", testdata.ExplicitBodyPrimitiveResultMultipleViewsDSL, testdata.ExplicitBodyPrimitiveResultMultipleViewsEncodeCode},
		{"explicit-body-user-result-multiple-views", testdata.ExplicitBodyUserResultMultipleViewsDSL, testdata.ExplicitBodyUserResultMultipleViewsEncodeCode},
		{"explicit-body-result-list", testdata.ExplicitBodyResultListDSL, testdata.ExplicitBodyResultListEncodeCode},
		{"explicit-content-type-result", testdata.ExplicitContentTypeResultDSL, testdata.ExplicitContentTypeResultEncodeCode},
		{"explicit-content-type-response", testdata.ExplicitContentTypeResponseDSL, testdata.ExplicitContentTypeResponseEncodeCode},

		{"tag-string", testdata.ResultTagStringDSL, testdata.ResultTagStringEncodeCode},
		{"tag-string-required", testdata.ResultTagStringRequiredDSL, testdata.ResultTagStringRequiredEncodeCode},
		{"tag-result-multiple-views", testdata.ResultMultipleViewsTagDSL, testdata.ResultMultipleViewsTagEncodeCode},

		{"empty-server-response", testdata.EmptyServerResponseDSL, testdata.EmptyServerResponseEncodeCode},
		{"empty-server-response-with-tags", testdata.EmptyServerResponseWithTagsDSL, testdata.EmptyServerResponseWithTagsEncodeCode},

		{"result-with-custom-pkg-type", testdata.ResultWithCustomPkgTypeDSL, testdata.ResultWithCustomPkgTypeEncodeCode},
		{"result-with-embedded-custom-pkg-type", testdata.EmbeddedCustomPkgTypeDSL, testdata.ResultWithEmbeddedCustomPkgTypeEncodeCode},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			RunHTTPDSL(t, c.DSL)
			fs := ServerFiles("", expr.Root)
			if len(fs) != 2 {
				t.Fatalf("got %d files, expected two", len(fs))
			}
			sections := fs[1].SectionTemplates
			if len(sections) < 2 {
				t.Fatalf("got %d sections, expected at least 2", len(sections))
			}
			code := codegen.SectionCode(t, sections[1])
			if code != c.Code {
				t.Errorf("invalid code, got:\n%s\ngot vs. expected:\n%s", code, codegen.Diff(t, code, c.Code))
			}
		})
	}
}

func TestEncodeMarshallingAndUnmarshalling(t *testing.T) {
	cases := []struct {
		Name           string
		DSL            func()
		Code           []string
		SectionsOffset int
	}{
		{"embedded-custom-pkg-type", testdata.EmbeddedCustomPkgTypeDSL, []string{
			testdata.EmbeddedCustomPkgTypeUnmarshalCode,
			testdata.EmbeddedCustomPkgTypeMarshalCode}, 3},
		{"array-alias-extended", testdata.ArrayAliasExtendedDSL, []string{
			testdata.ArrayAliasExtendedUnmarshalCode,
			testdata.ArrayAliasExtendedMarshalCode}, 3},
		{"extension-with-alias", testdata.ExtensionWithAliasDSL, []string{
			testdata.ExtensionWithAliasUnmarshalExtensionCode,
			testdata.ExtensionWithAliasUnmarshalBarCode,
			testdata.ExtensionWithAliasMarshalResultCode,
			testdata.ExtensionWithAliasMarshalExtensionCode,
			testdata.ExtensionWithAliasMarshalBarCode}, 4},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			RunHTTPDSL(t, c.DSL)
			fs := ServerFiles("", expr.Root)
			if len(fs) != 2 {
				t.Fatalf("got %d files, expected two", len(fs))
			}
			sections := fs[1].SectionTemplates
			totalSectionsExpected := c.SectionsOffset + len(c.Code)
			if len(sections) != totalSectionsExpected {
				t.Fatalf("got %d sections, expected %d", len(sections), totalSectionsExpected)
			}
			for i := 0; i < len(c.Code); i++ {
				code := codegen.SectionCode(t, sections[c.SectionsOffset+i])
				if code != c.Code[i] {
					t.Errorf("%s %d invalid code, got:\n%s\ngot vs. expected:\n%s", sections[c.SectionsOffset+i].Name, i, code, codegen.Diff(t, code, c.Code[i]))
				}
			}
		})
	}
}
