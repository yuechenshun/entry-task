package controller

import (
	"entry-task/model"
	"entry-task/response"
	"entry-task/server"
	"entry-task/util"
	"github.com/gin-gonic/gin"
)

var shop server.WebShopServer

// @BasePath /api

// CreateShop
// @Summary 创建店铺
// @Schemes
// @Description 卖家输入店铺信息，创建店铺
// @Tags 需要登陆
// @Accept json
// @Produce json
// @Param object body model.ShopCreateParam true "创建商铺参数"
// @Success 200 {string} json "{"msg":"店铺创建成功"}"
// @Failure 400 {string} json "{"msg":"店铺创建失败"}"
// @Router /shop_create [post]
func CreateShop(c *gin.Context) {
	var param model.ShopCreateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Fail(c, "请求参数无效")
		return
	}
	// 获取当前登陆用户
	username := c.MustGet("username").(string)
	u := util.GetUserExample(username)
	if u.Class != 1 {
		response.Fail(c, "您不是卖家，无法创建店铺")
		return
	}
	count := shop.Create(param, u.Id)
	if count == 0 {
		response.Fail(c, "创建店铺失败")
		return
	}
	response.Success(c, "创建店铺成功", count)

}

// @BasePath /api

// ViewShopList
// @Summary 查看店铺列表
// @Schemes
// @Description 卖家查看自己的商铺列表
// @Tags 需要登陆
// @Accept json
// @Produce json
// @Success 200 {string} json "{"msg":"店铺列表查询成功"}"
// @Failure 400 {string} json "{"msg":"店铺列表查询失败"}"
// @Router /shop_list [get]
func ViewShopList(c *gin.Context) {
	// 获取当前登陆用户
	username := c.MustGet("username").(string)
	u := util.GetUserExample(username)
	if u.Class != 1 {
		response.Fail(c, "您不是卖家，无法查看店铺列表")
		return
	}
	shopList := shop.ShopList(u.Id)
	if len(shopList) == 0 {
		response.Fail(c, "当前店铺列表为空")
		return
	}
	response.Success(c, "店铺列表查询成功", shopList)
}

// @BasePath /api

// ModifyShop
// @Summary 修改店铺信息
// @Schemes
// @Description 用于卖家修改店铺信息
// @Tags 需要登陆
// @Accept json
// @Produce json
// @Param object body model.ShopModifyParam true "修改店铺参数"
// @Success 200 {string} json "{"msg":"店铺信息修改成功"}"
// @Failure 400 {string} json "{"msg":"店铺新秀修改失败"}"
// @Router /shop_modify [put]
func ModifyShop(c *gin.Context) {
	var param model.ShopModifyParam
	if err := c.ShouldBind(&param); err != nil {
		response.Fail(c, "请求参数无效")
		return
	}
	// 获取当前登陆用户
	username := c.MustGet("username").(string)
	u := util.GetUserExample(username)
	if u.Class != 1 {
		response.Fail(c, "您不是卖家，无法修改店铺信息")
		return
	}
	count := shop.Modify(param, u.Id)
	if count == 0 {
		response.Fail(c, "店铺信息修改失败")
		return
	}
	response.Success(c, "店铺信息修改成功", count)

}

// @BasePath /api

// DeleteShop
// @Summary 删除店铺
// @Schemes
// @Description 用于卖家删除店铺
// @Tags 需要登陆
// @Accept json
// @Produce json
// @Param object body model.ShopDeleteParam true "删除店铺参数"
// @Success 200 {string} json "{"msg":"店铺删除成功"}"
// @Failure 400 {string} json "{"msg":"店铺删除失败"}"
// @Router /shop_delete [delete]
func DeleteShop(c *gin.Context) {
	var param model.ShopDeleteParam
	if err := c.ShouldBind(&param); err != nil {
		response.Fail(c, "请求参数无效")
		return
	}
	// 获取当前登陆用户
	username := c.MustGet("username").(string)
	u := util.GetUserExample(username)
	if u.Class != 1 {
		response.Fail(c, "您不是卖家，无法删除店铺")
		return
	}
	count := shop.Delete(param)
	if count == 0 {
		response.Fail(c, "店铺删除失败")
		return
	}
	response.Success(c, "店铺删除成功", count)
}

// @BasePath /api

// ViewShopProduct
// @Summary 查看店铺下商品
// @Schemes
// @Description 如果是买家则展示所有商品，卖家则可以根据商品状态查询对应商品
// @Tags 需要登陆
// @Accept json
// @Produce json
// @Param object query model.ShopProductParam true "商品列表查询参数"
// @Success 200 {string} json "{"msg":"商品列表查询成功"}"
// @Failure 400 {string} json "{"msg":"商品列表查询失败"}"
// @Router /shop_product [get]
func ViewShopProduct(c *gin.Context) {
	var param model.ShopProductParam
	if err := c.ShouldBind(&param); err != nil {
		response.Fail(c, "请求参数无效")
		return
	}
	// 获取当前登陆用户
	username := c.MustGet("username").(string)
	u := util.GetUserExample(username)
	// 如果是卖家,则可以根据卖家输入商品状态查询商品
	if u.Class == 1 {
		productList := shop.SellerShopProduct(param)
		if len(productList) == 0 {
			response.Fail(c, "商品列表查询失败")
			return
		}
		response.Success(c, "商品列表查询成功", productList)
	} else {
		productList := shop.BuyerShopProduct(param)
		if len(productList) == 0 {
			response.Fail(c, "商品列表查询失败")
			return
		}
		response.Success(c, "商品列表查询成功", productList)
	}
}
