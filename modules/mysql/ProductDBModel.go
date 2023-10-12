package mysql

type productMethod struct{}
type ProductDB interface {
	GetCategory() (category []CategoryType, err error)
	GetSubcategory() (subcategory []SubcategoryType, err error)
	GetProduct() (product []ProductType, err error)
}

var Product ProductDB

func GetProductDB() ProductDB {
	if Product == nil {
		Product = &productMethod{}
	}
	return Product
}

type CategoryType struct {
	ID           int    `gorm:"type:int(10);column:ID;UNSIGNED;not null;primary_key;autoIncrement:true"`
	CategoryName string `gorm:"type:varchar(255);column:CategoryName;not null"`
	Sort         int    `gorm:"type:int(10);column:Sort;UNSIGNED;not null"`
}

type SubcategoryType struct {
	ID              int    `gorm:"type:int(10);column:ID;UNSIGNED;not null;primary_key;autoIncrement:true"`
	SubcategoryName string `gorm:"type:varchar(255);column:SubcategoryName;not null"`
	CategoryID      int    `gorm:"type:int(10);column:CategoryID;UNSIGNED;not null;primary_key"`
	Sort            int    `gorm:"type:int(10);column:Sort;UNSIGNED;not null"`
}

type ProductType struct {
	ID            int    `gorm:"type:int(10);column:ID;UNSIGNED;not null;primary_key;autoIncrement:true"`
	CategoryID    int    `gorm:"type:int(10);column:CategoryID;UNSIGNED;not null"`
	SubcategoryID int    `gorm:"type:int(10);column:SubcategoryID;UNSIGNED;not null"`
	ProductName   string `gorm:"type:varchar(255);column:ProductName;not null"`
	Description   string `gorm:"type:text;column:Description;not null"`
	Price         int    `gorm:"type:int(10);column:Price;UNSIGNED;not null"`
	Image         string `gorm:"type:text;column:Image;not null"`
}
