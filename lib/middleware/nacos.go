package middleware

// func SetNacosNamingClient(nacosNamingClient naming_client.INamingClient) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		req := ctx.Request
// 		ctx.Request = req.WithContext(context.WithValue(req.Context(), "NacosNamingClient", &nacosNamingClient))
// 		ctx.Next()
// 	}
// }
