package validator_test

import (
	"testing"
	"validator"
)

var sv = validator.NewStructValidator()

/**
mandatory constraint validations test
*/
type ReqMsg struct {
	Name string `json:"name" constraints:"min-length=5"`
	Age  int    `json:"age" constraints:"min=10"`
}

func TestRequiredConstraintFail(t *testing.T) {
	msg := ReqMsg{
		Name: "Test",
		Age:  11,
	}
	err := sv.Validate(msg)
	got := err.Error()
	want := "mandatory fields not present"
	if got != want {
		t.Errorf("Expected: %s, got: %s", got, want)
	}
}

type ReqMsg2 struct {
	Name string `json:"name" constraints:"required=true;nillable=true"`
	Age  int    `json:"age" constraints:"required=true;nillable=true"`
}

func TestSuccessValidation(t *testing.T) {
	msg := ReqMsg2{
		Name: "Test",
		Age:  11,
	}
	if err := sv.Validate(msg); err != nil {
		t.Errorf("Error in validation: %s", err)
	}
}

/**
min, max validations test
*/

func TestNumericValidations(t *testing.T) {

	testsPass := []struct {
		Name  string
		input interface{}
	}{
		{
			Name: "Test1",
			input: struct {
				MinC1 int `json:"minC1" constraints:"required=true;nillable=true;min=10"`
				MaxC1 int `json:"maxC1" constraints:"required=true;nillable=true;max=49"`
			}{MinC1: 12, MaxC1: 45},
		},
		/**
		exclusive min/max validation test
		*/
		{
			Name: "Test4",
			input: struct {
				MinC4 int `json:"minC4" constraints:"required=true;nillable=true;exclusiveMin=10"`
				MaxC4 int `json:"maxC4" constraints:"required=true;nillable=true;exclusiveMax=50"`
			}{MinC4: 10, MaxC4: 50},
		},
		{
			Name: "Test7",
			input: struct {
				Num7 int `json:"num7" constraints:"required=true;nillable=true;multipleOf=5"`
			}{Num7: 10},
		},
	}

	for _, tt := range testsPass {
		t.Run(tt.Name, func(t *testing.T) {
			err := sv.Validate(tt.input)
			if err != nil {
				t.Errorf("Error in validation: %s", err)
			}
		})
	}

	testsError := []struct {
		Name  string
		input interface{}
		want  string
	}{
		{
			Name: "Test2",
			input: struct {
				MinC2 int `json:"minC2" constraints:"required=true;nillable=true;min=10"`
				MaxC2 int `json:"maxC2" constraints:"required=true;nillable=true;max=49"`
			}{MinC2: 7, MaxC2: 45},
			want: "min value validation failed",
		},
		{
			Name: "Test3",
			input: struct {
				MinC3 int `json:"minC3" constraints:"required=true;nillable=true;min=10"`
				MaxC3 int `json:"maxC3" constraints:"required=true;nillable=true;max=49"`
			}{MinC3: 12, MaxC3: 55},
			want: "max value validation failed",
		},
		/**
		exclusive min/max validation test
		*/
		{
			Name: "Test5",
			input: struct {
				MinC5 int `json:"minC5" constraints:"required=true;nillable=true;exclusiveMin=10"`
				MaxC5 int `json:"maxC5" constraints:"required=true;nillable=true;exclusiveMax=50"`
			}{MinC5: 9, MaxC5: 50},
			want: "exclusive min validation failed",
		},
		{
			Name: "Test6",
			input: struct {
				MinC6 int `json:"minC6" constraints:"required=true;nillable=true;exclusiveMin=10"`
				MaxC6 int `json:"maxC6" constraints:"required=true;nillable=true;exclusiveMax=50"`
			}{MinC6: 10, MaxC6: 51},
			want: "exclusive max validation failed",
		},
		{
			Name: "Test8",
			input: struct {
				Num8 int `json:"num8" constraints:"required=true;nillable=true;multipleOf=5"`
			}{Num8: 11},
			want: "multipleOf validation failed",
		},
	}

	for _, tt := range testsError {
		t.Run(tt.Name, func(t *testing.T) {
			err := sv.Validate(tt.input)
			if tt.want != err.Error() {
				t.Errorf("Got: %s, want: %s", err, tt.want)
			}
		})
	}
}

