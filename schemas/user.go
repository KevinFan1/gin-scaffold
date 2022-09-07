package schemas

type UserAddDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleId   uint   `json:"role_id" binding:"required"`
}
