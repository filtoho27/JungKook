package mysql

type accountMethod struct{}
type AccountDB interface {
	GetMemberShipByUserName(userName string) (memberShip MemberShip, err error)
	CreateMemberShip(memberShip MemberShip) (err error)
}

var Account AccountDB

func GetAccountDB() AccountDB {
	if Account == nil {
		Account = &accountMethod{}
	}
	return Account
}

type MemberShip struct {
	UserID       int    `gorm:"type:int(10);column:UserID;UNSIGNED;not null;primary_key;autoIncrement:true"`
	UserName     string `gorm:"type:varchar(255);column:UserName;not null;primary_key"`
	PassWord     string `gorm:"type:varchar(255);column:PassWord;not null"`
	CreateTime   string `gorm:"type:datetime;column:CreateTime;not null;default:null"`
	ModifiedTime string `gorm:"type:timestamp;column:ModifiedTime;not null;default:null"`
}
