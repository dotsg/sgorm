package tests

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/olachat/gola/golalib/testdata/users"
)

type SimpleUser struct {
	users.Name
	users.Email
}

func TestUserInsert(t *testing.T) {
	u := users.NewWithPK(11)
	u.SetEmail("hello")
	u.SetName("maou sheng")
	u.SetCreatedAt(111)
	u.SetUpdatedAt(222)
	u.Insert()
}

func TestUserDouble(t *testing.T) {
	u1 := users.FetchByPK(1)
	if u1.GetDoubleType() != 1.8729 {
		t.Errorf("FetchUserByPK GetDoubleType returns unexpected value: %f", u1.GetDoubleType())
	}
	if u1.GetFloatType() != 1.55555 {
		t.Errorf("FetchUserByPK GetFloatType returns unexpected value: %f", u1.GetFloatType())
	}

	u2 := users.FetchByPK(2)
	if u2.GetDoubleType() != 2.8239 {
		t.Errorf("FetchUserByPK GetDoubleType returns unexpected value: %f", u2.GetDoubleType())
	}
	if u2.GetFloatType() != 2.5 {
		t.Errorf("FetchUserByPK GetFloatType returns unexpected value: %f", u2.GetFloatType())
	}

	u3 := users.FetchByPK(3)
	if u3.GetDoubleType() != 334.8593 {
		t.Errorf("FetchUserByPK GetDoubleType returns unexpected value: %f", u3.GetDoubleType())
	}
	if u3.GetFloatType() != 3.5 {
		t.Errorf("FetchUserByPK GetFloatType returns unexpected value: %f", u3.GetFloatType())
	}

	u4 := users.FetchByPK(4)
	if u4.GetDoubleType() != 42234.83 {
		t.Errorf("FetchUserByPK GetDoubleType returns unexpected value: %f", u4.GetDoubleType())
	}
	if u4.GetFloatType() != 4.5 {
		t.Errorf("FetchUserByPK GetFloatType returns unexpected value: %f", u4.GetFloatType())
	}

	u4.SetDoubleType(5.1)
	u4.SetFloatType(4.0)
	u4.Update()

	u5 := users.FetchByPK(4)
	if u5.GetDoubleType() != 5.1 {
		t.Errorf("FetchUserByPK GetDoubleType returns unexpected value: %f", u4.GetDoubleType())
	}
	if u5.GetFloatType() != 4.0 {
		t.Errorf("FetchUserByPK GetFloatType returns unexpected value: %f", u4.GetFloatType())
	}
}

func TestUserHobby(t *testing.T) {
	u1 := users.FetchByPK(1)
	if u1.GetHobby() != users.UserHobbyRunning {
		t.Errorf("FetchUserByPK GetHobby returns unexpected value: %v", u1.GetHobby())
	}
	if u1.GetHobbyNoDefault() != users.UserHobbyNoDefaultSwimming {
		t.Errorf("FetchUserByPK GetHobbyNoDefault returns unexpected value: %v", u1.GetHobbyNoDefault())
	}

	u1.SetHobby(users.UserHobbySinging)
	u1.Update()
	u1 = users.FetchByPK(1)
	if u1.GetHobby() != users.UserHobbySinging {
		t.Errorf("FetchUserByPK GetHobby returns unexpected value: %v", u1.GetHobby())
	}

	u2 := users.FetchByPK(2)
	if u2.GetHobby() != users.UserHobbySwimming {
		t.Errorf("FetchUserByPK GetHobby returns unexpected value: %v", u2.GetHobby())
	}
	if u2.GetHobbyNoDefault() != users.UserHobbyNoDefaultRunning {
		t.Errorf("FetchUserByPK GetHobbyNoDefault returns unexpected value: %v", u2.GetHobbyNoDefault())
	}

	u3 := users.FetchByPK(3)
	if u3.GetHobby() != users.UserHobbySinging {
		t.Errorf("FetchUserByPK GetHobby returns unexpected value: %v", u3.GetHobby())
	}
	if u3.GetHobbyNoDefault() != users.UserHobbyNoDefaultSwimming {
		t.Errorf("FetchUserByPK GetHobbyNoDefault returns unexpected value: %v", u3.GetHobbyNoDefault())
	}

	u4 := users.FetchByPK(4)
	if u4.GetHobby() != users.UserHobbySinging {
		t.Errorf("FetchUserByPK GetHobby returns unexpected value: %v", u4.GetHobby())
	}
	if u4.GetHobbyNoDefault() != users.UserHobbyNoDefaultRunning {
		t.Errorf("FetchUserByPK GetHobbyNoDefault returns unexpected value: %v", u4.GetHobbyNoDefault())
	}
}

