package models

import "time"

type Student struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	UserID         uint      `gorm:"not null" json:"user_id"`
	NPM            string    `gorm:"column:npm;size:30;unique;not null" json:"npm"`
	StudyProgramID uint      `gorm:"not null" json:"study_program_id"`
	Angkatan       string    `gorm:"type:year;not null" json:"angkatan"`
	AdvisorID      *uint     `json:"advisor_id"`
	Status         string    `gorm:"type:enum('aktif','cuti','lulus','dropout');default:'aktif'" json:"status"`
	CreatedAt      time.Time `json:"created_at"`

	// Relationships
	User         User          `gorm:"foreignKey:UserID" json:"user"`
	StudyProgram StudyProgram  `gorm:"foreignKey:StudyProgramID" json:"study_program"`
	Advisor      *Lecturer     `gorm:"foreignKey:AdvisorID" json:"advisor"`
}

type Lecturer struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	UserID         uint      `gorm:"not null" json:"user_id"`
	NIDN           string    `gorm:"column:nidn;size:30;unique;not null" json:"nidn"`
	StudyProgramID *uint     `json:"study_program_id"`
	CreatedAt      time.Time `json:"created_at"`

	// Relationships
	User         User          `gorm:"foreignKey:UserID" json:"user"`
	StudyProgram *StudyProgram `gorm:"foreignKey:StudyProgramID" json:"study_program"`
}
