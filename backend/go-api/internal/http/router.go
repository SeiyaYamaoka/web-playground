package http

import "github.com/gin-gonic/gin"

// SetupRouter はAPIのエンドポイントを定義
func SetupRouter(router *gin.Engine, userHandler *UserHandler) {
	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			// CRUD操作のエンドポイント
			users.POST("", userHandler.CreateUser)       // ユーザー作成
			users.GET("/:id", userHandler.GetUser)       // ユーザー取得 (ID)
			users.GET("", userHandler.GetAllUsers)       // 全ユーザー取得
			users.PUT("/:id", userHandler.UpdateUser)    // ユーザー更新
			users.DELETE("/:id", userHandler.DeleteUser) // ユーザー削除
		}
	}
}
