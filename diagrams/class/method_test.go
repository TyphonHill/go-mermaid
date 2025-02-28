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

func TestMethod_SetVisibility(t *testing.T) {
	tests := []struct {
		name       string
		method     *Method
		visibility methodVisibility
		want       methodVisibility
	}{
		{
			name:       "Set public visibility",
			method:     NewMethod("test"),
			visibility: MethodVisibilityPublic,
			want:       MethodVisibilityPublic,
		},
		{
			name:       "Set private visibility",
			method:     NewMethod("test"),
			visibility: MethodVisibilityPrivate,
			want:       MethodVisibilityPrivate,
		},
		{
			name:       "Set protected visibility",
			method:     NewMethod("test"),
			visibility: MethodVisibilityProtected,
			want:       MethodVisibilityProtected,
		},
		{
			name:       "Set internal visibility",
			method:     NewMethod("test"),
			visibility: MethodVisibilityInternal,
			want:       MethodVisibilityInternal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.method.SetVisibility(tt.visibility)

			if result != tt.method {
				t.Error("SetVisibility() should return method for chaining")
			}

			if tt.method.Visibility != tt.want {
				t.Errorf("SetVisibility() = %v, want %v", tt.method.Visibility, tt.want)
			}
		})
	}
}

func TestMethod_SetReturnType(t *testing.T) {
	tests := []struct {
		name       string
		method     *Method
		returnType string
		want       string
	}{
		{
			name:       "Set void return type",
			method:     NewMethod("test"),
			returnType: "void",
			want:       "void",
		},
		{
			name:       "Set string return type",
			method:     NewMethod("test"),
			returnType: "string",
			want:       "string",
		},
		{
			name:       "Set custom type",
			method:     NewMethod("test"),
			returnType: "MyClass",
			want:       "MyClass",
		},
		{
			name:       "Change return type",
			method:     NewMethod("test"),
			returnType: "int",
			want:       "int",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.method.SetReturnType(tt.returnType)

			if result != tt.method {
				t.Error("SetReturnType() should return method for chaining")
			}

			if tt.method.ReturnType != tt.want {
				t.Errorf("SetReturnType() = %v, want %v", tt.method.ReturnType, tt.want)
			}
		})
	}
}

func TestMethod_SetClassifier(t *testing.T) {
	tests := []struct {
		name       string
		method     *Method
		classifier methodClassifier
		want       methodClassifier
	}{
		{
			name:       "Set abstract classifier",
			method:     NewMethod("test"),
			classifier: MethodClassifierAbstract,
			want:       MethodClassifierAbstract,
		},
		{
			name:       "Set static classifier",
			method:     NewMethod("test"),
			classifier: MethodClassifierStatic,
			want:       MethodClassifierStatic,
		},
		{
			name:       "Change classifier",
			method:     NewMethod("test"),
			classifier: MethodClassifierStatic,
			want:       MethodClassifierStatic,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.method.SetClassifier(tt.classifier)

			if result != tt.method {
				t.Error("SetClassifier() should return method for chaining")
			}

			if tt.method.Classifier != tt.want {
				t.Errorf("SetClassifier() = %v, want %v", tt.method.Classifier, tt.want)
			}
		})
	}
}
