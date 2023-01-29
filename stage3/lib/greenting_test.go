package lib_test

import (
	"github.com/kamikazezirou/bazel-monorepo/stage3/lib"
	"testing"
)

func TestGreeting(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test",
			want: "Hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lib.Greeting(); got != tt.want {
				t.Errorf("Greeting() = %v, want %v", got, tt.want)
			}
		})
	}
}
