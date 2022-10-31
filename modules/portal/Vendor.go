package portal

import "fmt"

func (pl *portalLib) SetHallVendorRedis() (result string, err error) {
	url := fmt.Sprintf("%s/amfphp/json.php/Live.Hall.setHallVendorRedis", pl.host)
	resp, err := pl.client.R().Get(url)
	result = resp.String()
	err = portalResponse(resp, "SetHallVendorRedis", err, result)
	return
}
