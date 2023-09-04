package codegen

import (
	"testing"

	"github.com/StemsDAO/goa-grpc-option/v3/codegen/testdata"
	"github.com/StemsDAO/goa-grpc-option/v3/expr"
)

func TestRecursiveValidationCode(t *testing.T) {
	root := RunDSL(t, testdata.ValidationTypesDSL)
	var (
		scope = NewNameScope()

		integerT = root.UserType("Integer")
		stringT  = root.UserType("String")
		floatT   = root.UserType("Float")
		aliasT   = root.UserType("AliasType")
		userT    = root.UserType("UserType")
		arrayUT  = root.UserType("ArrayUserType")
		arrayT   = root.UserType("Array")
		mapT     = root.UserType("Map")
		rtT      = root.UserType("Result")
		rtcolT   = root.UserType("List")
		colT     = root.UserType("TypeWithList")
	)
	cases := []struct {
		Name       string
		Type       expr.UserType
		Required   bool
		Pointer    bool
		UseDefault bool
		Code       string
	}{
		{"integer-required", integerT, true, false, false, testdata.IntegerRequiredValidationCode},
		{"integer-pointer", integerT, false, true, false, testdata.IntegerPointerValidationCode},
		{"integer-use-default", integerT, false, false, true, testdata.IntegerUseDefaultValidationCode},
		{"float-required", floatT, true, false, false, testdata.FloatRequiredValidationCode},
		{"float-pointer", floatT, false, true, false, testdata.FloatPointerValidationCode},
		{"float-use-default", floatT, false, false, true, testdata.FloatUseDefaultValidationCode},
		{"string-required", stringT, true, false, false, testdata.StringRequiredValidationCode},
		{"string-pointer", stringT, false, true, false, testdata.StringPointerValidationCode},
		{"string-use-default", stringT, false, false, true, testdata.StringUseDefaultValidationCode},
		{"alias-type", aliasT, true, false, false, testdata.AliasTypeValidationCode},
		{"user-type-required", userT, true, false, false, testdata.UserTypeRequiredValidationCode},
		{"user-type-pointer", userT, false, true, false, testdata.UserTypePointerValidationCode},
		{"user-type-default", userT, false, false, true, testdata.UserTypeUseDefaultValidationCode},
		{"user-type-array-required", arrayUT, true, true, false, testdata.UserTypeArrayValidationCode},
		{"array-required", arrayT, true, false, false, testdata.ArrayRequiredValidationCode},
		{"array-pointer", arrayT, false, true, false, testdata.ArrayPointerValidationCode},
		{"array-use-default", arrayT, false, false, true, testdata.ArrayUseDefaultValidationCode},
		{"map-required", mapT, true, false, false, testdata.MapRequiredValidationCode},
		{"map-pointer", mapT, false, true, false, testdata.MapPointerValidationCode},
		{"map-use-default", mapT, false, false, true, testdata.MapUseDefaultValidationCode},
		{"result-type-pointer", rtT, false, true, false, testdata.ResultTypePointerValidationCode},
		{"list-required", rtcolT, true, false, false, testdata.ResultListPointerValidationCode},
		{"list-pointer", rtcolT, false, true, false, testdata.ResultListPointerValidationCode},
		{"type-with-list-pointer", colT, false, true, false, testdata.TypeWithListPointerValidationCode},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			ctx := NewAttributeContext(c.Pointer, false, c.UseDefault, "", scope)
			code := RecursiveValidationCode(&expr.AttributeExpr{Type: c.Type}, ctx, c.Required, expr.IsAlias(c.Type), "target")
			code = FormatTestCode(t, "package foo\nfunc Validate() (err error){\n"+code+"}")
			if code != c.Code {
				t.Errorf("invalid code, got:\n%s\ngot vs. expected:\n%s", code, Diff(t, code, c.Code))
			}
		})
	}
}
