package server

import (
	"cmall/api"
	"cmall/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)
		//更新用户信息
		v1.PUT("user", api.UserUpdate)

		// 管理员登录
		v1.POST("admin/login", api.AdminLogin)

		// 需要登录保护的
		authed := v1.Group("/")
		authed.Use(middleware.AuthRequired())
		{
			// User Routing
			authed.GET("user/me", api.UserMe)
			authed.DELETE("user/logout", api.UserLogout)
		}
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
		//收藏夹操作
		v1.POST("favorites", api.CreateFavorite)
		v1.GET("favorites/:id", api.ShowFavorites)
		v1.DELETE("favorites", api.DeleteFavorite)
		//分类操作
		v1.POST("categories", api.CreateCategory)
		v1.GET("categories", api.ListCategories)
		//搜索操作
		v1.POST("searches", api.SearchProducts)
		//订单操作
		v1.POST("orders", api.CreateOrder)
		v1.GET("orders/:id", api.ShowOrders)
		//购物车
		v1.POST("carts", api.CreateCart)
		v1.GET("carts/:id", api.ShowCarts)
		v1.PUT("carts", api.UpdateCart)
		v1.DELETE("carts", api.DeleteCart)
		//热门
		v1.GET("EHots", api.ListProducts)
		v1.GET("PHots", api.ListProducts)
		//排行榜
		v1.GET("rankings/", api.ShowRanking)
		//README操作
		v1.GET("abouts", api.ReadMe)
		// 上传操作
		v1.POST("avatar", api.UploadToken)
	}

	return r
}
