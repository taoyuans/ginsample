package middleware

import (
	"context"
	"ginsample/lib/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// func SetDBMiddleware(gormDB *gorm.DB, next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		timeoutContext, _ := context.WithTimeout(context.Background(), time.Second)
// 		ctx := context.WithValue(r.Context(), "DB", gormDB.WithContext(timeoutContext))
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

func SetDBMiddleware(gormDB *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := ctx.Request
		ctx.Request = req.WithContext(context.WithValue(req.Context(), "DB", gormDB))
		// if err := setTokenInfo(ctx, req); err != nil {
		// 	fmt.Println(err)
		// }
		// ctx.Set("DB", gormDB)
		ctx.Next()
	}
}

func setTokenInfo(ctx *gin.Context, req *http.Request) error {
	userClaims, err := auth.GetTokenInfo(ctx)
	if err != nil {
		return err
	}
	ctx.Request = req.WithContext(context.WithValue(req.Context(), "Colleague", userClaims))
	return nil
}

// func SetDbContext(db *xorm.Engine) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			req := c.Request()
// 			switch req.Method {
// 			case "POST", "PUT", "DELETE":
// 				session := db.NewSession()
// 				defer session.Close()

// 				c.SetRequest(req.WithContext(context.WithValue(req.Context(), echomiddleware.ContextDBName, session)))
// 				if err := session.Begin(); err != nil {
// 					log.Println(err)
// 				}
// 				if err := next(c); err != nil {
// 					session.Rollback()
// 					return errs.Trace(err)
// 				}
// 				if c.Response().Status >= 500 {
// 					session.Rollback()
// 					return nil
// 				}
// 				if err := session.Commit(); err != nil {
// 					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 				}
// 			default:
// 				c.SetRequest(req.WithContext(context.WithValue(req.Context(), echomiddleware.ContextDBName, db)))
// 				return next(c)
// 			}

// 			return nil
// 		}
// 	}
// }
