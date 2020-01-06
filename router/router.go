package router

import (
	"backend/controllers"
	"backend/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "successed",
		})
	})

	router.Use(middleware.Cors())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	loginApi := router.Group("/login")
	{
		loginApi.POST("", controllers.Login) // 用户登录
	}
	//router.Use(jwtAuthenticateMiddleware)

	menuApi := router.Group("/menus")
	{
		menuApi.GET("", controllers.GetMenuList)
	}

	userApi := router.Group("/users")
	{
		userApi.POST("", controllers.CreateUser)               //创建用户
		userApi.GET("/detail/:userId", controllers.UserDetail) //用户详情
		userApi.GET("", controllers.GetUserList)               //获取用户列表
		userApi.PUT("", controllers.UpdateUser)                //更新用户
		userApi.PUT("/state", controllers.UpdateUserState)     //更新用户状态
		userApi.PUT("/role", controllers.UpdateUserRoleId)     //更新用户状态
		userApi.DELETE("/:userId", controllers.DeleteUser)     //删除用户
	}

	roleApi := router.Group("/roles")
	{
		roleApi.POST("", controllers.CreateRole)
		roleApi.DELETE("/:roleId", controllers.DeleteRole)
		roleApi.POST("/rights", controllers.SetRoleRights)
		roleApi.PUT("/rights", controllers.DeleteRoleRight)
		roleApi.GET("/detail/:roleId", controllers.GetRoleDetail)
		roleApi.PUT("", controllers.UpdateRole)
		roleApi.GET("", controllers.GetRoleList)
	}

	rightApi := router.Group("/rights")
	{
		rightApi.GET("", controllers.GetRights)
		rightApi.GET("/tree", controllers.GetRightsTree)
		rightApi.DELETE("/rightId", controllers.DeleteRight)
	}

	return router
}
