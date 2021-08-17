package play_test

import (
	"play"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPublish(t *testing.T) {
	t.Parallel()
	input := []string{"/usr/bin/play", "a", "b", "c"}
	want := "a b c"
	var got string
	got = play.Publish(input[1:])
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestPublishWithNoArgsReturnsEmptyString(t *testing.T) {
	t.Parallel()
	want := ""
	got := play.Publish([]string{})
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}