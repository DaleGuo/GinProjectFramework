package db

type Users struct {
	Id int `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Role string `gorm:"type:enum('professor','student');column:role"`
}

//检查用户名和密码
func Authen(userName string, password string) bool {
	var users Users
	err := db.Where("userName=? && password=?", userName,password).First(&users).Error
	if err != nil {
		return false
	}
	return true
}

//获取用户角色
func QueryRole(userName string) string {
	var users Users
	db.First(&users,"userName=?",userName)
	return users.Role
}