package basediagram

import (
	"strings"
	"testing"
)

func TestBaseProperty_Format(t *testing.T) {
	tests := []struct {
		name     string
		property BaseProperty
		contains []string
	}{
		{
			name: "Format string value",
			property: BaseProperty{
				Name: "testProp",
				Val:  "test value",
			},
			contains: []string{
				"    testProp: test value",
			},
		},
		{
			name: "Format integer value",
			property: BaseProperty{
				Name: "testProp",
				Val:  42,
			},
			contains: []string{
				"    testProp: 42",
			},
		},
		{
			name: "Format boolean value",
			property: BaseProperty{
				Name: "testProp",
				Val:  true,
			},
			contains: []string{
				"    testProp: true",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.property.Format()
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("Format() missing expected content %q in:\n%s", want, got)
				}
			}
		})
	}
}

func TestBaseProperty_Value(t *testing.T) {
	tests := []struct {
		name     string
		property BaseProperty
		want     interface{}
	}{
		{
			name: "Get string value",
			property: BaseProperty{
				Name: "testProp",
				Val:  "test value",
			},
			want: "test value",
		},
		{
			name: "Get integer value",
			property: BaseProperty{
				Name: "testProp",
				Val:  42,
			},
			want: 42,
		},
		{
			name: "Get boolean value",
			property: BaseProperty{
				Name: "testProp",
				Val:  true,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.property.Value()
			if got != tt.want {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringArrayProperty_Format(t *testing.T) {
	tests := []struct {
		name     string
		property StringArrayProperty
		contains []string
	}{
		{
			name: "Format empty array",
			property: StringArrayProperty{
				BaseProperty: BaseProperty{
					Name: "testProp",
					Val:  []string{},
				},
			},
			contains: []string{
				"    testProp: []",
			},
		},
		{
			name: "Format single value array",
			property: StringArrayProperty{
				BaseProperty: BaseProperty{
					Name: "testProp",
					Val:  []string{"value1"},
				},
			},
			contains: []string{
				`    testProp: ["value1"]`,
			},
		},
		{
			name: "Format multiple value array",
			property: StringArrayProperty{
				BaseProperty: BaseProperty{
					Name: "testProp",
					Val:  []string{"value1", "value2", "value3"},
				},
			},
			contains: []string{
				`    testProp: ["value1", "value2", "value3"]`,
			},
		},
		{
			name: "Format array with special characters",
			property: StringArrayProperty{
				BaseProperty: BaseProperty{
					Name: "testProp",
					Val:  []string{"value:1", "value,2", "value\"3"},
				},
			},
			contains: []string{
				`    testProp: ["value:1", "value,2", "value\"3"]`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.property.Format()
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("Format() missing expected content %q in:\n%s", want, got)
				}
			}
		})
	}
}

func TestPropertyTypes(t *testing.T) {
	tests := []struct {
		name     string
		property DiagramProperty
		value    interface{}
		contains []string
	}{
		{
			name: "Bool property",
			property: &BoolProperty{
				BaseProperty: BaseProperty{
					Name: "testBool",
					Val:  true,
				},
			},
			value: true,
			contains: []string{
				"    testBool: true",
			},
		},
		{
			name: "Int property",
			property: &IntProperty{
				BaseProperty: BaseProperty{
					Name: "testInt",
					Val:  42,
				},
			},
			value: 42,
			contains: []string{
				"    testInt: 42",
			},
		},
		{
			name: "Float property",
			property: &FloatProperty{
				BaseProperty: BaseProperty{
					Name: "testFloat",
					Val:  3.14,
				},
			},
			value: 3.14,
			contains: []string{
				"    testFloat: 3.14",
			},
		},
		{
			name: "String property",
			property: &StringProperty{
				BaseProperty: BaseProperty{
					Name: "testString",
					Val:  "test",
				},
			},
			value: "test",
			contains: []string{
				"    testString: test",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test Format()
			got := tt.property.Format()
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("Format() missing expected content %q in:\n%s", want, got)
				}
			}

			// Test Value()
			if got := tt.property.Value(); got != tt.value {
				t.Errorf("Value() = %v, want %v", got, tt.value)
			}
		})
	}
}
