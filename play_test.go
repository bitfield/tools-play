package play_test

import (
	"play"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPublish(t *testing.T) {
	t.Parallel()
	input := []string{"/usr/bin/play", "a", "b", "c"}
	want := "a"
	var got string
	got, err := play.Publish(input)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

