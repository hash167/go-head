package internal

// Write a test for the ReadFile function
// Use the following test cases:
// 1. Test the function with a file that does not exist
// 2. Test the function with a file that has 10 lines
// 3. Test the function with a file that has 10 bytes

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		n        int
		isByte   bool
		want     string
	}{
		{
			name:     "Test the function with a file that does not exist",
			filename: "testdata/nonexistent.txt",
			n:        0,
			isByte:   false,
			want:     "",
		},
		{
			name:     "Test the function with a file that has 10 lines",
			filename: "testdata/10lines.txt",
			n:        10,
			isByte:   false,
			want:     "1\n2\n3\n4\n5\n6\n7\n8\n9\n10",
		},
		{
			name:     "Test the function with a file that has 10 bytes",
			filename: "testdata/10bytes.txt",
			n:        10,
			isByte:   true,
			want:     "1234567890",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadFile(tt.filename, tt.n, tt.isByte); got != tt.want {
				t.Errorf("ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
