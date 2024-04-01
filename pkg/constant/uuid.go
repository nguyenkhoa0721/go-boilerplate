package constant

type UuidType uint64

const (
	UUID_USER = iota
	UUID_SESSION
	UUID_TOKEN
	UUID_AA_WALLET
	UUID_ON_CHAIN_TRANSACTION
	UUID_SOCIAL_ACCOUNT
	UUID_UTILITY
)