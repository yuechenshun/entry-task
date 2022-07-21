package server

import (
	"entry-task/global"
	"entry-task/model"
	"time"
)

type WebProductServer struct {
}

// GetProductList 获取商品列表
func (con *WebProductServer) GetProductList(param model.ProductListParam) []model.ProductList {
	productList := make([]model.ProductList, 0)
	// 如果用户输入商品类别编号则返回对应类别的商品
	if param.CategoryId != 0 {
		categoryId := param.CategoryId
		global.Db.Table("products").Where("category_id = ? and status = 1", categoryId).Find(&productList)
		return productList
	}
	// 否则返回所有商品
	global.Db.Table("products").Where("status = 1").Find(&productList)
	return productList
}

// GetProductDetail 获取商品信息
func (con *WebProductServer) GetProductDetail(param model.ProductDetailParam) model.ProductDetail {
	var product model.ProductDetail
	// 如果输入的商品id不为0，则返回商品信息
	if param.ProductId != 0 {
		productId := param.ProductId
		global.Db.Table("products").Where("id = ?", productId).Find(&product)
		return product
	}
	// 否则返回空
	return product
}

// Create 创建商品
func (con *WebProductServer) Create(param model.ProductCreateParam) int64 {
	product := model.Product{
		Name:       param.Name,
		Price:      param.Price,
		Store:      param.Store,
		Image:      param.Image,
		CategoryId: param.CategoryId,
		Type:       param.Type,
		Detail:     param.Detail,
		Status:     0,
		ShopId:     param.ShopId,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	// 返回插入记录的条数
	rows := global.Db.Create(&product).RowsAffected
	return rows
}

// Delete 删除商品
func (con *WebProductServer) Delete(param model.ProductDeleteParam) int64 {
	rows := global.Db.Delete(&model.Product{}, param.ProductId).RowsAffected
	return rows
}

// Modify 修改商品
func (con *WebProductServer) Modify(param model.ProductModifyParam) int64 {
	product := model.Product{
		Id:         param.Id,
		Name:       param.Name,
		Price:      param.Price,
		Store:      param.Store,
		Image:      param.Image,
		CategoryId: param.CategoryId,
		Type:       param.Type,
		Detail:     param.Detail,
		Status:     3,
		ShopId:     param.ShopId,
		UpdateTime: time.Now(),
	}
	rows := global.Db.Model(&product).Updates(product).RowsAffected
	return rows
}

// Search 搜索商品
func (con *WebProductServer) Search(param model.ProductSearchParam) []model.ProductSearch {
	productList := make([]model.ProductSearch, 0)
	keyWord := param.KeyWord
	// 如果没有输入关键字则返回所有商品
	if len(keyWord) == 0 {
		global.Db.Table("products").Where("status = 1").Find(&productList)
		return productList
	}
	global.Db.Table("products").Where("name like ? and status = 1", "%"+keyWord+"%").Find(&productList)
	return productList
}

// Flow 查看商品推荐流
func (con *WebProductServer) Flow(param model.ProductFlowParam) []model.ProductFlow {
	productList := make([]model.ProductFlow, 0)
	flowList := param.ProductList
	for _, v := range flowList {
		var product model.ProductFlow
		global.Db.Table("products").Where("id = ? and status = 1", v).Find(&product)
		productList = append(productList, product)
	}
	return productList
}

// Down 下架商品
func (con *WebProductServer) Down(param model.ProductDownParam) int64 {
	product := model.Product{
		Id:     param.ProductId,
		Status: 2,
	}
	rows := global.Db.Model(&product).Updates(product).RowsAffected
	return rows
}
