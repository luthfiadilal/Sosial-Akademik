package models

import "time"

type CourseMaterial struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ClassID      uint      `gorm:"not null" json:"class_id"`
	Title        string    `gorm:"size:200;not null" json:"title"`
	Description  string    `gorm:"type:text" json:"description"`
	FileUrl      string    `gorm:"size:255" json:"file_url"`
	MaterialType string    `gorm:"type:enum('file','link','video');default:'file'" json:"material_type"`
	CreatedBy    *uint     `json:"created_by"`
	CreatedAt    time.Time `json:"created_at"`

	Class    Class     `gorm:"foreignKey:ClassID" json:"class"`
	Lecturer *Lecturer `gorm:"foreignKey:CreatedBy" json:"lecturer"`
}

type Assignment struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	ClassID     uint       `gorm:"not null" json:"class_id"`
	Title       string     `gorm:"size:200;not null" json:"title"`
	Description string     `gorm:"type:text" json:"description"`
	DueDate     *time.Time `json:"due_date"`
	MaxScore    float64    `gorm:"type:decimal(5,2);default:100.00" json:"max_score"`
	CreatedBy   *uint      `json:"created_by"`
	CreatedAt   time.Time  `json:"created_at"`

	Class    Class     `gorm:"foreignKey:ClassID" json:"class"`
	Lecturer *Lecturer `gorm:"foreignKey:CreatedBy" json:"lecturer"`
}

type AssignmentSubmission struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	AssignmentID uint       `gorm:"not null" json:"assignment_id"`
	StudentID    uint       `gorm:"not null" json:"student_id"`
	FileUrl      string     `gorm:"size:255" json:"file_url"`
	Note         string     `gorm:"type:text" json:"note"`
	SubmittedAt  time.Time  `json:"submitted_at"`
	Score        *float64   `gorm:"type:decimal(5,2)" json:"score"`
	Feedback     string     `gorm:"type:text" json:"feedback"`
	GradedBy     *uint      `json:"graded_by"`
	GradedAt     *time.Time `json:"graded_at"`

	Assignment Assignment `gorm:"foreignKey:AssignmentID" json:"assignment"`
	Student    Student    `gorm:"foreignKey:StudentID" json:"student"`
	Lecturer   *Lecturer  `gorm:"foreignKey:GradedBy" json:"grader"`
}

type Forum struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ClassID   uint      `gorm:"not null" json:"class_id"`
	Title     string    `gorm:"size:200" json:"title"`
	CreatedBy *uint     `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`

	Class Class `gorm:"foreignKey:ClassID" json:"class"`
	User  *User `gorm:"foreignKey:CreatedBy" json:"user"`
}

type ForumThread struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ForumID   uint      `gorm:"not null" json:"forum_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Content   string    `gorm:"type:text" json:"content"`
	CreatedAt time.Time `json:"created_at"`

	Forum Forum `gorm:"foreignKey:ForumID" json:"forum"`
	User  User  `gorm:"foreignKey:UserID" json:"user"`
}

type ForumComment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ThreadID  uint      `gorm:"not null" json:"thread_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Comment   string    `gorm:"type:text" json:"comment"`
	CreatedAt time.Time `json:"created_at"`

	Thread ForumThread `gorm:"foreignKey:ThreadID" json:"thread"`
	User   User        `gorm:"foreignKey:UserID" json:"user"`
}

type Quiz struct {
	ID              uint       `gorm:"primaryKey" json:"id"`
	ClassID         uint       `gorm:"not null" json:"class_id"`
	Title           string     `gorm:"size:200" json:"title"`
	Description     string     `gorm:"type:text" json:"description"`
	StartTime       *time.Time `json:"start_time"`
	EndTime         *time.Time `json:"end_time"`
	DurationMinutes int        `json:"duration_minutes"`
	CreatedBy       *uint      `json:"created_by"`
	CreatedAt       time.Time  `json:"created_at"`

	Class    Class     `gorm:"foreignKey:ClassID" json:"class"`
	Lecturer *Lecturer `gorm:"foreignKey:CreatedBy" json:"lecturer"`
}

type QuizQuestion struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	QuizID   uint   `gorm:"not null" json:"quiz_id"`
	Question string `gorm:"type:text;not null" json:"question"`
	Type     string `gorm:"type:enum('multiple_choice','essay');default:'multiple_choice'" json:"type"`

	Quiz Quiz `gorm:"foreignKey:QuizID" json:"quiz"`
}

type QuizOption struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	QuestionID uint   `gorm:"not null" json:"question_id"`
	OptionText string `gorm:"size:255" json:"option_text"`
	IsCorrect  bool   `gorm:"default:false" json:"is_correct"`

	Question QuizQuestion `gorm:"foreignKey:QuestionID" json:"question"`
}

type QuizAnswer struct {
	ID               uint     `gorm:"primaryKey" json:"id"`
	QuestionID       uint     `gorm:"not null" json:"question_id"`
	StudentID        uint     `gorm:"not null" json:"student_id"`
	SelectedOptionID *uint    `json:"selected_option_id"`
	AnswerText       string   `gorm:"type:text" json:"answer_text"`
	IsCorrect        *bool    `json:"is_correct"`
	Score            *float64 `gorm:"type:decimal(5,2)" json:"score"`

	Question       QuizQuestion `gorm:"foreignKey:QuestionID" json:"question"`
	Student        Student      `gorm:"foreignKey:StudentID" json:"student"`
	SelectedOption *QuizOption  `gorm:"foreignKey:SelectedOptionID" json:"selected_option"`
}

type QuizResult struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	QuizID      uint      `gorm:"not null" json:"quiz_id"`
	StudentID   uint      `gorm:"not null" json:"student_id"`
	TotalScore  *float64  `gorm:"type:decimal(5,2)" json:"total_score"`
	SubmittedAt time.Time `json:"submitted_at"`

	Quiz    Quiz    `gorm:"foreignKey:QuizID" json:"quiz"`
	Student Student `gorm:"foreignKey:StudentID" json:"student"`
}
