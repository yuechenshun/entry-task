package controller

import (
	"entry-task/model"
	"entry-task/response"
	"entry-task/server"
	"entry-task/util"
	"github.com/gin-gonic/gin"
)

var product server.WebProductServer

// @BasePath /api

// ViewProductList
// @Summary 查看商品列表
// @Schemes
// @Description 用户可以根据输入商品category_id查看相应商品列表
// @Tags 需要登陆
// @Accept json
// @Produce json
// @Param object query model.ProductListParam true "查看商品列表参数"
// @Success 200 {string} json "{"msg":"商品列表查询成功"}"
// @Failure 400 {string} json "{"msg":"商品列表查询失败"}"
// @Router /product_list [get]
func ViewProductList(c *gin.Context) {
	var param model.ProductListParam
	if err := c.ShouldBind(&param); err != nil {
		response.Fail(c, "请求参数无效")
		return
	}
	productList := product.GetProductList(param)
	response.Success(c, "商品列表查询成功", productList)
}

// @BasePath /api

// ViewProductDetail
// @Summary 查看商品详情
// @Schemes
// @Description 用户可以根据输入商品id查看相应商品详情
// @Tags 需要登陆
// @Accept json
// @Produce json
// @Param object query model.ProductDetailParam true "查看商品详情参数"
// @Success 200 {string} json "{"msg":"商品详情查询成功"}"
// @Failure 400 {string} json "{"msg":"商品详情查询失败"}"
// @Router /product_detail [get]
func ViewProductDetail(c *gin.Context) {
	var param model.ProductDetailParam
	if err := c.ShouldBind(&param); err != nil {
		response.Fail(c, "请求参数无效")
		return
	}
	productDetail := product.GetProductDetail(param)
	if productDetail.Id != 0 {
		response.Success(c, "商品详情查询成功", productDetail)
		return
	}
	response.Fail(c, "查询失败，请输入需要查询的商品id")
}

// @BasePath /api

// CreateProduct
// @Summary 创建商品
// @Schemes
// @Description 用于用户创建商品
// @Tags 需要登陆
// @Accept json
// @Produce json
// @Param object body model.ProductCreateParam true "创建商品需要的参数"
// @Success 200 {string} json "{"msg":"商品创建成功"}"
// @Failure 400 {string} json "{"msg":"商品创建失败"}"
// @Router /product_create [post]
func CreateProduct(c *gin.Context) {
	var param model.ProductCreateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Fail(c, "请求参数无效")
		return
	}
	// 判断当前用户类型
	username := c.MustGet("username").(string)
	u := util.GetUserExample(username)
	if u.Class != 1 {
		response.Fail(c, "无法创建商品，您必须是卖家")
		return
	}
	count := product.Create(param)
	if count == 0 {
		response.Fail(c, "商品创建失败")
		return
	}
	response.Success(c, "商品创建成功", count)
}

// @BasePath /api

// DeleteProduct
// @Summary 删除商品
// @Schemes
// @Description 用于用户删除商品
// @Tags 需要登陆
// @Accept json
// @Produce json
// @Param object body model.ProductDeleteParam true "需要删除商品的id"
// @Success 200 {string} json "{"msg":"商品删除成功"}"
// @Failure 400 {string} json "{"msg":"商品删除失败"}"
// @Router /product_delete [delete]
func DeleteProduct(c *gin.Context) {
	var param model.ProductDeleteParam
	if err := c.ShouldBind(&param); err != nil {
		response.Fail(c, "请求参数无效")
		return
	}
	username := c.MustGet("username").(string)
	u := util.GetUserExample(username)
	if u.Class != 1 {
		response.Fail(c, "无法删除商品，您必须是卖家")
		return
	}
	count := product.Delete(param)
	if count == 0 {
		response.Fail(c, "商品删除失败")
		return
	}
	response.Success(c, "商品删除成功", count)
}

// @BasePath /api

// ModifyProduct
// @Summary 修改商品
// @Schemes
// @Description 用于用户修改商品
// @Tags 需要登陆
// @Accept json
// @Produce json
// @Param object body model.ProductModifyParam true "需要修改商品的参数"
// @Success 200 {string} json "{"msg":"商品修改成功"}"
// @Failure 400 {string} json "{"msg":"商品修改失败"}"
// @Router /product_modify [put]
func ModifyProduct(c *gin.Context) {
	var param model.ProductModifyParam
	if err := c.ShouldBind(&param); err != nil {
		response.Fail(c, "请求参数无效")
		return
	}
	username := c.MustGet("username").(string)
	u := util.GetUserExample(username)
	if u.Class != 1 {
		response.Fail(c, "无法修改商品，您必须是卖家")
		return
	}
	count := product.Modify(param)
	if count == 0 {
		response.Fail(c, "商品修改失败")
		return
	}
	response.Success(c, "商品修改成功", count)
}

// @BasePath /api

// SearchProduct
// @Summary 搜索商品
// @Schemes
// @Description 根据用户输入关键字搜索商品
// @Tags 需要登陆
// @Accept json
// @Produce json
// @Param object query model.ProductFlowParam true "商品推荐流关键字"
// @Success 200 {string} json "{"msg":"商品流推荐成功"}"
// @Failure 400 {string} json "{"msg":"商品流推荐失败"}"
// @Router /product_search [get]
func SearchProduct(c *gin.Context) {
	var param model.ProductSearchParam
	if err := c.ShouldBind(&param); err != nil {
		response.Fail(c, "请求参数无效")
		return
	}
	productList := product.Search(param)
	if len(productList) == 0 {
		response.Fail(c, "没有查询到相关商品")
		return
	}
	response.Success(c, "商品查询成功", productList)
}

// @BasePath /api

// ProductFlow
// @Summary 查看商品推荐流
// @Schemes
// @Description 接受商品id列表，返回对应商品信息
// @Tags 需要登陆
// @Accept json
// @Produce json
// @Param object query model.ProductDownParam true "商品搜索关键字"
// @Success 200 {string} json "{"msg":"商品下架成功"}"
// @Failure 400 {string} json "{"msg":"商品下架失败"}"
// @Router /product_down [put]
func ProductFlow(c *gin.Context) {
	var param model.ProductFlowParam
	if err := c.ShouldBind(&param); err != nil {
		response.Fail(c, "请求参数无效")
		return
	}
	productList := product.Flow(param)
	if len(productList) == 0 {
		response.Fail(c, "商品流查询失败")
		return
	}
	response.Success(c, "商品流查询成功", productList)
}

// @BasePath /api

// ProductDown
// @Summary 下架商品
// @Schemes
// @Description 接收商品id，下架对应商品
// @Tags 需要登陆
// @Accept json
// @Produce json
// @Param object body model.ProductSearchParam true "商品搜索关键字"
// @Success 200 {string} json "{"msg":"商品流查看成功"}"
// @Failure 400 {string} json "{"msg":"商品流查看失败"}"
// @Router /product_flow [get]
func ProductDown(c *gin.Context) {
	var param model.ProductDownParam
	if err := c.ShouldBind(&param); err != nil {
		response.Fail(c, "请求参数无效")
		return
	}
	username := c.MustGet("username").(string)
	u := util.GetUserExample(username)
	if u.Class != 1 {
		response.Fail(c, "无法修改商品，您必须是卖家")
		return
	}
	count := product.Down(param)
	if count == 0 {
		response.Fail(c, "商品下架失败")
		return
	}
	response.Success(c, "商品下架成功", count)
}
