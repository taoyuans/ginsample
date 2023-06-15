package models

import (
	"fmt"
	"ginsample/lib/auth"
	"ginsample/lib/test"
	"testing"
)

func TestMain(m *testing.M) {
	// 在这里注册要在所有测试用例执行完毕后调用的函数
	defer RemoveSqlite()

	// 执行测试用例
	m.Run()
}

func Test_User_GetByUserId(t *testing.T) {
	user, err := User{}.GetByUserId(ctx, 10000000000000001)

	test.Ok(t, err)
	test.Equals(t, user.Id, int64(1))
	test.Equals(t, user.UserId, int64(10000000000000001))
	test.Equals(t, user.UserName, "zhangsan")
	test.Equals(t, user.PhoneNo, "13456789091")
	test.Equals(t, user.Email, "13456789091@qq.com")
	test.Equals(t, user.Deleted, false)
}

func Test_Test(t *testing.T) {
	aa, err := Login{}.LoginByUserName(ctx, "zhangsan", "1234")
	fmt.Println(aa)
	fmt.Println(err)
	auth.CheckToken(aa["token"].(string))
}
