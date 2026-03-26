package models

import "time"

type UserProfile struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	UserID         uint      `gorm:"not null;unique" json:"user_id"`
	Username       string    `gorm:"size:50;unique" json:"username"`
	Bio            string    `gorm:"type:text" json:"bio"`
	ProfilePicture string    `gorm:"size:255" json:"profile_picture"`
	BannerImage    string    `gorm:"size:255" json:"banner_image"`
	TotalFollowers int       `gorm:"default:0" json:"total_followers"`
	TotalFollowing int       `gorm:"default:0" json:"total_following"`
	TotalPosts     int       `gorm:"default:0" json:"total_posts"`
	TotalLikes     int       `gorm:"default:0" json:"total_likes"`
	CreatedAt      time.Time `json:"created_at"`

	User User `gorm:"foreignKey:UserID" json:"user"`
}

type Follow struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	FollowerID  uint      `gorm:"not null" json:"follower_id"`
	FollowingID uint      `gorm:"not null" json:"following_id"`
	CreatedAt   time.Time `json:"created_at"`

	Follower  User `gorm:"foreignKey:FollowerID" json:"follower"`
	Following User `gorm:"foreignKey:FollowingID" json:"following"`
}

type Post struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	UserID        uint      `gorm:"not null" json:"user_id"`
	Content       string    `gorm:"type:text" json:"content"`
	Visibility    string    `gorm:"type:enum('public','private','followers');default:'public'" json:"visibility"`
	TotalLikes    int       `gorm:"default:0" json:"total_likes"`
	TotalComments int       `gorm:"default:0" json:"total_comments"`
	TotalShares   int       `gorm:"default:0" json:"total_shares"`
	CreatedAt     time.Time `json:"created_at"`

	User          User          `gorm:"foreignKey:UserID" json:"user"`
	PostMedia     []PostMedia   `gorm:"foreignKey:PostID" json:"media"`
	PostComments  []PostComment `gorm:"foreignKey:PostID" json:"comments"`
}

type PostMedia struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PostID    uint      `gorm:"not null" json:"post_id"`
	MediaUrl  string    `gorm:"size:255;not null" json:"media_url"`
	MediaType string    `gorm:"type:enum('image','video')" json:"media_type"`
	CreatedAt time.Time `json:"created_at"`
}

type PostLike struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PostID    uint      `gorm:"not null" json:"post_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`

	Post Post `gorm:"foreignKey:PostID" json:"post"`
	User User `gorm:"foreignKey:UserID" json:"user"`
}

type PostComment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PostID    uint      `gorm:"not null" json:"post_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Comment   string    `gorm:"type:text;not null" json:"comment"`
	ParentID  *uint     `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`

	UserUser  User         `gorm:"foreignKey:UserID" json:"user"`
	Parent    *PostComment `gorm:"foreignKey:ParentID" json:"parent"`
}

type PostShare struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PostID    uint      `gorm:"not null" json:"post_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Caption   string    `gorm:"type:text" json:"caption"`
	CreatedAt time.Time `json:"created_at"`

	Post Post `gorm:"foreignKey:PostID" json:"post"`
	User User `gorm:"foreignKey:UserID" json:"user"`
}

type PostView struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	PostID   uint      `json:"post_id"`
	UserID   uint      `json:"user_id"`
	ViewedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"viewed_at"`
}

type Conversation struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100" json:"name"`
	IsGroup   bool      `gorm:"default:false" json:"is_group"`
	CreatedBy *uint     `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`

	Creator User `gorm:"foreignKey:CreatedBy" json:"creator"`
}

type ConversationMember struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	ConversationID uint      `gorm:"not null" json:"conversation_id"`
	UserID         uint      `gorm:"not null" json:"user_id"`
	Role           string    `gorm:"type:enum('admin','member');default:'member'" json:"role"`
	JoinedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"joined_at"`
}

type Message struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	ConversationID uint      `gorm:"not null" json:"conversation_id"`
	SenderID       uint      `gorm:"not null" json:"sender_id"`
	Message        string    `gorm:"type:text" json:"message"`
	MessageType    string    `gorm:"type:enum('text','image','video','file');default:'text'" json:"message_type"`
	CreatedAt      time.Time `json:"created_at"`

	Sender User           `gorm:"foreignKey:SenderID" json:"sender"`
	Media  []MessageMedia `gorm:"foreignKey:MessageID" json:"media"`
}

type MessageMedia struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	MessageID uint   `gorm:"not null" json:"message_id"`
	MediaUrl  string `gorm:"size:255" json:"media_url"`
	MediaType string `gorm:"type:enum('image','video','file')" json:"media_type"`
}

type MessageRead struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	MessageID uint      `json:"message_id"`
	UserID    uint      `json:"user_id"`
	ReadAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"read_at"`
}
