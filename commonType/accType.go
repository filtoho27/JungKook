package commontype

type AccInterface interface {
	GetDomainData(hallID int) (result string, err error)
}

type AccTypeInterface interface {
	GetByVendor(vendor int) (acc AccInterface)
}

type AccResult struct {
	Result string `json:"result"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}
