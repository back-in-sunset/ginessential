package do

// User ..
type User struct {
	Name      string  `json:"name" db:"name" gorm:"column:name;type:varchar(20);not null;"`                                    // 用户名
	Password  string  `json:"password,omitempty" db:"password" gorm:"->:false;<-;column:password;type:varchar(100);not null;"` // 密码
	Telephone string  `json:"telephone" db:"telephone" gorm:"column:telephone;type:varchar(110);not null;unique;"`             // 手机号
	Email     *string `json:"email" db:"email" gorm:"column:email;size:255;index;"`                                            // 邮箱
}
