package entity

type User struct {
	ID       string `gorm:"primaryKey,type:varchar(20)" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
