package request

type AddForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Status   string `json:"status" bidding:"required"`
}
