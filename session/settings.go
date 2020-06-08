package session

import (
	"net/http"
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

// ReadCookie returns the Request cookie value associated to the CookieID setting
func (settings Settings) ReadCookie(r *http.Request) string {
	cookie, err := r.Cookie(settings.CookieID)
	if err != nil {
		return ""
	}
	return cookie.Value
}
