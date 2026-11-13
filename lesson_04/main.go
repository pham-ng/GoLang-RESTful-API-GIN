package main

import (
	"github.com/gin-gonic/gin"

	v1handler "myapp/lesson_04/internal/api/v1/handler"
	v2handler "myapp/lesson_04/internal/api/v2/handler"
	"myapp/lesson_04/utils"
)

func main() {
	r := gin.Default()
	if err := utils.RegisterValidators(); err != nil {
		panic(err)
	}

	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			userHandlerV1 := v1handler.NewUserHandler()
			user.GET("", userHandlerV1.GetUsersV1)
			user.GET("/:id", userHandlerV1.GetUserByIDV1)
			user.GET("/admin/:uuid", userHandlerV1.GetUserByUUIDV1)
			user.POST("", userHandlerV1.CreateUserV1)
			user.PUT("/:id", userHandlerV1.UpdateUserV1)
			user.DELETE("/:id", userHandlerV1.DeleteUserV1)
		}

		product := v1.Group("/products")
		{
			productHandlerV1 := v1handler.NewProductHandler()
			product.GET("", productHandlerV1.GetProductsV1)
			product.GET("/:slug", productHandlerV1.GetProductBySlugV1)
			product.POST("", productHandlerV1.PostProductV1)
			product.PUT("/:id", productHandlerV1.PutProductV1)
			product.DELETE("/:id", productHandlerV1.DeleteProductV1)
		}

		category := v1.Group("/categories")
		{
			categoryHandlerV1 := v1handler.NewCategoryHandler()
			category.GET("", categoryHandlerV1.GetCategoriesV1)
			category.GET("/:slug", categoryHandlerV1.GetCategoryBySlugV1)
			category.POST("", categoryHandlerV1.PostCategoryV1)
			category.PUT("/:id", categoryHandlerV1.PutCategoryV1)
			category.DELETE("/:id", categoryHandlerV1.DeleteCategoryV1)
		}
		news := v1.Group("/news")
		{
			newsHandlerV1 := v1handler.NewNewsHandler()
			news.GET("", newsHandlerV1.GetNewsV1)
			news.GET("/:slug", newsHandlerV1.GetNewsV1)
		}
	}

	v2 := r.Group("/api/v2")
	{
		user := v2.Group("/users")
		{
			userHandlerV2 := v2handler.NewUserHandler()
			user.GET("", userHandlerV2.GetUsersV2)
			user.GET("/:id", userHandlerV2.GetUserByIDV2)
			user.POST("", userHandlerV2.CreateUserV2)
			user.PUT("/:id", userHandlerV2.UpdateUserV2)
			user.DELETE("/:id", userHandlerV2.DeleteUserV2)
		}

		product := v2.Group("/products")
		{
			productHandlerV2 := v2handler.NewProductHandler()
			product.GET("", productHandlerV2.GetProductsV2)
			product.GET("/:id", productHandlerV2.GetProductByIdV2)
			product.POST("", productHandlerV2.PostProductV2)
			product.PUT("/:id", productHandlerV2.PutProductV2)
			product.DELETE("/:id", productHandlerV2.DeleteProductV2)
		}
	}

	r.Run(":8082")
}
