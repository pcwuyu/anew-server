package response

import (
	"anew-server/models"
)

// User login response structure
type LoginResp struct {
	Token   string           `json:"token"`   // jwt令牌
	Expires models.LocalTime `json:"expires"` // 过期时间, 秒
}

// 用户返回角色信息
type UserRolesResp struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Keyword string `json:"keyword"`
	Status  *bool  `json:"status"`
}

// 用户信息响应
type UserInfoResp struct {
	Id       uint        `json:"id"`
	Username string      `json:"username"`
	Mobile   string      `json:"mobile"`
	Avatar   string      `json:"avatar"`
	Name     string      `json:"name"`
	Email    string      `json:"mail"`
	Roles    []UserRolesResp `json:"roles"`
	LoginResp
}

// 用户列表信息响应, 字段含义见models.SysUser
type UserListResp struct {
	Id        uint             `json:"id"`
	Username  string           `json:"username"`
	Mobile    string           `json:"mobile"`
	Avatar    string           `json:"avatar"`
	Name      string           `json:"name"`
	Email     string           `json:"mail"`
	Status    *bool            `json:"status"`
	Roles    []UserRolesResp `json:"roles"`
	Creator   string           `json:"creator"`
	CreatedAt models.LocalTime `json:"createdAt"`
	UpdatedAt models.LocalTime `json:"updatedAt"`
}
