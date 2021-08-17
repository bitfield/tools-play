package play_test

import (
	"play"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPublish(t *testing.T) {
	t.Parallel()
	p, err := play.NewPublisherFromArgs([]string{"hello", "world"})
	if err != nil {
		t.Fatal(err)
	}
	want := "hello world"
	got := p.Publish()
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestPublishWithNoArgsReturnsEmptyString(t *testing.T) {
	t.Parallel()
	want := ""
	p, err := play.NewPublisherFromArgs([]string{})
	if err != nil {
		t.Fatal(err)
	}
	got := p.Publish()
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestNewPublisherFromArgs(t *testing.T) {
	t.Parallel()
	_, err := play.NewPublisherFromArgs([]string{"hello", "world"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestPublishAllCaps(t *testing.T) {
	t.Parallel()
	want := "HELLO WORLD"
	publisher, err := play.NewPublisherFromArgs([]string{"-c", "hello", "world"})
	if err != nil {
		t.Fatal(err)
	}
	got := publisher.Publish()
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}