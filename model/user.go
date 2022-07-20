package model

/*
	用户表映射模型
		1. 用户id，无符号整型，自增
		2. 用户名
		3. 密码
		4. 邮箱
		5. 头像
		6. 用户类型，整型，1表示卖家，2表示卖家
*/
type User struct {
	Id       uint64 `gorm:"column:id;primaryKey" json:"id"`
	UserName string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Email    string `gorm:"column:email" json:"email"`
	Avatar   string `gorm:"column:avatar" json:"avatar"`
	Class    int    `gorm:"column:class" json:"class"`
}

// UserLoginParam 用户登陆参数解析模型
type UserLoginParam struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// UserLogin 用户登陆信息传输模型
type UserLogin struct {
	Uid   uint64 `json:"uid"`
	Token string `json:"token"`
}

//// ViewUserInfoParam 用户信息查询参数解析模型
//type ViewUserInfoParam struct {
//	Id uint64 `form:"id"`
//}

// ViewUserInfo 用户信息查询传输模型
type ViewUserInfo struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}
