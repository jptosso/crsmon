package crsmon

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestLoadPolicy(t *testing.T) {
	dir, err := ioutil.TempDir("/tmp", "crstest*")
	defer os.Remove(dir)
	if err != nil {
		t.Error(err)
	}
	policy := NewPolicy(dir)
	err = policy.Build()
	if err != nil {
		t.Error(err)
	}
	if _, err := os.Stat(path.Join(dir, "crs.conf")); err != nil {
		t.Error(err)
	}
}
