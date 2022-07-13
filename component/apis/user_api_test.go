package apis

import (
	"encoding/json"
	"fmt"
	"ginsample/lib/test"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_AppApiController_GetApp(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users", nil)
	r.ServeHTTP(w, req)

	var v struct {
		Result []struct {
			Id     int64  `json:"id"`
			Code   string `json:"code"`
			Name   string `json:"name"`
			Enable bool   `json:"enable"`
		} `json:"result"`
		Success bool                   `json:"success"`
		Errors  map[string]interface{} `json:"error"`
	}

	test.Equals(t, http.StatusOK, w.Code)
	test.Ok(t, json.Unmarshal(w.Body.Bytes(), &v))

	fmt.Printf("%+v\n", v)

	test.Equals(t, v.Success, true)
	test.Equals(t, v.Result[0].Id, int64(1))
	test.Equals(t, v.Result[0].Code, "xiaoming")
	test.Equals(t, v.Result[0].Name, "小明")
	test.Equals(t, v.Result[0].Enable, true)
	test.Equals(t, v.Result[1].Id, int64(2))
	test.Equals(t, v.Result[1].Code, "xiaozhang")
	test.Equals(t, v.Result[1].Name, "小张")
	test.Equals(t, v.Result[1].Enable, true)
	fmt.Println("----------haha2----------")

}
