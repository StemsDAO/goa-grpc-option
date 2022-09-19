package testdata

const ResultWithMultipleViewsCode = `// ResultType is the viewed result type that is projected based on a view.
type ResultType struct {
	// Type to project
	Projected *ResultTypeView
	// View to render
	View string
}

// ResultTypeView is a type that runs validations on a projected type.
type ResultTypeView struct {
	A *string
	B *string
}

var (
	// ResultTypeMap is a map indexing the attribute names of ResultType by view
	// name.
	ResultTypeMap = map[string][]string{
		"default": {
			"a",
			"b",
		},
		"tiny": {
			"a",
		},
	}
)

// ValidateResultType runs the validations defined on the viewed result type
// ResultType.
func ValidateResultType(result *ResultType) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateResultTypeView(result.Projected)
	case "tiny":
		err = ValidateResultTypeViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "tiny"})
	}
	return
}

// ValidateResultTypeView runs the validations defined on ResultTypeView using
// the "default" view.
func ValidateResultTypeView(result *ResultTypeView) (err error) {
	if result.A == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("a", "result"))
	}
	if result.B == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("b", "result"))
	}
	return
}

// ValidateResultTypeViewTiny runs the validations defined on ResultTypeView
// using the "tiny" view.
func ValidateResultTypeViewTiny(result *ResultTypeView) (err error) {
	if result.A == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("a", "result"))
	}
	return
}
`

const ResultListMultipleViewsCode = `// ResultTypeList is the viewed result type that is projected based on a
// view.
type ResultTypeList struct {
	// Type to project
	Projected ResultTypeListView
	// View to render
	View string
}

// ResultTypeListView is a type that runs validations on a projected type.
type ResultTypeListView []*ResultTypeView

// ResultTypeView is a type that runs validations on a projected type.
type ResultTypeView struct {
	A *string
	B *string
}

var (
	// ResultTypeListMap is a map indexing the attribute names of
	// ResultTypeList by view name.
	ResultTypeListMap = map[string][]string{
		"default": {
			"a",
			"b",
		},
		"tiny": {
			"a",
		},
	}
	// ResultTypeMap is a map indexing the attribute names of ResultType by view
	// name.
	ResultTypeMap = map[string][]string{
		"default": {
			"a",
			"b",
		},
		"tiny": {
			"a",
		},
	}
)

// ValidateResultTypeList runs the validations defined on the viewed
// result type ResultTypeList.
func ValidateResultTypeList(result ResultTypeList) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateResultTypeListView(result.Projected)
	case "tiny":
		err = ValidateResultTypeListViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "tiny"})
	}
	return
}

// ValidateResultTypeListView runs the validations defined on
// ResultTypeListView using the "default" view.
func ValidateResultTypeListView(result ResultTypeListView) (err error) {
	for _, item := range result {
		if err2 := ValidateResultTypeView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateResultTypeListViewTiny runs the validations defined on
// ResultTypeListView using the "tiny" view.
func ValidateResultTypeListViewTiny(result ResultTypeListView) (err error) {
	for _, item := range result {
		if err2 := ValidateResultTypeViewTiny(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateResultTypeView runs the validations defined on ResultTypeView using
// the "default" view.
func ValidateResultTypeView(result *ResultTypeView) (err error) {
	if result.A == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("a", "result"))
	}
	if result.B == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("b", "result"))
	}
	return
}

// ValidateResultTypeViewTiny runs the validations defined on ResultTypeView
// using the "tiny" view.
func ValidateResultTypeViewTiny(result *ResultTypeView) (err error) {
	if result.A == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("a", "result"))
	}
	return
}
`

