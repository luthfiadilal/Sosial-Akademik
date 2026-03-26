package models

import "time"

type Faculty struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Code      string    `gorm:"size:10;unique;not null" json:"code"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type StudyProgram struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FacultyID uint      `gorm:"not null" json:"faculty_id"`
	Code      string    `gorm:"size:10;unique;not null" json:"code"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Degree    string    `gorm:"type:enum('D3','S1','S2','S3');not null" json:"degree"`
	CreatedAt time.Time `json:"created_at"`

	Faculty Faculty `gorm:"foreignKey:FacultyID" json:"faculty"`
}

type Curriculum struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	StudyProgramID uint      `gorm:"not null" json:"study_program_id"`
	Year           string    `gorm:"size:50;not null" json:"year"`
	IsActive       bool      `gorm:"default:true" json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`

	StudyProgram StudyProgram `gorm:"foreignKey:StudyProgramID" json:"study_program"`
}

type CurriculumDetail struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	CurriculumID uint       `gorm:"not null" json:"curriculum_id"`
	CourseID     uint       `gorm:"not null" json:"course_id"`
	
	Curriculum Curriculum `gorm:"foreignKey:CurriculumID" json:"curriculum"`
	Course     Course     `gorm:"foreignKey:CourseID" json:"course"`
}
