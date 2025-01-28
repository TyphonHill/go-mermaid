package class_diagram

import (
	"reflect"
	"testing"
)

func TestNewNamespace(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name             string
		args             args
		wantNewNamespace *Namespace
	}{
		{
			name: "Nominal test",
			args: args{
				name: "TestNamespace",
			},
			wantNewNamespace: &Namespace{
				Name: "TestNamespace",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewNamespace := NewNamespace(tt.args.name); !reflect.DeepEqual(gotNewNamespace, tt.wantNewNamespace) {
				t.Errorf("NewNamespace() = %v, want %v", gotNewNamespace, tt.wantNewNamespace)
			}
		})
	}
}

func TestNamespace_AddClass(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name         string
		namespace    *Namespace
		args         args
		wantNewClass *Class
	}{
		{
			name:      "Nominal test",
			namespace: NewNamespace("TestNamespace"),
			args: args{
				name: "TestClass",
			},
			wantNewClass: &Class{
				Name: "TestClass",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Namespace{
				Name:    tt.namespace.Name,
				classes: tt.namespace.classes,
			}
			if gotNewClass := n.AddClass(tt.args.name); !reflect.DeepEqual(gotNewClass, tt.wantNewClass) {
				t.Errorf("Namespace.AddClass() = %v, want %v", gotNewClass, tt.wantNewClass)
			}
		})
	}
}

func TestNamespace_String(t *testing.T) {
	ns := NewNamespace("TestNamespace")
	ns.AddClass("Test1")
	ns.AddClass("Test2")

	tests := []struct {
		name      string
		namespace *Namespace
		want      string
	}{
		{
			name:      "Nominal test with no classes",
			namespace: NewNamespace("TestNamespace"),
			want:      "",
		},
		{
			name:      "Namespace with classes",
			namespace: ns,
			want: `	namespace TestNamespace{
		class Test1{
		}
		class Test2{
		}
	}
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Namespace{
				Name:    tt.namespace.Name,
				classes: tt.namespace.classes,
			}
			if got := n.String(); got != tt.want {
				t.Errorf("Namespace.String() = \n%v\n, want \n%v\n", got, tt.want)
			}
		})
	}
}
