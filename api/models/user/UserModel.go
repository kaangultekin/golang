package user

import "time"

type User struct {
	ID        uint      `json:"id"`
	Status    int       `json:"status" gorm:"default:1"`
	Name      string    `json:"name" gorm:"type:varchar(50);not null"`
	Surname   string    `json:"surname" gorm:"type:varchar(50);not null"`
	Email     string    `json:"email" gorm:"type:varchar(50);not null;unique"`
	Password  string    `json:"-" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:current_timestamp"`
	DeletedAt time.Time `json:"deleted_at" gorm:"default:null"`
}
