package user

type RegisterRequest struct {
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required,min=6"`
	Role           string `json:"role" binding:"required,oneof=member trainer admin"`
	MembershipTier string `json:"membership_tier" binding:"required,oneof=basic premium vip"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Role           string `json:"role"`
	MembershipTier string `json:"membership_tier"`
	Active         bool   `json:"active"`
}

func ToUserResponse(u *User) *UserResponse {
	return &UserResponse{
		ID:             u.ID,
		Name:           u.Name,
		Email:          u.Email,
		Role:           u.Role,
		MembershipTier: u.MembershipTier,
		Active:         u.Active,
	}
}
