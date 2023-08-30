package commontype

type ModuleType struct {
	Redis    RedisInterface
	Time     TimeInterface
	Bcrypt   BcryptInterface
	Google   GoogleInterface
}

type ExcelType struct {
	FileName  string
	SheetName string
	Header    []string
	Data      [][]string
}
