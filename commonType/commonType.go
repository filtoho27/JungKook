package commontype

type ModuleType struct {
	Redis    RedisInterface
	Acc      AccTypeInterface
	Time     TimeInterface
	Bcrypt   BcryptInterface
	Google   GoogleInterface
	Portal   PortalInterface
	Vendor   int
	Operator string
}

type ExcelType struct {
	FileName  string
	SheetName string
	Header    []string
	Data      [][]string
}
