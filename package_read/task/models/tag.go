package models

func CleanAllTag()bool{
	db.Unscoped().Where("deleted_on != ?",0).Delete(&Tag{})
	return true
}
