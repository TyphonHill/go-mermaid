package classdiagram

import (
	"reflect"
	"testing"
)

func TestNewClass(t *testing.T) {
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
				name: "TestClass",
			},
			wantNewClass: &Class{Name: "TestClass"},
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

func TestClass_AddMethod(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name          string
		class         *Class
		args          args
		wantNewMethod *Method
	}{
		{
			name:          "Nominal test",
			class:         NewClass("TestClass"),
			args:          args{name: "TestMethod"},
			wantNewMethod: &Method{Name: "TestMethod", Visibility: MethodVisibilityPublic},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Class{
				Name:       tt.class.Name,
				Label:      tt.class.Label,
				Annotation: tt.class.Annotation,
				methods:    tt.class.methods,
				fields:     tt.class.fields,
			}
			if gotNewMethod := c.AddMethod(tt.args.name); !reflect.DeepEqual(gotNewMethod, tt.wantNewMethod) {
				t.Errorf("Class.AddMethod() = %v, want %v", gotNewMethod, tt.wantNewMethod)
			}
		})
	}
}

func TestClass_AddField(t *testing.T) {
	type args struct {
		fieldName string
		fieldType string
	}
	tests := []struct {
		name         string
		class        *Class
		args         args
		wantNewField *Field
	}{
		{
			name:         "Nominal test",
			class:        NewClass("TestClass"),
			args:         args{fieldName: "TestField", fieldType: "int32"},
			wantNewField: &Field{Name: "TestField", Type: "int32", Visibility: FieldVisibilityPublic},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Class{
				Name:       tt.class.Name,
				Label:      tt.class.Label,
				Annotation: tt.class.Annotation,
				methods:    tt.class.methods,
				fields:     tt.class.fields,
			}
			if gotNewField := c.AddField(tt.args.fieldName, tt.args.fieldType); !reflect.DeepEqual(gotNewField, tt.wantNewField) {
				t.Errorf("Class.AddField() = %v, want %v", gotNewField, tt.wantNewField)
			}
		})
	}
}

func TestClass_String(t *testing.T) {
	class := NewClass("TestClass")
	class.Annotation = ClassAnnotationAbstract
	class.Label = "TestLabel"
	class.AddField("TestField1", "int32")
	class.AddField("TestField2", "string")
	m1 := class.AddMethod("Method1")
	m1.ReturnType = "float32"
	m1.Visibility = MethodVisibilityProtected
	m1.AddParameter("Param1", "string")
	m1.AddParameter("Param2", "int32")
	class.AddMethod("Method2")

	type args struct {
		curIndentation string
	}
	tests := []struct {
		name  string
		class *Class
		args  args
		want  string
	}{
		{
			name:  "Nominal test",
			class: NewClass("TestClass"),
			args:  args{curIndentation: "%s"},
			want: `	class TestClass{
	}
`,
		},
		{
			name:  "Complex class",
			class: class,
			args:  args{curIndentation: "%s"},
			want: `	class TestClass["TestLabel"]{
		<<Abstract>>
		+int32 TestField1
		+string TestField2
		#Method1(Param1:string,Param2:int32) float32
		+Method2() 
	}
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Class{
				Name:       tt.class.Name,
				Label:      tt.class.Label,
				Annotation: tt.class.Annotation,
				methods:    tt.class.methods,
				fields:     tt.class.fields,
			}
			if got := c.String(tt.args.curIndentation); got != tt.want {
				t.Errorf("Class.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
