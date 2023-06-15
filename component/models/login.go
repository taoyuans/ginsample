package models

import (
	"context"
	"ginsample/lib/auth"
	"ginsample/lib/errs"
)

type Login struct {
}

func (Login) LoginByUserName(ctx context.Context, userName string, password string) (map[string]interface{}, error) {
	user, err := User{}.CheckPassword(ctx, userName, password)
	if err != nil {
		return nil, errs.Trace(err)
	}

	if user == nil || user.Id == 0 {
		return nil, errs.New("userName or password is invalid!")
	}
	token, err := auth.NewToken(map[string]interface{}{
		"userId":  user.UserId,
		"phoneNo": user.PhoneNo,
		"email":   user.Email,
	})
	if err != nil {
		return nil, errs.Trace(err)
	}

	return map[string]interface{}{"token": token}, nil
}
