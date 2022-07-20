package model

import (
	"time"
)

// Product 商品数据库映射模型
type Product struct {
	Id         uint64    `gorm:"column:id;primaryKey" json:"id"`        // 商品id
	Name       string    `gorm:"column:name" json:"name"`               // 商品名称
	Price      float64   `gorm:"column:price" json:"price"`             // 商品价格
	Store      uint64    `gorm:"column:store" json:"store"`             // 商品库存
	Image      string    `gorm:"column:image" json:"image"`             // 商品图片信息
	CategoryId uint64    `gorm:"column:category_id" json:"category_id"` // 商品类别id
	Type       string    `gorm:"type" json:"type"`                      // 商品种类名称
	Detail     string    `gorm:"column:detail" json:"detail"`           // 商品详情
	Status     uint8     `gorm:"column:status" json:"status"`           // 商品状态 1：正常 2：下架 3：审核
	ShopId     uint64    `gorm:"column:shop_id" json:"shop_id"`         // 商品所属店铺id
	CreateTime time.Time `gorm:"create_time" json:"create_time"`        // 创建时间
	UpdateTime time.Time `gorm:"update_time" json:"update_time"`        // 修改时间
}

// ProductListParam 查询商品列表参数模型
type ProductListParam struct {
	CategoryId uint64 `form:"category_id"`
}

// ProductList 查询商品列表模型
type ProductList struct {
	Id     uint64  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Store  uint64  `json:"store"`
	Image  string  `json:"image"`
	Type   string  `json:"type"`
	Detail string  `json:"detail"`
}

// ProductDetailParam 查询商品详情参数模型
type ProductDetailParam struct {
	ProductId uint64 `form:"product_id"`
}

// ProductDetail 查询商品详情模型
type ProductDetail struct {
	Id         uint64  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Store      uint64  `json:"store"`
	Image      string  `json:"image"`
	CategoryId uint64  `json:"category_id"`
	Type       string  `json:"type"`
	Detail     string  `json:"detail"`
	Status     uint8   `json:"status"`
}

// ProductCreateParam 创建商品参数模型
type ProductCreateParam struct {
	Name       string  `json:"name" form:"name" binding:"required"`
	Price      float64 `json:"price" form:"price" binding:"required,gt=0"`
	Store      uint64  `json:"store" form:"store" binging:"required,gt=0"`
	Image      string  `json:"image" form:"image"`
	CategoryId uint64  `json:"category_id" form:"category_id" binging:"required,gt=0"`
	Type       string  `json:"type" form:"type"`
	Detail     string  `json:"detail" form:"detail"`
	ShopId     uint64  `json:"shop_id" form:"shop_id"`
}

// ProductDeleteParam 删除商品参数模型
type ProductDeleteParam struct {
	ProductId uint64 `json:"product_id" form:"product_id"`
}

// ProductModifyParam 修改商品参数模型
type ProductModifyParam struct {
	Id         uint64  `json:"id" form:"id" binding:"required,gt=0"`
	Name       string  `json:"name" form:"name" binding:"required"`
	Price      float64 `json:"price" form:"price" binding:"required,gt=0"`
	Store      uint64  `json:"store" form:"store" binging:"required,gt=0"`
	Image      string  `json:"image" form:"image"`
	CategoryId uint64  `json:"category_id" form:"category_id" binging:"required,gt=0"`
	Type       string  `json:"type" form:"type"`
	Detail     string  `json:"detail" form:"detail"`
	ShopId     uint64  `json:"shop_id" form:"shop_id"`
}

// ProductSearchParam 搜索商品参数模型
type ProductSearchParam struct {
	KeyWord string `json:"key_word" form:"key_word"`
}

// ProductSearch 搜索商品信息模型
type ProductSearch struct {
	Id     uint64  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Image  string  `json:"image"`
	Type   string  `json:"type"`
	Detail string  `json:"detail"`
}

// ProductFlowParam 查看商品推荐流参数模型
type ProductFlowParam struct {
	ProductList []uint64 `json:"product_list" form:"product_list"`
}

// ProductFlow 商品推荐流信息模型
type ProductFlow struct {
	Id     uint64  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Image  string  `json:"image"`
	Type   string  `json:"type"`
	Detail string  `json:"detail"`
}
