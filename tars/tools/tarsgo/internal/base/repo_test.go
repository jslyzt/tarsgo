package base

import (
	"context"
	"testing"
)

func TestRepo(t *testing.T) {
	r := NewRepo("https://github.com/jslyzt/tarsgo.git", "", "")
	if err := r.Clone(context.Background()); err != nil {
		t.Fatal(err)
	}
	if err := r.CopyTo(context.Background(), nil, "/tmp/test_repo", "github.com/jslyzt/tarsgo", nil); err != nil {
		t.Fatal(err)
	}
}
