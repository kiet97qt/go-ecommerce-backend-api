package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents an application user with roles (many-to-many).
type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Username string    `gorm:"size:255;uniqueIndex;not null" json:"username"`
	IsActive bool      `gorm:"default:true" json:"isActive"`
	Roles    []Role    `gorm:"many2many:user_roles;" json:"roles"`
}

// Role represents a permission role that can be assigned to users.
type Role struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	RoleName string `gorm:"size:100;uniqueIndex;not null" json:"roleName"`
	RoleNote string `gorm:"size:255" json:"roleNote"`
	Users    []User `gorm:"many2many:user_roles;" json:"-"`
}
