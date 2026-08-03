package account

import (
	"bytes"
	"testing"
)

func TestVerificationUploadRequest_Encode(t *testing.T) {
	t.Parallel()

	req := VerificationUploadRequest{
		ImageFile1:     bytes.NewBufferString("front"),
		ImageFile1Name: "front.jpg",
		ImageFile2:     bytes.NewBufferString("back"),
	}
	fields := req.Encode()
	if len(fields) != 2 {
		t.Fatalf("len(Encode()) = %d, want 2", len(fields))
	}

	tests := []struct {
		name          string
		index         int
		expectedKey   string
		expectedValue string
	}{
		{
			name:          "primary document",
			index:         0,
			expectedKey:   "image_file1",
			expectedValue: "front.jpg",
		},
		{
			name:          "secondary document uses default filename",
			index:         1,
			expectedKey:   "image_file2",
			expectedValue: "image_file2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := fields[tt.index]
			if field.Key != tt.expectedKey {
				t.Errorf("Key = %q, want %q", field.Key, tt.expectedKey)
			}
			if field.Value != tt.expectedValue {
				t.Errorf("Value = %q, want %q", field.Value, tt.expectedValue)
			}
			if field.Reader == nil {
				t.Error("Reader is nil")
			}
		})
	}
}
