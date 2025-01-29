package class

import (
	"strings"
	"testing"
)

func TestNewMethod(t *testing.T) {
	tests := []struct {
		name       string
		methodName string
		want       *Method
	}{
		{
			name:       "Create new method",
			methodName: "testMethod",
			want: &Method{
				Name:       "testMethod",
				Parameters: []Parameter{},
				ReturnType: "",
				Visibility: MethodVisibilityPublic,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMethod(tt.methodName)

			if got.Name != tt.want.Name {
				t.Errorf("NewMethod() Name = %v, want %v", got.Name, tt.want.Name)
			}

			if got.Visibility != tt.want.Visibility {
				t.Errorf("NewMethod() Visibility = %v, want %v", got.Visibility, tt.want.Visibility)
			}

			// Check that parameters are initialized as empty slice
			if len(got.Parameters) != 0 {
				t.Errorf("NewMethod() did not initialize empty Parameters slice")
			}
		})
	}
}

func TestMethod_AddParameter(t *testing.T) {
	method := NewMethod("testMethod")

	method.AddParameter("param1", "string")
	method.AddParameter("param2", "int")

	if len(method.Parameters) != 2 {
		t.Errorf("AddParameter() did not add parameters to method")
	}

	if method.Parameters[0].Name != "param1" || method.Parameters[0].Type != "string" {
		t.Errorf("First parameter incorrect. Got %v:%v, want param1:string",
			method.Parameters[0].Name, method.Parameters[0].Type)
	}

	if method.Parameters[1].Name != "param2" || method.Parameters[1].Type != "int" {
		t.Errorf("Second parameter incorrect. Got %v:%v, want param2:int",
			method.Parameters[1].Name, method.Parameters[1].Type)
	}
}

func TestMethod_String(t *testing.T) {
	tests := []struct {
		name     string
		method   *Method
		contains []string
	}{
		{
			name: "Public method with no params",
			method: func() *Method {
				m := NewMethod("simpleMethod")
				return m
			}(),
			contains: []string{
				"+simpleMethod()",
			},
		},
		{
			name: "Method with parameters",
			method: func() *Method {
				m := NewMethod("calculateSum")
				m.AddParameter("a", "int")
				m.AddParameter("b", "int")
				m.ReturnType = "int"
				return m
			}(),
			contains: []string{
				"+calculateSum(a:int,b:int) int",
			},
		},
		{
			name: "Method with different visibility",
			method: func() *Method {
				m := NewMethod("privateMethod")
				m.Visibility = MethodVisibilityPrivate
				return m
			}(),
			contains: []string{
				"-privateMethod()",
			},
		},
		{
			name: "Static method",
			method: func() *Method {
				m := NewMethod("staticMethod")
				m.Classifier = MethodClassifierStatic
				return m
			}(),
			contains: []string{
				"+staticMethod()$",
			},
		},
		{
			name: "Abstract method",
			method: func() *Method {
				m := NewMethod("abstractMethod")
				m.Classifier = MethodClassifierAbstract
				m.ReturnType = "void"
				return m
			}(),
			contains: []string{
				"+abstractMethod()*",
				"void",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := tt.method.String()

			for _, expectedContent := range tt.contains {
				if !strings.Contains(output, expectedContent) {
					t.Errorf("String() output missing expected content: %q", expectedContent)
				}
			}
		})
	}
}
