package tests

import (
	"testing"

	"github.com/olachat/gola/golalib/testdata/myusers"
	"github.com/volatiletech/null/v8"
)

func TestMyUser(t *testing.T) {
	u := myusers.New()

	if u.GetId() != 0 {
		t.Error("myuser pk init failed")
	}

	u.SetName(null.StringFrom("my name"))

	err := u.Insert()
	if err != nil {
		t.Error(err)
	}
	if u.GetId() != 1 {
		t.Error("Insert myuser pk failed")
	}

	u = myusers.FetchByPK(1)
	if u.GetName().String != "my name" {
		t.Error("myuser get name failed")
	}
}
