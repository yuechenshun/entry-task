package server

import (
	"entry-task/global"
	"entry-task/model"
	"time"
)

type WebShopServer struct {
}

// Create 创建店铺
func (con *WebShopServer) Create(param model.ShopCreateParam, ownerId uint64) int64 {
	shop := model.Shop{
		Name:       param.Name,
		Type:       param.Type,
		Detail:     param.Detail,
		OwnerId:    ownerId,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	rows := global.Db.Create(&shop).RowsAffected
	return rows
}

// ShopList 查看店铺列表
func (con *WebShopServer) ShopList(ownerId uint64) []model.ViewShopList {
	shopList := make([]model.ViewShopList, 0)
	global.Db.Table("shops").Where("owner_id = ?", ownerId).Find(&shopList)
	return shopList
}

// Modify 修改店铺信息
func (con *WebShopServer) Modify(param model.ShopModifyParam, ownerId uint64) int64 {
	shop := model.Shop{
		Id:         param.Id,
		Name:       param.Name,
		Type:       param.Type,
		Detail:     param.Detail,
		OwnerId:    ownerId,
		UpdateTime: time.Now(),
	}
	rows := global.Db.Model(&shop).Updates(shop).RowsAffected
	return rows
}

// Delete 删除店铺
func (con *WebShopServer) Delete(param model.ShopDeleteParam) int64 {
	rows := global.Db.Delete(&model.Shop{}, param.ShopId).RowsAffected
	return rows
}

// BuyerShopProduct 买家查询商品列表
func (con *WebShopServer) BuyerShopProduct(param model.ShopProductParam) []model.ShopProduct {
	productList := make([]model.ShopProduct, 0)
	// 返回所有已上架到商品
	global.Db.Table("products").Where("status = 1 and shop_id = ?", param.Id).Find(&productList)
	return productList
}

// SellerShopProduct 卖家查询商品列表
func (con *WebShopServer) SellerShopProduct(param model.ShopProductParam) []model.ShopProduct {
	productList := make([]model.ShopProduct, 0)
	status := param.Status
	// 卖家输入0则返回所有商品
	if status == 0 {
		global.Db.Table("products").Where("shop_id = ?", param.Id).Find(&productList)
		return productList
	}
	global.Db.Table("products").Where("status = ? and shop_id = ?", status, param.Id).Find(&productList)
	return productList
}
