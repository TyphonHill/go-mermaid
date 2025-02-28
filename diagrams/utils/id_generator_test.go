package utils

import (
	"testing"
)

func TestNewIDGenerator(t *testing.T) {
	got := NewIDGenerator()
	if got == nil {
		t.Error("NewIDGenerator() returned nil")
	}
	if got.nextID != 0 {
		t.Errorf("NewIDGenerator() initial nextID = %v, want 0", got.nextID)
	}
}

func TestDefaultIDGenerator_NextID(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*DefaultIDGenerator)
		sequence []string
	}{
		{
			name: "Generate sequence from 0",
			sequence: []string{
				"0", "1", "2", "3", "4",
			},
		},
		{
			name: "Generate sequence after reset",
			setup: func(g *DefaultIDGenerator) {
				g.NextID() // consume some IDs
				g.NextID()
				g.Reset()
			},
			sequence: []string{
				"0", "1", "2",
			},
		},
		{
			name: "Generate sequence from non-zero",
			setup: func(g *DefaultIDGenerator) {
				g.nextID = 10
			},
			sequence: []string{
				"10", "11", "12",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generator := NewIDGenerator()
			if tt.setup != nil {
				tt.setup(generator)
			}

			for i, want := range tt.sequence {
				got := generator.NextID()
				if got != want {
					t.Errorf("NextID() iteration %d = %v, want %v", i, got, want)
				}
			}
		})
	}
}

func TestDefaultIDGenerator_Reset(t *testing.T) {
	tests := []struct {
		name  string
		setup func(*DefaultIDGenerator)
		want  int
	}{
		{
			name: "Reset unused generator",
			want: 0,
		},
		{
			name: "Reset after generating IDs",
			setup: func(g *DefaultIDGenerator) {
				g.NextID()
				g.NextID()
				g.NextID()
			},
			want: 0,
		},
		{
			name: "Reset after manual set",
			setup: func(g *DefaultIDGenerator) {
				g.nextID = 100
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generator := NewIDGenerator()
			if tt.setup != nil {
				tt.setup(generator)
			}

			result := generator.Reset()

			// Test nextID value
			if generator.nextID != tt.want {
				t.Errorf("Reset() nextID = %v, want %v", generator.nextID, tt.want)
			}

			// Test method chaining
			if result != generator {
				t.Error("Reset() should return generator for chaining")
			}

			// Test next generated ID after reset
			got := generator.NextID()
			if got != "0" {
				t.Errorf("NextID() after Reset() = %v, want '0'", got)
			}
		})
	}
}

func TestIDGenerator_Interface(t *testing.T) {
	var _ IDGenerator = (*DefaultIDGenerator)(nil)
}