const ResultWithUserTypeCode = `// ResultType is the viewed result type that is projected based on a view.
type ResultType struct {
	// Type to project
	Projected *ResultTypeView
	// View to render
	View string
}

// ResultTypeView is a type that runs validations on a projected type.
type ResultTypeView struct {
	A *UserTypeView
	B *string
}

// UserTypeView is a type that runs validations on a projected type.
type UserTypeView struct {
	A *string
}

var (
	// ResultTypeMap is a map indexing the attribute names of ResultType by view
	// name.
	ResultTypeMap = map[string][]string{
		"default": {
			"a",
			"b",
		},
		"tiny": {
			"a",
		},
	}
)

// ValidateResultType runs the validations defined on the viewed result type
// ResultType.
func ValidateResultType(result *ResultType) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateResultTypeView(result.Projected)
	case "tiny":
		err = ValidateResultTypeViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "tiny"})
	}
	return
}

// ValidateResultTypeView runs the validations defined on ResultTypeView using
// the "default" view.
func ValidateResultTypeView(result *ResultTypeView) (err error) {
	if result.A == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("a", "result"))
	}
	return
}

// ValidateResultTypeViewTiny runs the validations defined on ResultTypeView
// using the "tiny" view.
func ValidateResultTypeViewTiny(result *ResultTypeView) (err error) {
	if result.A == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("a", "result"))
	}
	return
}

// ValidateUserTypeView runs the validations defined on UserTypeView.
func ValidateUserTypeView(result *UserTypeView) (err error) {

	return
}
`

const ResultWithResultTypeCode = `// RT is the viewed result type that is projected based on a view.
type RT struct {
	// Type to project
	Projected *RTView
	// View to render
	View string
}

// RTView is a type that runs validations on a projected type.
type RTView struct {
	A *string
	B *RT2View
	C *RT3View
}

// RT2View is a type that runs validations on a projected type.
type RT2View struct {
	C *string
	D *UserTypeView
	E *string
}

// UserTypeView is a type that runs validations on a projected type.
type UserTypeView struct {
	P *string
}

// RT3View is a type that runs validations on a projected type.
type RT3View struct {
	X []string
	Y map[int]*UserTypeView
	Z *string
}

var (
	// RTMap is a map indexing the attribute names of RT by view name.
	RTMap = map[string][]string{
		"default": {
			"a",
			"b",
			"c",
		},
		"tiny": {
			"b",
			"c",
		},
	}
	// RT2Map is a map indexing the attribute names of RT2 by view name.
	RT2Map = map[string][]string{
		"default": {
			"c",
			"d",
		},
		"extended": {
			"c",
			"d",
			"e",
		},
		"tiny": {
			"d",
		},
	}
	// RT3Map is a map indexing the attribute names of RT3 by view name.
	RT3Map = map[string][]string{
		"default": {
			"x",
			"y",
		},
		"tiny": {
			"x",
		},
	}
)

// ValidateRT runs the validations defined on the viewed result type RT.
func ValidateRT(result *RT) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateRTView(result.Projected)
	case "tiny":
		err = ValidateRTViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "tiny"})
	}
	return
}

// ValidateRTView runs the validations defined on RTView using the "default"
// view.
func ValidateRTView(result *RTView) (err error) {

	if result.B != nil {
		if err2 := ValidateRT2ViewExtended(result.B); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.C != nil {
		if err2 := ValidateRT3View(result.C); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateRTViewTiny runs the validations defined on RTView using the "tiny"
// view.
func ValidateRTViewTiny(result *RTView) (err error) {

	if result.B != nil {
		if err2 := ValidateRT2ViewTiny(result.B); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.C != nil {
		if err2 := ValidateRT3View(result.C); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateRT2View runs the validations defined on RT2View using the "default"
// view.
func ValidateRT2View(result *RT2View) (err error) {
	if result.C == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("c", "result"))
	}
	if result.D == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("d", "result"))
	}
	return
}

// ValidateRT2ViewExtended runs the validations defined on RT2View using the
// "extended" view.
func ValidateRT2ViewExtended(result *RT2View) (err error) {
	if result.C == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("c", "result"))
	}
	if result.D == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("d", "result"))
	}
	return
}

// ValidateRT2ViewTiny runs the validations defined on RT2View using the "tiny"
// view.
func ValidateRT2ViewTiny(result *RT2View) (err error) {
	if result.D == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("d", "result"))
	}
	return
}

// ValidateUserTypeView runs the validations defined on UserTypeView.
func ValidateUserTypeView(result *UserTypeView) (err error) {

	return
}

// ValidateRT3View runs the validations defined on RT3View using the "default"
// view.
func ValidateRT3View(result *RT3View) (err error) {
	if result.X == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("x", "result"))
	}
	if result.Y == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("y", "result"))
	}
	return
}

// ValidateRT3ViewTiny runs the validations defined on RT3View using the "tiny"
// view.
func ValidateRT3ViewTiny(result *RT3View) (err error) {
	if result.X == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("x", "result"))
	}
	return
}
`

