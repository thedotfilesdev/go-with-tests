package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := RepeatCharacter("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatCharacter("a", 5)
	}
}

func ExampleRepeat() {
	repeated := RepeatCharacter("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
