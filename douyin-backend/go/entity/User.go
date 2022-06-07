package entity

// 数据库表明自定义，默认为model的复数形式，比如这里默认为 users
func (User) TableName() string {
	return "user"
}

type User struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Password       string `json:"password"`
	Fllow_Count    int    `json:"follow_count"`
	Follower_Count int    `json:"follower_count"`
	Is_follow      int    `json:"is_follow"`
}
