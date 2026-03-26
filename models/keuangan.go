package models

import "time"

type Billing struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	StudentID    uint       `gorm:"not null" json:"student_id"`
	Semester     int        `gorm:"not null" json:"semester"`
	AcademicYear string     `gorm:"size:9;not null" json:"academic_year"`
	Amount       float64    `gorm:"type:decimal(12,2);not null" json:"amount"`
	Status       string     `gorm:"type:enum('unpaid','partial','paid');default:'unpaid'" json:"status"`
	DueDate      *time.Time `json:"due_date"`
	CreatedAt    time.Time  `json:"created_at"`

	Student Student `gorm:"foreignKey:StudentID" json:"student"`
}

type Payment struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	BillingID   uint      `gorm:"not null" json:"billing_id"`
	Amount      float64   `gorm:"type:decimal(12,2);not null" json:"amount"`
	PaymentDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"payment_date"`
	Method      string    `gorm:"size:50" json:"method"`
	ReferenceNo string    `gorm:"size:100" json:"reference_no"`
	Status      string    `gorm:"type:enum('pending','success','failed');default:'pending'" json:"status"`

	Billing Billing `gorm:"foreignKey:BillingID" json:"billing"`
}

type Scholarship struct {
	ID    uint    `gorm:"primaryKey" json:"id"`
	Name  string  `gorm:"size:100" json:"name"`
	Type  string  `gorm:"type:enum('percentage','fixed')" json:"type"`
	Value float64 `gorm:"type:decimal(10,2)" json:"value"`
}

type StudentScholarship struct {
	ID            uint `gorm:"primaryKey" json:"id"`
	StudentID     uint `gorm:"not null" json:"student_id"`
	ScholarshipID uint `gorm:"not null" json:"scholarship_id"`

	Student     Student     `gorm:"foreignKey:StudentID" json:"student"`
	Scholarship Scholarship `gorm:"foreignKey:ScholarshipID" json:"scholarship"`
}
