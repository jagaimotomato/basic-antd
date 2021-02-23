package v1

import (
	"basic-antd/app/api"
	"basic-antd/app/api/public"
	"basic-antd/app/middleware"
	"basic-antd/app/middleware/handler"
	"basic-antd/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(r *gin.Engine, auth *jwt.GinJWTMiddleware) *gin.RouterGroup {
	var (
		g         = r.Group("")
		authGroup = r.Group("/auth")
		//p         = r.Group("") // Public
	)
	sysPingRouter(g)

	{
		/* 公共路由 无需权限*/
		sysStaticFileRouter(g)
	}

	{
		/* 登录,登出相关路由 */
		registerAuthRouter(authGroup, auth)
		registerLogoutRouter(authGroup, auth)
	}
	{
		/* 系统基础路由 */
		registerUserInfoRouter(g, auth)
		registerSystemBaseRouter(g, auth)
		registerSystemPageRouter(g, auth)
		registerSystemListRouter(g, auth)
		registerSystemTreeRouter(g, auth)
		registerUploadRouter(g, auth)
		sysAuthPingRouter(g, auth)
		registerSystemOtherRouter(g, auth)
	}

	//v1 := r.Group("/api/v1")
	return g
}

func sysPingRouter(r *gin.RouterGroup) {
	r.GET("/", api.Pong)
}

func sysAuthPingRouter(r *gin.RouterGroup, auth *jwt.GinJWTMiddleware) {
	r.Use(auth.MiddlewareFunc()).GET("/auth/ping", api.AuthPong)
}

func registerAuthRouter(v1 *gin.RouterGroup, auth *jwt.GinJWTMiddleware) {
	v1.POST("/login", auth.LoginHandler)
}

func registerLogoutRouter(v1 *gin.RouterGroup, auth *jwt.GinJWTMiddleware) {
	v1.Use(auth.MiddlewareFunc()).POST("/logout", handler.Logout)
}

func registerUserInfoRouter(v1 *gin.RouterGroup, auth *jwt.GinJWTMiddleware) {
	baseRouter := v1.Group("info").Use(auth.MiddlewareFunc()).Use(middleware.AuthCheck())
	{
		baseRouter.GET("/user", api.GetUserInfo)
		baseRouter.GET("/nav", api.GetUserNav)
		// baseRouter.GET("/ossConfig", public.GetOssConfig)
		// baseRouter.GET("/qiniu", public.GetQiNiuToken)
	}
}

func registerSystemPageRouter(g *gin.RouterGroup, auth *jwt.GinJWTMiddleware) {
	pageGroup := g.Group("/page").Use(auth.MiddlewareFunc()).Use(middleware.AuthCheck())
	{
		pageGroup.GET("/user", api.GetUserPage)
		/* permission page */
		pageGroup.GET("/permission", api.GetPermissionPage)
		pageGroup.GET("/role", api.GetRolePage)
		pageGroup.GET("/department", api.GetDepartmentPage)
		pageGroup.GET("/loginLog", api.GetLoginLogPage)
		pageGroup.GET("/operationLog", api.GetOperationLogPage)
		pageGroup.GET("/dictType", api.GetDictTypePage)
		pageGroup.GET("/dictData", api.GetDictDataPage)
	}
}

func registerSystemListRouter(g *gin.RouterGroup, auth *jwt.GinJWTMiddleware) {
	listGroup := g.Group("/list").Use(auth.MiddlewareFunc()).Use(middleware.AuthCheck())
	{
		/* permissionList */
		listGroup.GET("/permissionList", api.GetPermissionList)
		listGroup.GET("/roleList", api.GetRoleList)
		listGroup.GET("/dictDataList", api.GetDictDataList)
		listGroup.GET("/dictTypeList", api.GetDictTypeList)
	}
}

func registerSystemTreeRouter(g *gin.RouterGroup, auth *jwt.GinJWTMiddleware) {
	treeGroup := g.Group("/tree").Use(auth.MiddlewareFunc()).Use(middleware.AuthCheck())
	{
		treeGroup.GET("/permission", api.GetPermissionTree)
		treeGroup.GET("/permission/:roleId", api.GetRolePermissionTree)

		treeGroup.GET("/department", api.GetDepartmentTree)
	}
}

func registerSystemOtherRouter(g *gin.RouterGroup, auth *jwt.GinJWTMiddleware) {
	otherGroup := g.Group("/other").Use(auth.MiddlewareFunc()).Use(middleware.AuthCheck())
	{
		otherGroup.GET("/dictData/:type", api.GetDictDataList)
	}
}

func registerSystemBaseRouter(g *gin.RouterGroup, auth *jwt.GinJWTMiddleware) {
	userGroup := g.Group("/user").Use(auth.MiddlewareFunc()).Use(middleware.AuthCheck())
	{
		userGroup.GET("/:userId", api.GetUser)
		userGroup.POST("", api.CreateUser)
		userGroup.PUT("", api.UpdateUser)
		userGroup.PUT("/status", api.UpdateUserStatus)
		userGroup.DELETE("/:userId", api.DeleteUser)
	}
	/* permission */
	permissionGroup := g.Group("/permission").Use(auth.MiddlewareFunc()).Use(middleware.AuthCheck())
	{
		permissionGroup.GET("/:permissionId", api.GetPermission)
		permissionGroup.POST("", api.CreatePermission)
		permissionGroup.PUT("", api.UpdatePermission)
		permissionGroup.DELETE("/:permissionId", api.DeletePermission)
	}
	/* role */
	roleGroup := g.Group("/role").Use(auth.MiddlewareFunc()).Use(middleware.AuthCheck())
	{
		roleGroup.GET("/:roleId", api.GetRole)
		roleGroup.POST("", api.CreateRole)
		roleGroup.PUT("", api.UpdateRole)
		roleGroup.PUT("/status", api.UpdateRoleStatus)
		roleGroup.DELETE("/:roleId", api.DeleteRole)
	}
	/* department */
	departmentGroup := g.Group("/department").Use(auth.MiddlewareFunc()).Use(middleware.AuthCheck())
	{
		departmentGroup.GET("/:department_id", api.GetDepartment)
		departmentGroup.POST("", api.CreateDepartment)
		departmentGroup.PUT("", api.UpdateDepartment)
		departmentGroup.DELETE("/:department_id", api.DeleteDepartment)
	}
	/* dict */
	dictGroup := g.Group("/dict").Use(auth.MiddlewareFunc()).Use(middleware.AuthCheck())
	{
		dictGroup.GET("/data/:dictCode", api.GetDictData)
		dictGroup.POST("/data", api.CreateDictData)
		dictGroup.PUT("/data", api.UpdateDictData)
		dictGroup.DELETE("/data/:dictCode", api.DeleteDictData)

		dictGroup.GET("/type/:dictId", api.GetDictType)
		dictGroup.POST("/type", api.CreateDictType)
		dictGroup.PUT("/type", api.UpdateDictType)
		dictGroup.DELETE("/type/:dictId", api.DeleteDictType)
	}
}

func registerUploadRouter(v1 *gin.RouterGroup, auth *jwt.GinJWTMiddleware) {
	uploadGroup := v1.Group("/upload").Use(auth.MiddlewareFunc())
	{
		uploadGroup.POST("", public.UploadFile)
	}
}

func sysStaticFileRouter(g *gin.RouterGroup) {
	g.Static("/static", "./static")
}
