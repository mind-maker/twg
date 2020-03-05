package underscore

import "testing"

// func TestCamel(t *testing.T) {
// 	testCases := []struct {
// 		arg  string
// 		want string
// 	}{
// 		{"thisIsACamelCaseString", "this_is_a_camel_case_string"},
// 		{"with a space", "with a space"},
// 		{"endsWithA", "ends_with_a"},
// 	}
// 	for _, tc := range testCases {
// 		arg := tc.arg
// 		want := tc.want
// 		got := Camel(arg)
// 		if got != want {
// 			t.Errorf("Camel(%q) = %q; want %q", arg, got, want)
// 		}
// 	}
// }

func TestCamel(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"simple", args{"thisIsACamelCaseString"}, "this_is_a_camel_case_string"},
		{"spaces", args{"with a space"}, "with a space"},
		{"ends with capital", args{"endsWithA"}, "ends_with_a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Camel(tt.args.str); got != tt.want {
				t.Errorf("Camel() = %v, want %v", got, tt.want)
			}
		})
	}
}
