package models

import "time"

type Krs struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	StudentID    uint       `gorm:"not null" json:"student_id"`
	Semester     int        `gorm:"not null" json:"semester"`
	AcademicYear string     `gorm:"size:9;not null" json:"academic_year"`
	Status       string     `gorm:"type:enum('draft','diajukan','disetujui','ditolak');default:'draft'" json:"status"`
	ApprovedBy   *uint      `json:"approved_by"`
	ApprovedAt   *time.Time `json:"approved_at"`

	Student  Student     `gorm:"foreignKey:StudentID" json:"student"`
	Approver *Lecturer   `gorm:"foreignKey:ApprovedBy" json:"approver"`
}

type KrsDetail struct {
	ID      uint `gorm:"primaryKey" json:"id"`
	KrsID   uint `gorm:"not null" json:"krs_id"`
	ClassID uint `gorm:"not null" json:"class_id"`

	Krs   Krs   `gorm:"foreignKey:KrsID" json:"krs"`
	Class Class `gorm:"foreignKey:ClassID" json:"class"`
}