const ResultWithRecursiveResultTypeCode = `// RT is the viewed result type that is projected based on a view.
type RT struct {
	// Type to project
	Projected *RTView
	// View to render
	View string
}

// RTView is a type that runs validations on a projected type.
type RTView struct {
	A *RTView
}

var (
	// RTMap is a map indexing the attribute names of RT by view name.
	RTMap = map[string][]string{
		"default": {
			"a",
		},
		"tiny": {
			"a",
		},
	}
)

// ValidateRT runs the validations defined on the viewed result type RT.
func ValidateRT(result *RT) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateRTView(result.Projected)
	case "tiny":
		err = ValidateRTViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "tiny"})
	}
	return
}

// ValidateRTView runs the validations defined on RTView using the "default"
// view.
func ValidateRTView(result *RTView) (err error) {
	if result.A == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("a", "result"))
	}
	if result.A != nil {
		if err2 := ValidateRTViewTiny(result.A); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateRTViewTiny runs the validations defined on RTView using the "tiny"
// view.
func ValidateRTViewTiny(result *RTView) (err error) {
	if result.A == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("a", "result"))
	}
	if result.A != nil {
		if err2 := ValidateRTView(result.A); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}
`

const ResultWithCustomFieldsCode = `// RT is the viewed result type that is projected based on a view.
type RT struct {
	// Type to project
	Projected *RTView
	// View to render
	View string
}

// RTView is a type that runs validations on a projected type.
type RTView struct {
	CustomA *string
	B       *int
}

var (
	// RTMap is a map indexing the attribute names of RT by view name.
	RTMap = map[string][]string{
		"default": {
			"a",
			"b",
		},
		"tiny": {
			"a",
		},
	}
)

// ValidateRT runs the validations defined on the viewed result type RT.
func ValidateRT(result *RT) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateRTView(result.Projected)
	case "tiny":
		err = ValidateRTViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "tiny"})
	}
	return
}

// ValidateRTView runs the validations defined on RTView using the "default"
// view.
func ValidateRTView(result *RTView) (err error) {
	if result.CustomA == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("a", "result"))
	}
	if result.B == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("b", "result"))
	}
	return
}

// ValidateRTViewTiny runs the validations defined on RTView using the "tiny"
// view.
func ValidateRTViewTiny(result *RTView) (err error) {
	if result.CustomA == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("a", "result"))
	}
	return
}
`

