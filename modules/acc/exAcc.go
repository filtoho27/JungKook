package acc

import (
	"fmt"
)

func (acc *accLib) GetDomainData(hallID int) (result string, err error) {
	url := fmt.Sprintf("%s/api/domain/%d", acc.host, hallID)
	resp, err := acc.client.R().Get(url)
	result = resp.String()
	err = accResponse(resp, "GetDomainData", err, result)
	return
}
