package mysql

type coreMethod struct{}
type CoreDB interface {
	GetUserInfo(acc string) (userInfo UserInfo, err error)
}

var Core CoreDB

func GetCoreDB() CoreDB {
	if Core == nil {
		Core = &coreMethod{}
	}
	return Core
}

type RD3Admin struct {
	ID              int    `gorm:"type:int(10);column:ID;NOT NULL;primary_key;autoIncrement:true"`
	Name            string `gorm:"type:varchar(191);column:Name;NOT NULL"`
	Password        string `gorm:"type:varchar(191);column:Password;NOT NULL"`
	Google2faSecret string `gorm:"type:text;column:Google2faSecret;default:''"`
	Enable          int    `gorm:"type:tinyint(1);column:Enable;UNSIGNED;NOT NULL;default:1"`
	LoginError      int    `gorm:"type:tinyint(1);column:LoginError;UNSIGNED;NOT NULL;default:0"`
}
