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
	got, err := play.Publish(input[1:])
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestPublishWithNoArgsReturnsError(t *testing.T) {
	t.Parallel()
	_, err := play.Publish([]string{})
	if err == nil {
		t.Fatal("want error for zero args, but got nil")
	}
}