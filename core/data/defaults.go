package data

import (
	"github.com/owncast/owncast/config"
	"github.com/owncast/owncast/models"
)

// HasPopulatedDefaults will determine if the defaults have been inserted into the database.
func HasPopulatedDefaults() bool {
	hasPopulated, err := _datastore.GetBool("HAS_POPULATED_DEFAULTS")
	if err != nil {
		return false
	}
	return hasPopulated
}

// PopulateDefaults will set default values in the database.
func PopulateDefaults() {
	defaults := config.GetDefaults()

	if HasPopulatedDefaults() {
		return
	}

	_ = SetStreamKey(defaults.StreamKey)
	_ = SetHTTPPortNumber(float64(defaults.WebServerPort))
	_ = SetRTMPPortNumber(float64(defaults.RTMPServerPort))
	_ = SetLogoPath(defaults.Logo)
	_ = SetServerMetadataTags([]string{"UFAXLIVE", "UFAPRO888"})
	_ = SetServerSummary("UFABET คาสิโน 18+ แทงบอลกำไรดี มีพริตตี้ไลฟ์สดโชว์ Visit https://ufapro888s.info/ more.")
	_ = SetServerWelcomeMessage("เว็บแทงบอล UFAPRO888 แห่งนี้ไม่ได้มีดีแค่การเดิมพัน พนันออนไลน์ อย่างเดียว ทางเราได้เพิ่มบริการพิเศษเข้าไปอีก 1 บริการ คือการสมัครเข้า กลุ่มลับพริตตี้ไลฟ์สด20+")
	_ = SetServerName("UFAXLIVE")
	_ = SetStreamKey(defaults.StreamKey)
	_ = SetExtraPageBodyContent("UFABET คาสิโน 18+ แทงบอลกำไรดี มีพริตตี้ไลฟ์สดโชว์")
	_ = SetSocialHandles([]models.SocialHandle{
		{
			Platform: "facebook",
			URL:      "https://www.facebook.com/ufax249999/",
		},
	})

	_datastore.warmCache()
	_ = _datastore.SetBool("HAS_POPULATED_DEFAULTS", true)
}
