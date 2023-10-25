package flowchart

import (
	"reflect"
	"testing"
)

func TestNewClass(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	type args struct {
		name string
	}
	tests := []struct {
		name         string
		args         args
		wantNewClass *Class
	}{
		{
			name: "Nominal test",
			args: args{
				name: "TestName",
			},
			wantNewClass: &Class{
				Name: "TestName",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewClass := NewClass(tt.args.name); !reflect.DeepEqual(gotNewClass, tt.wantNewClass) {
				t.Errorf("NewClass() = %v, want %v", gotNewClass, tt.wantNewClass)
			}
		})
	}
}

func TestClass_String(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	tests := []struct {
		name  string
		class *Class
		want  string
	}{
		{
			name: "Nominal test",
			class: &Class{
				Name: "TestName",
				Style: &NodeStyle{
					Color:       "red",
					StrokeWidth: 5,
				},
			},
			want: "\tclassDef TestName color:red,stroke-width:5\n",
		},
		{
			name: "No style defined",
			class: &Class{
				Name: "Test Name",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.class.String(); got != tt.want {
				t.Errorf("Class.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
