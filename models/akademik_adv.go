package models

import "time"

type Attendance struct {
	ID            uint       `gorm:"primaryKey" json:"id"`
	StudentID     uint       `gorm:"not null" json:"student_id"`
	ClassID       uint       `gorm:"not null" json:"class_id"`
	MeetingNumber int        `gorm:"not null" json:"meeting_number"`
	Date          *time.Time `json:"date"`
	Status        string     `gorm:"type:enum('hadir','izin','sakit','alfa');default:'hadir'" json:"status"`

	Student Student `gorm:"foreignKey:StudentID" json:"student"`
	Class   Class   `gorm:"foreignKey:ClassID" json:"class"`
}

type Transcript struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	StudentID uint      `gorm:"not null" json:"student_id"`
	TotalSks  int       `gorm:"default:0" json:"total_sks"`
	Ipk       float64   `gorm:"type:decimal(3,2)" json:"ipk"`
	UpdatedAt time.Time `json:"updated_at"`

	Student Student `gorm:"foreignKey:StudentID" json:"student"`
}

type TranscriptDetail struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	TranscriptID uint    `gorm:"not null" json:"transcript_id"`
	CourseID     uint    `gorm:"not null" json:"course_id"`
	ClassID      *uint   `json:"class_id"`
	NilaiHuruf   string  `gorm:"size:2" json:"nilai_huruf"`
	NilaiAngka   float64 `gorm:"type:decimal(5,2)" json:"nilai_angka"`
	Sks          int     `json:"sks"`

	Transcript Transcript `gorm:"foreignKey:TranscriptID" json:"transcript"`
	Course     Course     `gorm:"foreignKey:CourseID" json:"course"`
}

type StudentGpa struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	StudentID    uint    `gorm:"not null" json:"student_id"`
	Semester     int     `json:"semester"`
	AcademicYear string  `gorm:"size:9" json:"academic_year"`
	Ips          float64 `gorm:"type:decimal(3,2)" json:"ips"`
	Sks          int     `json:"sks"`

	Student Student `gorm:"foreignKey:StudentID" json:"student"`
}
