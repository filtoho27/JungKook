package mysql

type productMethod struct{}
type ProductDB interface {
	GetCategory() (category []Category, err error)
	GetSubcategory() (subcategory []Subcategory, err error)
}

var Product ProductDB

func GetProductDB() ProductDB {
	if Product == nil {
		Product = &productMethod{}
	}
	return Product
}

type Category struct {
	ID           int    `gorm:"type:int(10);column:ID;UNSIGNED;not null;primary_key;autoIncrement:true"`
	CategoryName string `gorm:"type:varchar(255);column:CategoryName;not null"`
	Sort         int    `gorm:"type:int(10);column:Sort;UNSIGNED;not null"`
}

type Subcategory struct {
	ID              int    `gorm:"type:int(10);column:ID;UNSIGNED;not null;primary_key;autoIncrement:true"`
	SubcategoryName string `gorm:"type:varchar(255);column:SubcategoryName;not null"`
	CategoryID      int    `gorm:"type:int(10);column:CategoryID;UNSIGNED;not null;primary_key"`
	Sort            int    `gorm:"type:int(10);column:Sort;UNSIGNED;not null"`
}
