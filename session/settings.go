package session

import (
	"time"

	"taylz.io/keygen"
)

// Settings is configuration for Cache cookies usage
type Settings struct {
	CookieID string
	Keygen   keygen.I
	Lifetime time.Duration
}

// SettingsDefault is a var Settings for using in basic case
var SettingsDefault = Settings{
	CookieID: "SessionID",
	Keygen: &keygen.Settings{
		KeySize: 12,
		CharSet: keygen.CharsetAlphaCapitalNumeric,
		Rand:    keygen.DefaultSettings.Rand,
	},
}
