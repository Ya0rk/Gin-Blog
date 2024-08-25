package model

import (
	"Gin-Blog/utils/errmsg"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

// 实现 Tabler 接口来更改默认复数表名为单数
type Tabler interface {
	TableName() string
}

/*
定义了一个名为 User 的 Go 语言结构体，
用于表示一个数据库中的用户表。结构体中的
字段被绑定到数据库的列，并且使用了 GORM
提供的标签来指定字段的属性。
*/
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=20" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

// 将默认的 users 表名重写为单数 user
func (User) TableName() string {
	return "user"
}

/*
查询用户是否存在,返回code
*/
func CheckUser(name string) (code int) {
	var users User
	// 使用了 GORM 来查询数据库
	db.Select("id").Where("username = ?", name).First(&users)
	// 用户是否存在
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

/*
新增用户, 返回code，判断是否成功
*/
func CreateUser(data *User) int {
	//data.Password = ScryptPw(data.Password) // 也可以使用gorm的钩子函数实现覆盖原来的明文密码
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS // 200
}

/*
查询用户列表，分页查询
pageSize：每页显示的记录数
pageNum：当前页码
返回 []User 类型的切片
*/
func GetUsers(pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

/*
编辑用户：只能编辑username和role
*/
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

/*
删除用户：软删除
*/
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

/*
密码加密 ： scrypt加密方法
https://pkg.go.dev/golang.org/x/crypto/scrypt
*/

// 钩子函数
func (u *User) BeforeSave(*gorm.DB) error {
	u.Password = ScryptPw(u.Password)
	return nil
}

func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}
	// 进行加密处理
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	// 编码为 base64
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

// 登录验证
func CheckLogin(username, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCESS
}