func TestStringValidation(t *testing.T) {
	testsPass := []struct {
		Name  string
		input interface{}
	}{
		{
			Name: "Test1",
			input: struct {
				Str1T1 string `json:"str1T1" constraints:"required=true;nillable=true;min-length=10"`
				Str2T1 string `json:"str2T1" constraints:"required=true;nillable=true;max-length=15"`
			}{Str1T1: "hello_world", Str2T1: "hello_world_go"},
		},
		/**
		pattern validations
		*/
		{
			Name: "Test4",
			input: struct {
				Str4 string `json:"str4" constraints:"required=true;nillable=false;pattern=^[tes]{4}.*"`
			}{Str4: "test1234"},
		},
		{
			Name: "Test7",
			input: struct {
				Str4 string `json:"str4" constraints:"required=true;nillable=false;pattern=gray|grey"`
			}{Str4: "grey"},
		},
	}

	for _, tt := range testsPass {
		t.Run(tt.Name, func(t *testing.T) {
			err := sv.Validate(tt.input)
			if err != nil {
				t.Errorf("Error in validation: %s", err)
			}
		})
	}

	testsFail := []struct {
		Name  string
		input interface{}
		want  string
	}{
		{
			Name: "Test2",
			input: struct {
				Str1T2 string `json:"str1T2" constraints:"required=true;nillable=true;min-length=10"`
				Str2T2 string `json:"str2T2" constraints:"required=true;nillable=true;max-length=15"`
			}{Str1T2: "hell_worl", Str2T2: "hello_world_go"},
			want: "min-length validation failed",
		},
		{
			Name: "Test3",
			input: struct {
				Str1T3 string `json:"str1T3" constraints:"required=true;nillable=true;min-length=10"`
				Str2T3 string `json:"str2T3" constraints:"required=true;nillable=true;max-length=15"`
			}{Str1T3: "hello_world", Str2T3: "hello_world_from_go"},
			want: "max-length validation failed",
		},
		/**
		pattern validations
		*/
		{
			Name: "Test5",
			input: struct {
				Str5 string `json:"str5" constraints:"required=true;nillable=false;pattern=^[tes]{4}.*"`
			}{Str5: "abcd1234"},
			want: "pattern validation failed",
		},
		{
			Name: "Test6",
			input: struct {
				Str6 string `json:"str6" constraints:"required=true;nillable=false;pattern=["`
			}{Str6: "tsst1234"},
			want: "invalid constraint value",
		},
		{
			Name: "Test8",
			input: struct {
				Str string `json:"str" constraints:"required=true;nillable=false;pattern=gray|grey"`
			}{Str: "gry"},
			want: "pattern validation failed",
		},
	}

	for _, tt := range testsFail {
		t.Run(tt.Name, func(t *testing.T) {
			err := sv.Validate(tt.input)
			if tt.want != err.Error() {
				t.Errorf("Got: %s, want: %s", err, tt.want)
			}
		})
	}
}

/**
required validations test
*/

func TestReqValidation(t *testing.T) {

	testsPass := []struct {
		Name  string
		input interface{}
	}{
		{
			Name: "Test1",
			input: struct {
				Name1 string `json:"name1" constraints:"required=true;nillable=false"`
			}{Name1: "abcd"},
		},
	}
	for _, tt := range testsPass {
		t.Run(tt.Name, func(t *testing.T) {
			err := sv.Validate(tt.input)
			if err != nil {
				t.Errorf("Error in validation: %s", err)
			}
		})
	}

	testsFail := []struct {
		Name  string
		input interface{}
		want  string
	}{
		{
			Name: "Test2",
			input: struct {
				Name2 string `json:"name2" constraints:"required=true;nillable=false"`
			}{Name2: ""},
			want: "required validation failed",
		},
	}

	for _, tt := range testsFail {
		t.Run(tt.Name, func(t *testing.T) {
			err := sv.Validate(tt.input)
			if tt.want != err.Error() {
				t.Errorf("Got: %s, want: %s", err, tt.want)
			}
		})
	}
}

/**
Nested Structure Testing
**Not working as per latest algo**
*/

