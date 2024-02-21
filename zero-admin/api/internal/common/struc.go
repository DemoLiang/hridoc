package common

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// 用户
type User struct {
	ID       string         // ID编号
	Name     string         // 用户名
	NickName sql.NullString // 昵称
	Avatar   sql.NullString // 头像
	Password string         // 密码
	Salt     string         // 加密盐
	Email    sql.NullString // 邮箱
	Mobile   sql.NullString // 手机号
	Status   int64          // 状态  0：禁用   1：正常

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// 证书记录
type Certificate struct {
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt
	ID                int64
	Name              string //证书名字
	CertificateTypeID string //证书分类ID
	Path              string //存储路径
	UserID            string //持有用户
}

// 证书分类
type CertificateType struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	ID        int64
	Name      string //分类名
}
