package class_diagram

import (
	"reflect"
	"testing"
)

func TestNewMethod(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name          string
		args          args
		wantNewMethod *Method
	}{
		{
			name: "Nominal test",
			args: args{
				name: "NewMethod",
			},
			wantNewMethod: &Method{Name: "NewMethod", Visibility: MethodVisibilityPublic},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewMethod := NewMethod(tt.args.name); !reflect.DeepEqual(gotNewMethod, tt.wantNewMethod) {
				t.Errorf("NewMethod() = %v, want %v", gotNewMethod, tt.wantNewMethod)
			}
		})
	}
}

func TestMethod_AddParameter(t *testing.T) {
	type args struct {
		paramName string
		paramType string
	}
	tests := []struct {
		name   string
		method *Method
		args   args
	}{
		{
			name:   "Nominal test",
			method: NewMethod("TestMethod"),
			args: args{
				paramName: "testParam",
				paramType: "int32",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Method{
				Name:       tt.method.Name,
				Parameters: tt.method.Parameters,
				ReturnType: tt.method.ReturnType,
				Visibility: tt.method.Visibility,
				Classifier: tt.method.Classifier,
			}

			m.AddParameter(tt.args.paramName, tt.args.paramType)
		})
	}
}

func TestMethod_String(t *testing.T) {
	method := NewMethod("TestMethod")
	method.ReturnType = "float32"
	method.Visibility = MethodVisibilityProtected
	method.AddParameter("Param1", "string")
	method.AddParameter("Param2", "int32")

	tests := []struct {
		name   string
		method *Method
		want   string
	}{
		{
			name:   "Nominal test",
			method: NewMethod("TestMethod"),
			want:   `	+TestMethod() `,
		},
		{
			name:   "Method with params",
			method: method,
			want:   `	#TestMethod(Param1:string,Param2:int32) float32`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Method{
				Name:       tt.method.Name,
				Parameters: tt.method.Parameters,
				ReturnType: tt.method.ReturnType,
				Visibility: tt.method.Visibility,
				Classifier: tt.method.Classifier,
			}
			if got := m.String(); got != tt.want {
				t.Errorf("Method.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
