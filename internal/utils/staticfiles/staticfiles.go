package staticfiles

import (
	"github.com/CaSe-SuperCashback/go-utilities/appimage"
)

type Endpoint struct {
	cdn string
}

var endpoint = Endpoint{}

func Init(cdnEndpoint string) {
	endpoint.cdn = cdnEndpoint
}

func GetImageURL(image *appimage.Image, isMobile bool) string {
	url := ""
	if image != nil {
		url = image.GetURL(endpoint.cdn, isMobile)
	}
	return url
}
