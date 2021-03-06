package session

import (
	"taylz.io/keygen"
	"taylz.io/types"
	"taylz.io/z/charset"
)

// Settings is configuration for sessions
type Settings struct {
	CookieID string
	Secure   bool
	Strict   bool
	Keygen   keygen.I
	Lifetime types.Duration
}

// SettingsDefault is a var Settings for using in basic case
var SettingsDefault = Settings{
	CookieID: "SessionID",
	Strict:   true,
	Keygen: &keygen.Settings{
		KeySize: 12,
		CharSet: charset.AlphaCapitalNumeric,
		Rand:    keygen.DefaultSettings.Rand,
	},
	Lifetime: 12 * types.Hour,
}
