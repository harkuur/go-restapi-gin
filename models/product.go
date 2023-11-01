package models

type Product struct {
	Id int64 `gorm:"primaryKey" json:"id"`
	NamaProduct string `gorm:"tipe:varchar(300)" json:"nama_product"`
	Deskripsi string `gorm:"tipe:text" json:"deskripsi"`

}