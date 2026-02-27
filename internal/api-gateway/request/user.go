package request

// SendTextMessage 绑定 JSON
type SendTextMessage struct {
	Phone string `form:"phone" json:"phone" xml:"phone"  binding:"required"`
}

// Register 绑定 JSON
type Register struct {
	Phone            string `form:"phone" json:"phone" xml:"phone"  binding:"required"`
	VerificationCode string `form:"verification_code" json:"verification_code" xml:"verification_code"  binding:"required"`
	Password         string `form:"password" json:"password" xml:"password"  binding:"required"`
}

// Login 绑定 JSON
type Login struct {
	Phone            string `form:"phone" json:"phone" xml:"phone"  binding:"required"`
	Password         string `form:"password" json:"password" xml:"password"  binding:"required"`
	VerificationCode string `form:"verification_code" json:"verification_code" xml:"verification_code"  binding:"required"`
}