type Example struct {
	Reference
	Summary     string      `json:"summary,omitempty" constraints:"required=true;nillable=false"`
	Description string      `json:"description,omitempty" constraints:"required=true;nillable=false"`
	Value       interface{} `json:"example,omitempty" constraints:"required=true;nillable=false"`
}

type Reference struct {
	Ref            string `json:"ref" constraints:"required=true;nillable=false"`
	RefDescription string `json:"ref-description" constraints:"required=true;nillable=false"`
	RefSummary     string `json:"ref-summary" constraints:"required=true;nillable=false"`
}

func TestNested(t *testing.T) {
	msg := Example{
		Reference: Reference{
			Ref:            "reference",
			RefSummary:     "ref summary",
			RefDescription: "ref description",
		},
		Summary:     "summary",
		Description: "description",
		Value:       nil,
	}

	if err := sv.Validate(msg); err != nil {
		t.Errorf("Error in validation: %s", err)
	}
}

type ExampleFail struct {
	ReferenceFail
	Summary     string      `json:"summary,omitempty" constraints:"required=true;nillable=false"`
	Description string      `json:"description,omitempty" constraints:"required=true;nillable=false"`
	Value       interface{} `json:"example,omitempty" constraints:"required=true;nillable=false"`
}

type ReferenceFail struct {
	Ref            string `json:"ref" constraints:"required=true;nillable=false;min-length=10"`
	RefDescription string `json:"ref-description" constraints:"required=true;nillable=false"`
	RefSummary     string `json:"ref-summary" constraints:"required=true;nillable=false"`
}

func TestNestedFail(t *testing.T) {
	msg := ExampleFail{
		ReferenceFail: ReferenceFail{
			Ref:            "reference",
			RefSummary:     "ref summary",
			RefDescription: "ref description",
		},
		Summary:     "summary",
		Description: "description",
		Value:       nil,
	}

	err := sv.Validate(msg)
	got := err.Error()
	want := "min-length validation failed"
	if got != want {
		t.Errorf("Expected: %s, got: %s", got, want)
	}
}

/**
Empty Struct Validation
// TODO
*/

type EmptyExample struct {
	Field string `json:"field" constraints:"required=false;nillable=false"`
}

func TestEmptyStruct(t *testing.T) {
	msg := EmptyExample{}
	if err := sv.Validate(msg); err != nil {
		t.Errorf("Error in validation: %s", err)
	}

}

type ConstExample struct {
	Summary     string `json:"summary" constraints:"-"`
	Description string `json:"description" constraints:"required=false;nillable=false"`
}

func TestConstStruct(t *testing.T) {
	msg := ConstExample{
		Summary:     "testing",
		Description: "this is testing",
	}
	if err := sv.Validate(msg); err != nil {
		t.Errorf("Error in validation: %s", err)
	}
}

func TestEnumValidation(t *testing.T) {

	testsPass := []struct {
		Name  string
		input interface{}
	}{
		{
			Name: "Test1",
			input: struct {
				Status     string `json:"status" constraints:"required=true;nillable=true;enum=Success,Error,Not Reachable"`
				StatusCode int    `json:"statusCode" constraints:"required=true;nillable=true;enum=200,404,500"`
			}{Status: "Success", StatusCode: 200},
		},
	}

	for _, tt := range testsPass {
		t.Run(tt.Name, func(t *testing.T) {
			err := sv.Validate(tt.input)
			if err != nil {
				t.Errorf("Error in validation: %s", err)
			}
		})
	}

	testsError := []struct {
		Name  string
		input interface{}
		want  string
	}{
		{
			Name: "Test2",
			input: struct {
				Status2     string `json:"status2" constraints:"required=true;nillable=true;enum=Success,Error,Not Reachable"`
				StatusCode2 int    `json:"statusCode2" constraints:"required=true;nillable=true;enum=200,404,500"`
			}{Status2: "Success", StatusCode2: 503},
			want: "enum validation failed",
		},
	}

	for _, tt := range testsError {
		t.Run(tt.Name, func(t *testing.T) {
			err := sv.Validate(tt.input)
			if tt.want != err.Error() {
				t.Errorf("Got: %s, want: %s", err, tt.want)
			}
		})
	}
}