func TestUserSports(t *testing.T) {
	u1 := users.FetchByPK(1)
	if len(u1.GetSports()) != 1 {
		t.Errorf("FetchUserByPK GetSports returns unexpected value: %v", u1.GetSports())
	}
	if !contains(u1.GetSports(), users.UserSportsTennis) {
		t.Errorf("FetchUserByPK GetSports should contain swim. Actual: %v", u1.GetSports())
	}

	if len(u1.GetSportsNoDefault()) != 2 {
		t.Errorf("FetchUserByPK GetSportsNoDefault returns unexpected value: %v", u1.GetSportsNoDefault())
	}
	if !contains(u1.GetSportsNoDefault(), users.UserSportsNoDefaultSwim) {
		t.Errorf("FetchUserByPK GetSportsNoDefault should contain swim. Actual: %v", u1.GetSportsNoDefault())
	}
	if !contains(u1.GetSportsNoDefault(), users.UserSportsNoDefaultTennis) {
		t.Errorf("FetchUserByPK GetSportsNoDefault should contain tennis. Actual: %v", u1.GetSportsNoDefault())
	}

	u2 := users.FetchByPK(2)
	if len(u2.GetSports()) != 1 {
		t.Errorf("FetchUserByPK GetSports returns unexpected value: %v", u2.GetSports())
	}
	if !contains(u2.GetSports(), users.UserSportsFootball) {
		t.Errorf("FetchUserByPK GetSports should contain football. Actual: %v", u2.GetSports())
	}

	if len(u2.GetSportsNoDefault()) != 1 {
		t.Errorf("FetchUserByPK GetSportsNoDefault returns unexpected value: %v", u2.GetSportsNoDefault())
	}
	if !contains(u2.GetSportsNoDefault(), users.UserSportsNoDefaultBasketball) {
		t.Errorf("FetchUserByPK GetSportsNoDefault should contain basketball. Actual: %v", u2.GetSportsNoDefault())
	}
}

func TestUserSports2(t *testing.T) {
	u3 := users.FetchByPK(3)
	if len(u3.GetSports()) != 2 {
		t.Errorf("FetchUserByPK GetSports returns unexpected value: %v", u3.GetSports())
	}
	if !contains(u3.GetSports(), users.UserSportsSquash) {
		t.Errorf("FetchUserByPK GetSports should contain swim. Actual: %v", u3.GetSports())
	}
	if !contains(u3.GetSports(), users.UserSportsTennis) {
		t.Errorf("FetchUserByPK GetSports should contain football. Actual: %v", u3.GetSports())
	}

	if len(u3.GetSportsNoDefault()) != 2 {
		t.Errorf("FetchUserByPK GetSportsNoDefault returns unexpected value: %v", u3.GetSportsNoDefault())
	}
	if !contains(u3.GetSportsNoDefault(), users.UserSportsNoDefaultBadminton) {
		t.Errorf("FetchUserByPK GetSportsNoDefault should contain badminton. Actual: %v", u3.GetSportsNoDefault())
	}
	if !contains(u3.GetSportsNoDefault(), users.UserSportsNoDefaultSquash) {
		t.Errorf("FetchUserByPK GetSportsNoDefault should contain squash. Actual: %v", u3.GetSportsNoDefault())
	}

	u4 := users.FetchByPK(4)
	if len(u4.GetSports()) != 2 {
		t.Errorf("FetchUserByPK GetSports returns unexpected value: %v", u4.GetSports())
	}
	if !contains(u4.GetSports(), users.UserSportsBadminton) {
		t.Errorf("FetchUserByPK GetSports should contain swim. Actual: %v", u4.GetSports())
	}
	if !contains(u4.GetSports(), users.UserSportsBasketball) {
		t.Errorf("FetchUserByPK GetSports should contain football. Actual: %v", u4.GetSports())
	}

	if len(u4.GetSportsNoDefault()) != 1 {
		t.Errorf("FetchUserByPK GetSportsNoDefault returns unexpected value: %v", u4.GetSportsNoDefault())
	}
	if !contains(u4.GetSportsNoDefault(), users.UserSportsNoDefaultTennis) {
		t.Errorf("FetchUserByPK GetSportsNoDefault should contain tennis. Actual: %v", u4.GetSportsNoDefault())
	}
}

func contains[T comparable](slice []T, item T) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func TestUserMethods(t *testing.T) {
	u := users.FetchFieldsByPK[struct {
		users.Email
	}](1)
	if u.GetEmail() != "john@doe.com" {
		t.Error("Failed to FetchByPK with email using id 1")
	}

	u2 := users.FetchFieldsByPK[users.User](1)
	if u2.GetEmail() != "john@doe.com" && u2.GetName() != "John Doe" {
		t.Error("Failed to FetchByPK with User using id 1")
	}
	u2.SetEmail("joe@doe.com")
	u2.SetName("Joe Doe")
	u2.Update()

	u2 = users.FetchFieldsByPK[users.User](1)
	if u2.GetEmail() != "joe@doe.com" && u2.GetName() != "JOe Doe" {
		t.Error("Failed to FetchByPK with User using id 1 after update")
	}
	u2.SetEmail("john@doe.com")
	u2.SetName("John Doe")
	u2.Update()

	u3 := users.FetchByPK(1)
	if u2.GetEmail() != u3.GetEmail() && u2.GetName() != u3.GetName() {
		t.Error("FetchUserByPK and FetchByPK[User] returns different result")
	}

	u4 := users.FetchByPK(0)
	if u4 != nil {
		t.Error("FetchUserByPK must return nil for id 0")
	}

	objs := users.FetchFieldsByPKs[SimpleUser](1, 2)
	if len(objs) != 2 {
		t.Error("FetchByPKs[SimpleUser]([]int{1, 2}) failed")
	}
	if objs[0].GetEmail() != u.GetEmail() {
		t.Error("FetchByPK and FetchByPKs[SimpleUser] returns different result")
	}

	objs2 := users.FetchByPKs(3, 4)
	if len(objs2) != 2 {
		t.Error("FetchUserByPKs([]int{3, 4}) failed")
	}
}
