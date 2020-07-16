package server

import (
	"cmall/api"
	"cmall/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	// 路由
	v1 := r.Group("/api/v1")
	{
		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 管理员登录
		v1.POST("admin/login", api.AdminLogin)

		//商品操作
		v1.POST("products", api.CreateProduct)
		v1.GET("products", api.ListProducts)
		v1.GET("products/:id", api.ShowProduct)
		v1.GET("categories/:category_id", api.ShowCategory)
		v1.DELETE("products/:id", api.DeleteProduct)
		v1.PUT("products", api.UpdateProduct)
		//轮播图操作
		v1.POST("carousels", api.CreateCarousel)
		v1.GET("carousels", api.ListCarousels)
		//商品图片操作
		v1.POST("pictures", api.CreatePicture)
		v1.GET("pictures/:id", api.ShowPictures)
		//分类操作
		v1.POST("categories", api.CreateCategory)
		v1.GET("categories", api.ListCategories)
		//搜索操作
		v1.POST("searches", api.SearchProducts)
		//排行榜
		v1.GET("rankings", api.ShowRanking)
		//README操作
		v1.GET("abouts", api.ReadMe)
		// 需要登录保护的
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			//验证token
			authed.GET("ping", api.CheckToken)
			//用户操作
			authed.PUT("user", api.UserUpdate)
			authed.DELETE("user/logout", api.UserLogout)
			// 上传操作
			authed.POST("avatar", api.UploadToken)
			//收藏夹操作
			authed.GET("favorites/:id", api.ShowFavorites)
			authed.POST("favorites", api.CreateFavorite)
			authed.DELETE("favorites", api.DeleteFavorite)
			//订单操作
			authed.POST("orders", api.CreateOrder)
			authed.GET("orders/:id", api.ShowOrders)
			//购物车
			authed.POST("carts", api.CreateCart)
			authed.GET("carts/:id", api.ShowCarts)
			authed.PUT("carts", api.UpdateCart)
			authed.DELETE("carts", api.DeleteCart)
		}

	}
	v2 := r.Group("/api/v2")
	{
		// 管理员注册
		v2.POST("admin/register", api.AdminRegister)
		// 管理员登录
		v2.POST("admin/login", api.AdminLogin)
		v2.Use(middleware.JWT())
		v2.GET("products", api.ListProducts)
	}
	return r
}
