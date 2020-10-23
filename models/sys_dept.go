package models

// 系统部门表
type SysDept struct {
	Model
	Name     string    `gorm:"comment:'部门名称';size:128" json:"name"`
	Status   *bool     `gorm:"type:tinyint(1);default:1;comment:'菜单状态(正常/禁用, 默认正常)'" json:"status"` // 由于设置了默认值, 这里使用ptr, 可避免赋值失败
	Creator  string    `gorm:"comment:'创建人';size:128" json:"creator"`
	Sort       int       `gorm:"type:int(3);comment:'菜单顺序(同级菜单, 从0开始, 越小显示越靠前)'" json:"sort"`
	ParentId uint      `gorm:"default:0;comment:'父菜单编号(编号为0时表示根菜单)'" json:"parentId"`
	Children []SysDept `gorm:"-" json:"children"` // 下属部门集合
	Users    []SysUser `gorm:"foreignkey:DeptId"` // 一个部门有多个user
}

func (m SysDept) TableName() string {
	return m.Model.TableName("sys_dept")
}