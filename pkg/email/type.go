package email

// Verify Account

type AuthorizeNewDeviceEmail struct {
	Email       string
	ConfirmLink string
}

// 2FA Email

// Setup 2FA
type Setup2FASuccessEmail struct {
	Email string
}
