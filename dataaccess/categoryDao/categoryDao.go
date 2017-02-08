package categoryDao

import "github.com/NiciiA/GoGraphQL/domain/model/categoryModel"

var CategoryList map[string]categoryModel.Category = make(map[string]categoryModel.Category)

func GetByKey(key string) categoryModel.Category {
	return CategoryList[key]
}

func AddCategory(c categoryModel.Category) {
	CategoryList[c.Name] = c
}