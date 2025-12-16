package user

import "time"

const (
	RoleMember  = "member"
	RoleTrainer = "trainer"
	RoleAdmin   = "admin"

	MembershipBasic   = "basic"
	MembershipPremium = "premium"
	MembershipVIP     = "vip"
)

type User struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Name           string    `json:"name"`
	Email          string    `gorm:"uniqueIndex" json:"email"`
	PasswordHash   string    `json:"-"`
	Role           string    `gorm:"type:varchar(20)" json:"role"`
	MembershipTier string    `gorm:"type:varchar(20)" json:"membership_tier"`
	Active         bool      `json:"active"`
}
