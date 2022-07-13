package models

import (
	"ginsample/lib/test"
	"testing"
)

func Test_GetApps(t *testing.T) {
	user, err := User{}.GetById(ctx, int64(1))
	test.Ok(t, err)
	test.Equals(t, user.Id, int64(1))
	test.Equals(t, user.Code, "xiaoming")
	test.Equals(t, user.Name, "小明")
}
