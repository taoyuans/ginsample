package middleware

// func SetRedisContext(valueKey string, redisCache cache.Cache) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			req := c.Request()
// 			c.SetRequest(req.WithContext(context.WithValue(req.Context(), valueKey, redisCache)))
// 			return next(c)
// 		}
// 	}
// }
