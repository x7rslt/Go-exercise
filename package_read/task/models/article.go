package models

func CleanAllArticle()bool{
	db.Unscoped().Where("deleted_on!=?",0).Delete(&Article{})
	return true
}
