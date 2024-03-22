package validator

const (
	defaultUsernameRules = `^[a-z0-9]{4,15}$`
	defaultPasswordRules = `^[a-zA-Z0-9]{6,15}$`
)

var usernameBlackList = []string{
	"administrator",
	"admin",
	"root",
	"marketing",
	"postmaster",
	"security",
	"support",
	"sysadmin",
	"trouble",
	"usenet",
	"webmaster",
	"production",
}
