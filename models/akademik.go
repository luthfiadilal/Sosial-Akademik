package models

import "time"

type Course struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Code      string    `gorm:"size:20;unique;not null" json:"code"`
	Name      string    `gorm:"size:150;not null" json:"name"`
	SKS       int       `gorm:"not null" json:"sks"`
	Semester  int       `gorm:"not null" json:"semester"`
	CreatedAt time.Time `json:"created_at"`
}

type Class struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CourseID     uint      `gorm:"not null" json:"course_id"`
	ClassName    string    `gorm:"size:10" json:"class_name"`
	Capacity     int       `gorm:"default:30" json:"capacity"`
	Semester     int       `gorm:"not null" json:"semester"`
	AcademicYear string    `gorm:"size:9;not null" json:"academic_year"`
	CreatedAt    time.Time `json:"created_at"`

	Course Course `gorm:"foreignKey:CourseID" json:"course"`
}

type ClassLecturer struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	ClassID     uint   `gorm:"not null" json:"class_id"`
	LecturerID  uint   `gorm:"not null" json:"lecturer_id"`
	Role        string `gorm:"type:enum('utama','asisten');default:'utama'" json:"role"`

	Class    Class    `gorm:"foreignKey:ClassID" json:"class"`
	Lecturer Lecturer `gorm:"foreignKey:LecturerID" json:"lecturer"`
}

type Grade struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	StudentID  uint      `gorm:"not null" json:"student_id"`
	ClassID    uint      `gorm:"not null" json:"class_id"`
	NilaiAngka float64   `gorm:"type:decimal(5,2)" json:"nilai_angka"`
	NilaiHuruf string    `gorm:"size:2" json:"nilai_huruf"`
	InputBy    *uint     `json:"input_by"`
	InputAt    time.Time `json:"input_at"`

	Student Student   `gorm:"foreignKey:StudentID" json:"student"`
	Class   Class     `gorm:"foreignKey:ClassID" json:"class"`
	Lecturer *Lecturer `gorm:"foreignKey:InputBy" json:"lecturer_input"`
}

type Schedule struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	ClassID   uint   `gorm:"not null" json:"class_id"`
	Day       string `gorm:"type:enum('senin','selasa','rabu','kamis','jumat','sabtu')" json:"day"`
	StartTime string `gorm:"type:time" json:"start_time"`
	EndTime   string `gorm:"type:time" json:"end_time"`
	Room      string `gorm:"size:50" json:"room"`

	Class Class `gorm:"foreignKey:ClassID" json:"class"`
}
