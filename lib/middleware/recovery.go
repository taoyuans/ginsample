package middleware

//RecoveryMiddleware捕获所有panic，并且返回错误信息
// func RecoveryMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		defer func() {
// 			if err := recover(); err != nil {
// 				if lib.ConfBase.DebugMode != "debug" {
// 					ResponseError(c, 500, errors.New("内部错误"))
// 					return
// 				} else {
// 					ResponseError(c, 5001, errors.New(fmt.Sprint(err)))
// 					return
// 				}
// 			}
// 		}()
// 		c.Next()
// 	}
// }
