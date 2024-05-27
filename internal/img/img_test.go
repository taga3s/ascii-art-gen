package img

import (
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name string
		path string
	}{
		{
			name: "valid png image",
			path: "../../examples/tests/cat.png",
		},
		{
			name: "valid jpg image",
			path: "../../examples/tests/cat.jpg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Load(tt.path)
			if err != nil {
				t.Errorf("Load() error = %v", err)
			}
		})
	}
}

func TestUnSync(t *testing.T) {
	tests := []struct {
		name string
		path string
	}{
		{
			name: "valid png image",
			path: "../../examples/tests/cat.png",
		},
		{
			name: "valid jpg image",
			path: "../../examples/tests/cat.jpg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dest, _ := Load(tt.path)

			if err := UnSync(dest); err != nil {
				t.Errorf("UnSync() error = %v", err)
			}
		})
	}
}
