package models

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func CleanAllTag()bool{
	db.Unscoped().Where("deleted_on != ?",0).Delete(&Tag{})
	return true
}
