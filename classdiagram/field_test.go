package classdiagram

import (
	"reflect"
	"testing"
)

func TestNewField(t *testing.T) {
	type args struct {
		fieldName string
		fieldType string
	}
	tests := []struct {
		name         string
		args         args
		wantNewField *Field
	}{
		{
			name: "Nominal test",
			args: args{
				fieldName: "testField",
				fieldType: "string",
			},
			wantNewField: &Field{
				Name:       "testField",
				Type:       "string",
				Visibility: FieldVisibilityPublic,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewField := NewField(tt.args.fieldName, tt.args.fieldType); !reflect.DeepEqual(gotNewField, tt.wantNewField) {
				t.Errorf("NewField() = %v, want %v", gotNewField, tt.wantNewField)
			}
		})
	}
}

func TestField_String(t *testing.T) {
	field := NewField("NewField", "int32")
	field.Classifier = FieldClassifierStatic
	field.Visibility = FieldVisibilityProtected

	tests := []struct {
		name  string
		field *Field
		want  string
	}{
		{
			name:  "Nominal test",
			field: NewField("FieldName", "string"),
			want:  `	+string FieldName`,
		},
		{
			name:  "Field with modifiers",
			field: field,
			want:  `	#int32 NewField$`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Field{
				Name:       tt.field.Name,
				Type:       tt.field.Type,
				Visibility: tt.field.Visibility,
				Classifier: tt.field.Classifier,
			}
			if got := m.String(); got != tt.want {
				t.Errorf("Field.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
