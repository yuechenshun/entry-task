package model

import "time"

// Shop 店铺数据库映射模型
type Shop struct {
	Id         uint64    `gorm:"column:id;primaryKey" json:"id"`        // 店铺id
	Name       string    `gorm:"column:name" json:"name"`               // 店铺名称
	Type       string    `gorm:"column:type" json:"type"`               // 店铺类别
	Detail     string    `gorm:"column:detail" json:"detail"`           // 店铺详情
	OwnerId    uint64    `gorm:"column:owner_id" json:"owner_id"`       // 店铺所有者id
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` // 店铺创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` // 店铺更新时间
}

// ShopCreateParam 创建商铺参数模型
type ShopCreateParam struct {
	Name   string `json:"name" form:"name" binding:"required"`
	Type   string `json:"type" form:"type" binding:"required"`
	Detail string `json:"detail" form:"detail"`
}

// ViewShopList 查看店铺列表模型
type ViewShopList struct {
	Id     uint64 `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Detail string `json:"detail"`
}

// ShopModifyParam 修改店铺参数模型
type ShopModifyParam struct {
	Id     uint64 `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	Type   string `json:"type" form:"type"`
	Detail string `json:"detail" form:"detail"`
}

// ShopDeleteParam 删除店铺参数模型
type ShopDeleteParam struct {
	ShopId uint64 `json:"shop_id" form:"shop_id"`
}

// ShopProductParam 店铺商品参数模型
type ShopProductParam struct {
	Id     uint64 `json:"id" form:"id"`         // 需要查看的店铺id
	Status uint8  `json:"status" form:"status"` // 需要查看的商品状态
}

// ShopProduct 店铺商品信息模型
type ShopProduct struct {
	Id         uint64  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Store      uint64  `json:"store"`
	Image      string  `json:"image"`
	CategoryId uint64  `json:"category_id"`
	Type       string  `json:"type"`
	Detail     string  `json:"detail"`
}