const ResultWithRecursiveListOfResultTypeCode = `// SomeRT is the viewed result type that is projected based on a view.
type SomeRT struct {
	// Type to project
	Projected *SomeRTView
	// View to render
	View string
}

// AnotherResult is the viewed result type that is projected based on a view.
type AnotherResult struct {
	// Type to project
	Projected *AnotherResultView
	// View to render
	View string
}

// SomeRTView is a type that runs validations on a projected type.
type SomeRTView struct {
	A SomeRTListView
}

// SomeRTListView is a type that runs validations on a projected type.
type SomeRTListView []*SomeRTView

// AnotherResultView is a type that runs validations on a projected type.
type AnotherResultView struct {
	A AnotherResultListView
}

// AnotherResultListView is a type that runs validations on a projected
// type.
type AnotherResultListView []*AnotherResultView

var (
	// SomeRTMap is a map indexing the attribute names of SomeRT by view name.
	SomeRTMap = map[string][]string{
		"default": {
			"a",
		},
		"tiny": {
			"a",
		},
	}
	// AnotherResultMap is a map indexing the attribute names of AnotherResult by
	// view name.
	AnotherResultMap = map[string][]string{
		"default": {
			"a",
		},
	}
	// SomeRTListMap is a map indexing the attribute names of
	// SomeRTList by view name.
	SomeRTListMap = map[string][]string{
		"default": {
			"a",
		},
		"tiny": {
			"a",
		},
	}
	// AnotherResultListMap is a map indexing the attribute names of
	// AnotherResultList by view name.
	AnotherResultListMap = map[string][]string{
		"default": {
			"a",
		},
	}
)

// ValidateSomeRT runs the validations defined on the viewed result type SomeRT.
func ValidateSomeRT(result *SomeRT) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateSomeRTView(result.Projected)
	case "tiny":
		err = ValidateSomeRTViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "tiny"})
	}
	return
}

// ValidateAnotherResult runs the validations defined on the viewed result type
// AnotherResult.
func ValidateAnotherResult(result *AnotherResult) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateAnotherResultView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateSomeRTView runs the validations defined on SomeRTView using the
// "default" view.
func ValidateSomeRTView(result *SomeRTView) (err error) {

	if result.A != nil {
		if err2 := ValidateSomeRTListViewTiny(result.A); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateSomeRTViewTiny runs the validations defined on SomeRTView using the
// "tiny" view.
func ValidateSomeRTViewTiny(result *SomeRTView) (err error) {

	if result.A != nil {
		if err2 := ValidateSomeRTListView(result.A); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateSomeRTListView runs the validations defined on
// SomeRTListView using the "default" view.
func ValidateSomeRTListView(result SomeRTListView) (err error) {
	for _, item := range result {
		if err2 := ValidateSomeRTView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateSomeRTListViewTiny runs the validations defined on
// SomeRTListView using the "tiny" view.
func ValidateSomeRTListViewTiny(result SomeRTListView) (err error) {
	for _, item := range result {
		if err2 := ValidateSomeRTViewTiny(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateAnotherResultView runs the validations defined on AnotherResultView
// using the "default" view.
func ValidateAnotherResultView(result *AnotherResultView) (err error) {

	if result.A != nil {
		if err2 := ValidateAnotherResultListView(result.A); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateAnotherResultListView runs the validations defined on
// AnotherResultListView using the "default" view.
func ValidateAnotherResultListView(result AnotherResultListView) (err error) {
	for _, item := range result {
		if err2 := ValidateAnotherResultView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}
`

const ResultWithMultipleMethodsCode = `// RT is the viewed result type that is projected based on a view.
type RT struct {
	// Type to project
	Projected *RTView
	// View to render
	View string
}

// RTView is a type that runs validations on a projected type.
type RTView struct {
	A *string
}

var (
	// RTMap is a map indexing the attribute names of RT by view name.
	RTMap = map[string][]string{
		"default": {
			"a",
		},
	}
)

// ValidateRT runs the validations defined on the viewed result type RT.
func ValidateRT(result *RT) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateRTView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateRTView runs the validations defined on RTView using the "default"
// view.
func ValidateRTView(result *RTView) (err error) {

	return
}
`

const ResultWithEnumType = `// Result is the viewed result type that is projected based on a view.
type Result struct {
	// Type to project
	Projected *ResultView
	// View to render
	View string
}

// ResultView is a type that runs validations on a projected type.
type ResultView struct {
	T []UserTypeView
}

// UserTypeView is a type that runs validations on a projected type.
type UserTypeView string

var (
	// ResultMap is a map indexing the attribute names of Result by view name.
	ResultMap = map[string][]string{
		"default": {
			"t",
		},
	}
)

// ValidateResult runs the validations defined on the viewed result type Result.
func ValidateResult(result *Result) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateResultView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateResultView runs the validations defined on ResultView using the
// "default" view.
func ValidateResultView(result *ResultView) (err error) {
	for _, e := range result.T {
		if !(string(e) == "a" || string(e) == "b") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.t[*]", string(e), []interface{}{"a", "b"}))
		}
	}
	return
}

// ValidateUserTypeView runs the validations defined on UserTypeView.
func ValidateUserTypeView(result UserTypeView) (err error) {
	if !(string(result) == "a" || string(result) == "b") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError("result", string(result), []interface{}{"a", "b"}))
	}
	return
}
`
