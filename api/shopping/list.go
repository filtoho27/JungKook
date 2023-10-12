package shopping

import (
	Type "jungkook/commonType"
	// Error "jungkook/error"
	"jungkook/modules/mysql"
	"fmt"
)

type shoppingListType struct {
	CategoryID   int               `json:"categoryID"`
	CategoryName string            `json:"categoryName"`
	Subcategory  []subcategoryType `json:"subcategory"`
	Product      []productType     `json:"product"`
}

type subcategoryType struct {
	SubcategoryID   int    `json:"subcategoryID"`
	SubcategoryName string `json:"subcategoryName"`
}

type productType struct {
	ProductID     int    `json:"productID"`
	CategoryID    int    `json:"categoryID"`
	SubcategoryID int    `json:"subcategoryID"`
	ProductName   string `json:"productName"`
	Description   string `json:"description"`
	Price         int    `json:"price"`
	Image         string `json:"image"`
}

func getShoppingList(module *Type.ModuleType) (shoppingList []shoppingListType, err error) {
	productDB := mysql.GetProductDB()
	// 分類
	category, err := productDB.GetCategory()
	if err != nil {
		return
	}
	// 子分類
	subcategory, err := productDB.GetSubcategory()
	if err != nil {
		return
	}
	// 整理同分類的子分類
	subcategoryInfo := make(map[int][]subcategoryType, len(subcategory))
	for _, v := range subcategory {
		subcategoryInfo[v.CategoryID] = append(subcategoryInfo[v.CategoryID], subcategoryType{v.ID, v.SubcategoryName})
	}
	// 商品
	product, err := productDB.GetProduct()
	if err != nil {
		return
	}
	// 整理同分類的商品
	productInfo := make(map[int][]productType, len(product))
	for _, v := range product {
		data := productType{
			v.ID,
			v.CategoryID,
			v.SubcategoryID,
			v.ProductName,
			v.Description,
			v.Price,
			v.Image,
		}
		productInfo[v.CategoryID] = append(productInfo[v.CategoryID], data)
	}

	for _, v := range category {
		data := shoppingListType{v.ID, v.CategoryName, []subcategoryType{}, []productType{}}
		subcategoryData, subcategoryExist := subcategoryInfo[v.ID]
		if subcategoryExist {
			data.Subcategory = subcategoryData
		}
		productData, productExist := productInfo[v.ID]
		if productExist {
			data.Product = productData
		}
		shoppingList = append(shoppingList, data)
	}
	fmt.Printf("%#v \n", subcategoryInfo)
	return
}
