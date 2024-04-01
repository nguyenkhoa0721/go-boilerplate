package constant

import "time"

const (
	EMAIL_VERIFIED_TOKEN_EXPIRE_AT    = time.Minute * 15
	SESSION_EXPIRE_AT                 = time.Hour * 24 * 15
	EMAIL_CONFIRMATION_CODE_EXPIRE_AT = time.Minute * 5
	TOTP_EXPIRE_AT                    = time.Minute * 5
)

const (
	MFA_EMAIL = "email"
	MFA_TOTP  = "totp"
)

const (
	TELEGRAM_PROVIDER = "telegram"
	GOOGLE_PROVIDER   = "google"
)

const (
	VERIFICATION_PHONE = "phone"
	VERIFICATION_EMAIL = "email"
)

const (
	AUTH_PHONE_NUMBER = "phoneNumber"
	AUTH_EMAIL        = "email"
)

const (
	METADATA_CHANGE_EMAIL = "changeEmail"
)

const (
	VAULT_USER_PRIV_KEY       = "user_privkey"
	VAULT_SYS_PRIV_KEY        = "sys_privkey"
	VAULT_TOTP_CODE           = "totp_code"
	VAULT_TOTP_RECOVERY_CODES = "totp_recovery_codes"
)

const (
	TEST_USER_ID           = "15000000000000000051"
	TEST_MAGIC_LINK_SECRET = "test_magic_link_secret"
)
