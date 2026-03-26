package models

import "time"

type KrsApprovalLog struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	KrsID      uint      `gorm:"not null" json:"krs_id"`
	LecturerID uint      `gorm:"not null" json:"lecturer_id"`
	Status     string    `gorm:"type:enum('diajukan','disetujui','ditolak')" json:"status"`
	Note       string    `gorm:"type:text" json:"note"`
	CreatedAt  time.Time `json:"created_at"`

	Krs      Krs      `gorm:"foreignKey:KrsID" json:"krs"`
	Lecturer Lecturer `gorm:"foreignKey:LecturerID" json:"lecturer"`
}

type StudentStatusLog struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	StudentID uint       `gorm:"not null" json:"student_id"`
	Status    string     `gorm:"type:enum('aktif','cuti','lulus','dropout')" json:"status"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`

	Student Student `gorm:"foreignKey:StudentID" json:"student"`
}

type Notification struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Title     string    `gorm:"size:150" json:"title"`
	Message   string    `gorm:"type:text" json:"message"`
	IsRead    bool      `gorm:"default:false" json:"is_read"`
	CreatedAt time.Time `json:"created_at"`

	User User `gorm:"foreignKey:UserID" json:"user"`
}
