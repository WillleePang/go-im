package logic

// User 用户结构体
type User struct {
	// ID
	Id int32 `json:"id"`
	// 昵称
	Name string `json:"name"`
	// 邮箱
	Mail string `json:"mail"`
	// 手机
	Phone string `json:"phone"`
}

// users 用户列表
var users = map[int32]*User{
	1: {
		Id:    1,
		Name:  "张三",
		Mail:  "zhangsan@mail.com",
		Phone: "12345678901",
	},
	2: {
		Id:    2,
		Name:  "李四",
		Mail:  "lisi@mail.com",
		Phone: "12345678902",
	},
	3: {
		Id:    3,
		Name:  "王五",
		Mail:  "wangwu@mail.com",
		Phone: "12345678903",
	},
	4: {
		Id:    4,
		Name:  "赵六",
		Mail:  "zhaoliu@mail.com",
		Phone: "12345678904",
	},
}
