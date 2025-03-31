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

	if u.GetAge().Valid {
		t.Error("Age not null")
	}

	u = myusers.FetchByPK(1)
	if u.GetName().String != "my name" {
		t.Error("myuser get name failed")
	}

	u = myusers.New()
	u.SetAge(null.IntFrom(25))
	u.SetName(null.StringFrom("Age 25"))
	err = u.Insert()
	if err != nil {
		t.Error(err)
	}

	users := myusers.Select().WhereAgeEQ(25).Limit(0, 1)
	if len(users) != 1 {
		t.Error("Myuser lookup error")
	}

	u = users[0]
	if !u.GetAge().Valid {
		t.Error("Age null")
	}

	if u.GetAge().Int != 25 || u.GetName().String != "Age 25" {
		t.Error("Myuser select error")
	}

	users2 := myusers.SelectFields[struct {
		myusers.Id
		myusers.Age
	}]().WhereAgeIN(25, 20).All()

	if len(users2) != 1 {
		t.Error("Myuser select fields error")
	}

	u2 := users2[0]
	if !u2.GetAge().Valid {
		t.Error("Age null")
	}

	if u2.GetAge().Int != 25 || u2.GetId() != 2 {
		t.Error("Myuser select fields value error")
	}
}
