package basediagram

import (
	"strings"
	"testing"
)

func TestNewMarkdownFencer(t *testing.T) {
	got := NewMarkdownFencer()
	if got.markdownFence {
		t.Error("NewMarkdownFencer() should create with fence disabled")
	}
}

func TestMarkdownFencer_EnableMarkdownFence(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "Enable markdown fence",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fencer := NewMarkdownFencer()
			result := fencer.EnableMarkdownFence()

			if !fencer.markdownFence {
				t.Error("EnableMarkdownFence() did not enable fence")
			}

			if result != &fencer {
				t.Error("EnableMarkdownFence() should return fencer for chaining")
			}
		})
	}
}

func TestMarkdownFencer_DisableMarkdownFence(t *testing.T) {
	tests := []struct {
		name  string
		setup func(*MarkdownFencer)
		want  bool
	}{
		{
			name: "Disable when enabled",
			setup: func(m *MarkdownFencer) {
				m.EnableMarkdownFence()
			},
			want: false,
		},
		{
			name:  "Disable when already disabled",
			setup: nil,
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fencer := NewMarkdownFencer()
			if tt.setup != nil {
				tt.setup(&fencer)
			}

			fencer.DisableMarkdownFence()

			if fencer.markdownFence != tt.want {
				t.Errorf("DisableMarkdownFence() = %v, want %v", fencer.markdownFence, tt.want)
			}
		})
	}
}

func TestMarkdownFencer_IsMarkdownFenceEnabled(t *testing.T) {
	tests := []struct {
		name  string
		setup func(*MarkdownFencer)
		want  bool
	}{
		{
			name:  "Check when disabled",
			setup: nil,
			want:  false,
		},
		{
			name: "Check when enabled",
			setup: func(m *MarkdownFencer) {
				m.EnableMarkdownFence()
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fencer := NewMarkdownFencer()
			if tt.setup != nil {
				tt.setup(&fencer)
			}

			if got := fencer.IsMarkdownFenceEnabled(); got != tt.want {
				t.Errorf("IsMarkdownFenceEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarkdownFencer_WrapWithFence(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		setup    func(*MarkdownFencer)
		contains []string
	}{
		{
			name:    "No fence when disabled",
			content: "test content",
			setup:   nil,
			contains: []string{
				"test content",
			},
		},
		{
			name:    "With fence when enabled",
			content: "test content",
			setup: func(m *MarkdownFencer) {
				m.EnableMarkdownFence()
			},
			contains: []string{
				"```mermaid\n",
				"test content",
				"\n```\n",
			},
		},
		{
			name:    "Empty content with fence",
			content: "",
			setup: func(m *MarkdownFencer) {
				m.EnableMarkdownFence()
			},
			contains: []string{
				"```mermaid\n",
				"\n```\n",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fencer := NewMarkdownFencer()
			if tt.setup != nil {
				tt.setup(&fencer)
			}

			got := fencer.WrapWithFence(tt.content)
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("WrapWithFence() missing expected content %q in:\n%s", want, got)
				}
			}
		})
	}
}
