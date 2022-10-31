package shopping

import (
	Type "jungkook/commonType"
	// Error "jungkook/error"
	"jungkook/modules/mysql"
	"fmt"
)

func getShoppingList(module *Type.ModuleType) (err error) {
	productDB := mysql.GetProductDB()
	Category, err := productDB.GetCategory()
	if err != nil {
		return
	}
	Subcategory, err := productDB.GetSubcategory()
	if err != nil {
		return
	}

	fmt.Printf("%#v \n", Category)
	fmt.Printf("%#v \n", Subcategory)
	return
}
