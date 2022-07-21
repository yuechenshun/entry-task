package _init

import (
	"entry-task/controller"
	docs "entry-task/docs"
	"entry-task/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(cfg *Config) {
	r := gin.Default()

	// swagger
	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 查看主页信息和登陆不需要身份认证
	r.GET("/index", controller.Test{}.TestFunc)
	r.POST("/login", controller.Login)

	// 以下接口需要开启jwt身份认证
	api := r.Group("/api", middleware.JwtAuth())
	{

		api.GET("/user_info", controller.ViewUserInfo) // 查看用户信息

		// 商品相关
		api.GET("/product_list", controller.ViewProductList)     // 查看商品列表
		api.GET("/product_detail", controller.ViewProductDetail) // 查看商品详情
		api.POST("/product_create", controller.CreateProduct)    // 增加商品
		api.DELETE("/product_delete", controller.DeleteProduct)  // 删除商品
		api.PUT("/product_modify", controller.ModifyProduct)     // 编辑商品
		api.GET("/product_search", controller.SearchProduct)     // 搜索商品
		api.GET("/product_flow", controller.ProductFlow)         //查看商品推荐流
		api.PUT("/product_down", controller.ProductDown)         // 下架商品
		// 店铺相关
		api.POST("/shop_create", controller.CreateShop)      // 创建店铺
		api.GET("/shop_list", controller.ViewShopList)       // 查看店铺列表
		api.PUT("/shop_modify", controller.ModifyShop)       // 修改店铺
		api.DELETE("/shop_delete", controller.DeleteShop)    // 删除店铺
		api.GET("/shop_product", controller.ViewShopProduct) // 查看店铺商品
	}

	r.Run(cfg.AppHost + ":" + cfg.AppPort)
}
