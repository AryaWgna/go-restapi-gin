package models

type Mahasiswa struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	NamaMahasiswa string `gorm:"type:varchar(300)" json:"nama_mahasiswa"`
	Prodi   string `gorm:"type:text" json:"prodi"`
}