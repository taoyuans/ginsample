package models

import (
	"ginsample/lib/test"
	"testing"
)

func Test_GetUsers(t *testing.T) {
	users, err := User{}.GetUsers(ctx)
	test.Ok(t, err)
	test.Equals(t, users[0].Id, int64(1))
	test.Equals(t, users[0].Code, "xiaoming")
	test.Equals(t, users[0].Name, "小明")
	test.Equals(t, users[0].Enable, true)
	test.Equals(t, users[1].Id, int64(2))
	test.Equals(t, users[1].Code, "xiaozhang")
	test.Equals(t, users[1].Name, "小张")
	test.Equals(t, users[1].Enable, true)
}

func Test_GetById(t *testing.T) {
	user, err := User{}.GetById(ctx, int64(1))
	test.Ok(t, err)
	test.Equals(t, user.Id, int64(1))
	test.Equals(t, user.Code, "xiaoming")
	test.Equals(t, user.Name, "小明")
}
