package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"ginsample/lib/test"
)

func TestMain(m *testing.M) {
	// 在这里注册要在所有测试用例执行完毕后调用的函数
	defer RemoveSqlite()

	// 执行测试用例
	m.Run()
}

func Test_UserApi_GetUsers(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users?userId=10000000000000001", nil)
	r.ServeHTTP(w, req)

	var v struct {
		Rows struct {
			Id       int64  `json:"id"`
			UserId   int64  `json:"userId"`
			PhoneNo  string `json:"phoneNo"`
			Email    string `json:"email"`
			UserName string `json:"userName"`
			Password string `json:"-"`
			Deleted  bool   `json:"deleted"`
		} `json:"rows"`
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}

	test.Equals(t, http.StatusOK, w.Code)
	fmt.Printf("%+v\n", w.Body)
	test.Ok(t, json.Unmarshal(w.Body.Bytes(), &v))

	test.Equals(t, v.Code, 200)

	test.Equals(t, v.Rows.Id, int64(1))
	test.Equals(t, v.Rows.UserId, int64(10000000000000001))
	test.Equals(t, v.Rows.UserName, "zhangsan")
	test.Equals(t, v.Rows.PhoneNo, "13456789091")
	test.Equals(t, v.Rows.Email, "13456789091@qq.com")
	test.Equals(t, v.Rows.Deleted, false)
}

func Test_UserApi_LoginByUserName(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/login-user-name", strings.NewReader(`{ "userName": "zhangsan", "password": "1234" }`))
	r.ServeHTTP(w, req)

	var v struct {
		Rows struct {
			Token string `json:"token"`
		} `json:"rows"`
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}

	fmt.Printf("%+v\n", w)

	test.Equals(t, http.StatusOK, w.Code)
	test.Ok(t, json.Unmarshal(w.Body.Bytes(), &v))

	test.Equals(t, len(v.Rows.Token) > 0, true)
}
