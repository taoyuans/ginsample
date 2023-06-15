package auth

// CorsConfig generates a config to use in gin cors middleware based on server configuration
// func CorsConfig(conf *config.Configuration) cors.Config {
// 	corsConf := cors.Config{
// 		MaxAge:                 12 * time.Hour,
// 		AllowWildcard:          conf.Server.Cors.AllowWildcard,
// 		AllowBrowserExtensions: conf.Server.Cors.AllowBrowserExtensions,
// 		AllowWebSockets:        conf.Server.Cors.AllowWebSockets,
// 	}

// 	if mode.IsDev() {
// 		corsConf.AllowAllOrigins = true
// 		corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS", "PUT"}
// 		corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
// 			"Connection", "Accept-Encoding", "Accept-Language", "Host"}
// 	} else {
// 		corsConf.AllowOrigins = conf.Server.Cors.AllowOrigins
// 		corsConf.AllowMethods = conf.Server.Cors.AllowMethods
// 		corsConf.AllowHeaders = conf.Server.Cors.AllowHeaders
// 	}

// 	return corsConf
// }
