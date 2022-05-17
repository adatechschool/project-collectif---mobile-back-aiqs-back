package model

import "gorm.io/gorm"

type Surf struct {
	/* gorm.Model */
	ID                      string  `gorm:"primaryKey"`
	SurfBreak               string  `json:"SurfBreak" gorm:"<-"`
	DifficultyLevel         int64   `json:"DifficultyLevel" gorm:"<-"`
	Destination             string  `json:"Destination" gorm:"<-"`
	MagicSeaweedLink        string  `json:"MagicSeaweedLink" gorm:"<-"`
	Photos                  string  `json:"Photos" gorm:"<-"`
	PeakSurfSeasonBegins    string  `json:"PeakSurfSeasonBegins" gorm:"<-"`
	DestinationStateCountry string  `json:"DestinationStateCountry" gorm:"<-"`
	PeakSurfSeasonEnds      string  `json:"PeakSurfSeasonEnds" gorm:"<-"`
	Address                 string  `json:"Address" gorm:"<-"`
	Lat                     float64 `json:"Lat" gorm:"<-"`
	Lng                     float64 `json:"Lng" gorm:"<-"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Surf{})
	return db
}
