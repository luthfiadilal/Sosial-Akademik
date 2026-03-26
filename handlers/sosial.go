package handlers

import (
	"akademik-backend/config"
	"akademik-backend/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ---------------------------------------------------------
// POSTS (FEEDS)
// ---------------------------------------------------------

// Get Feed Posts
func GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := config.DB.Preload("User").Preload("PostMedia").Preload("PostComments").Order("created_at desc").Limit(20).Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil timeline feed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

// Create Post
func CreatePost(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var input struct {
		Content    string `json:"content" binding:"required"`
		Visibility string `json:"visibility"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{
		UserID:     uint(userID.(float64)),
		Content:    input.Content,
		Visibility: input.Visibility,
	}
	
	if post.Visibility == "" {
		post.Visibility = "public"
	}

	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat status"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": post})
}

// Delete Post
func DeletePost(c *gin.Context) {
	postID := c.Param("id")
	userID, _ := c.Get("user_id")

	var post models.Post
	if err := config.DB.Where("id = ?", postID).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if post.UserID != uint(userID.(float64)) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to delete this post"})
		return
	}

	config.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

// ---------------------------------------------------------
// INTERACTIONS (LIKE, COMMENT, SHARE, FOLLOW)
// ---------------------------------------------------------

// Like Post
func LikePost(c *gin.Context) {
	postIDStr := c.Param("id")
	postID, _ := strconv.ParseUint(postIDStr, 10, 32)
	userID, _ := c.Get("user_id")
	actionUserID := uint(userID.(float64))

	var existingLike models.PostLike
	if err := config.DB.Where("post_id = ? AND user_id = ?", postID, actionUserID).First(&existingLike).Error; err == nil {
		// Unlike
		config.DB.Delete(&existingLike)
		config.DB.Model(&models.Post{}).Where("id = ?", postID).UpdateColumn("total_likes", config.DB.Raw("total_likes - 1"))
		c.JSON(http.StatusOK, gin.H{"message": "Post unliked"})
		return
	}

	var post models.Post
	if err := config.DB.Where("id = ?", postID).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	like := models.PostLike{PostID: uint(postID), UserID: actionUserID}
	
	post.TotalLikes += 1
	config.DB.Save(&post)
	config.DB.Create(&like)

	// WEBSOCKET NOTIFICATION
	if post.UserID != actionUserID {
		notif := models.Notification{
			UserID:  post.UserID,
			Title:   "New Like",
			Message: "Seseorang menyukai postingan Anda!",
		}
		config.DB.Create(&notif)

		if AppHub != nil {
			payload, _ := json.Marshal(map[string]interface{}{
				"type": "new_notification",
				"data": notif,
			})
			AppHub.SendToUser(post.UserID, payload)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post liked"})
}

// Comment On Post
func CommentOnPost(c *gin.Context) {
	postIDStr := c.Param("id")
	postID, _ := strconv.ParseUint(postIDStr, 10, 32)
	userID, _ := c.Get("user_id")
	actionUserID := uint(userID.(float64))

	var input struct {
		Comment string `json:"comment" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var post models.Post
	if err := config.DB.Where("id = ?", postID).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	comment := models.PostComment{
		PostID:  uint(postID),
		UserID:  actionUserID,
		Comment: input.Comment,
	}
	config.DB.Create(&comment)

	// Increment comments count
	post.TotalComments += 1
	config.DB.Save(&post)

	// WEBSOCKET NOTIFICATION
	if post.UserID != actionUserID {
		notif := models.Notification{
			UserID:  post.UserID,
			Title:   "New Comment",
			Message: "Ada komentar baru di postingan Anda!",
		}
		config.DB.Create(&notif)
		if AppHub != nil {
			payload, _ := json.Marshal(map[string]interface{}{"type": "new_notification", "data": notif})
			AppHub.SendToUser(post.UserID, payload)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": comment})
}

// Get Post Comments
func GetPostComments(c *gin.Context) {
	postID := c.Param("id")
	var comments []models.PostComment
	if err := config.DB.Preload("UserUser").Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil komentar"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": comments})
}

// Share Post
func SharePost(c *gin.Context) {
	postIDStr := c.Param("id")
	postID, _ := strconv.ParseUint(postIDStr, 10, 32)
	userID, _ := c.Get("user_id")

	var input struct {
		Caption string `json:"caption"`
	}
	c.ShouldBindJSON(&input)

	share := models.PostShare{
		PostID:  uint(postID),
		UserID:  uint(userID.(float64)),
		Caption: input.Caption,
	}

	if err := config.DB.Create(&share).Error; err == nil {
		config.DB.Model(&models.Post{}).Where("id = ?", postID).UpdateColumn("total_shares", config.DB.Raw("total_shares + 1"))
		c.JSON(http.StatusCreated, gin.H{"message": "Post shared", "data": share})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to share post"})
	}
}

// Follow User
func FollowUser(c *gin.Context) {
	targetIDStr := c.Param("id")
	targetID, _ := strconv.ParseUint(targetIDStr, 10, 32)
	followerIDInterface, _ := c.Get("user_id")
	followerID := uint(followerIDInterface.(float64))

	if uint(targetID) == followerID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot follow yourself"})
		return
	}

	var existingFollow models.Follow
	if err := config.DB.Where("follower_id = ? AND following_id = ?", followerID, targetID).First(&existingFollow).Error; err == nil {
		// Unfollow logic
		config.DB.Delete(&existingFollow)
		c.JSON(http.StatusOK, gin.H{"message": "Successfully unfollowed user"})
		return
	}

	follow := models.Follow{
		FollowerID:  followerID,
		FollowingID: uint(targetID),
	}
	config.DB.Create(&follow)

	// WEBSOCKET NOTIFICATION
	notif := models.Notification{
		UserID:  uint(targetID),
		Title:   "New Follower",
		Message: "Seseorang mulai mengikuti Anda!",
	}
	config.DB.Create(&notif)
	if AppHub != nil {
		payload, _ := json.Marshal(map[string]interface{}{"type": "new_notification", "data": notif})
		AppHub.SendToUser(uint(targetID), payload)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully followed user"})
}

// ---------------------------------------------------------
// CHATS AND MESSAGES
// ---------------------------------------------------------

// Get Conversations (Chat List)
func GetConversations(c *gin.Context) {
	userID, _ := c.Get("user_id")
	
	var members []models.ConversationMember
	if err := config.DB.Where("user_id = ?", userID).Find(&members).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil list pesan"})
		return
	}

	var conversationIDs []uint
	for _, m := range members {
		conversationIDs = append(conversationIDs, m.ConversationID)
	}

	var conversations []models.Conversation
	if len(conversationIDs) > 0 {
		config.DB.Where("id IN ?", conversationIDs).Find(&conversations)
	}
	
	c.JSON(http.StatusOK, gin.H{"data": conversations})
}

// Create Conversation (Start a 1v1 Chat or Group)
func CreateConversation(c *gin.Context) {
	userIDInterface, _ := c.Get("user_id")
	creatorID := uint(userIDInterface.(float64))

	var input struct {
		Name          string `json:"name"`
		IsGroup       bool   `json:"is_group"`
		ParticipantID uint   `json:"participant_id"` // For 1v1 chat
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tCreator := creatorID // Assigning creatorID proxy variable to keep original var type safety logic in pointers
	conv := models.Conversation{
		Name:      input.Name,
		IsGroup:   input.IsGroup,
		CreatedBy: &tCreator,
	}
	config.DB.Create(&conv)

	// Add Creator
	config.DB.Create(&models.ConversationMember{ConversationID: conv.ID, UserID: creatorID, Role: "admin"})

	// Add Participant if 1v1
	if !input.IsGroup && input.ParticipantID != 0 {
		config.DB.Create(&models.ConversationMember{ConversationID: conv.ID, UserID: input.ParticipantID, Role: "member"})
	}

	c.JSON(http.StatusCreated, gin.H{"data": conv})
}

// Get Conversation Messages
func GetConversationMessages(c *gin.Context) {
	convID := c.Param("id")

	var messages []models.Message
	if err := config.DB.Preload("Sender").Where("conversation_id = ?", convID).Order("created_at asc").Limit(100).Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch messages"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data": messages})
}

// Send Message
func SendMessage(c *gin.Context) {
	convIDStr := c.Param("id")
	convID, _ := strconv.ParseUint(convIDStr, 10, 32)
	userID, _ := c.Get("user_id")
	senderID := uint(userID.(float64))

	var input struct {
		Message string `json:"message" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg := models.Message{
		ConversationID: uint(convID),
		SenderID:       senderID,
		Message:        input.Message,
		MessageType:    "text",
	}

	if err := config.DB.Create(&msg).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save message"})
		return
	}

	var members []models.ConversationMember
	config.DB.Where("conversation_id = ?", convID).Find(&members)

	payload, _ := json.Marshal(map[string]interface{}{
		"type": "new_message",
		"data": msg,
	})
	
	for _, m := range members {
		if m.UserID != senderID && AppHub != nil {
			AppHub.SendToUser(m.UserID, payload)
		}
	}

	c.JSON(http.StatusCreated, gin.H{"data": msg})
}

// ---------------------------------------------------------
// PROFILES
// ---------------------------------------------------------

// Get Social Profile
func GetSocialProfile(c *gin.Context) {
	username := c.Param("username")
	var profile models.UserProfile
	if err := config.DB.Preload("User").Where("username = ?", username).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}
	
	// Get User's Posts
	var posts []models.Post
	config.DB.Preload("User").Preload("PostMedia").Preload("PostComments").Where("user_id = ?", profile.UserID).Order("created_at desc").Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"profile": profile,
		"posts":   posts,
	})
}

// Update Social Profile
func UpdateSocialProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var input struct {
		Username       string `json:"username"`
		Bio            string `json:"bio"`
		ProfilePicture string `json:"profile_picture"`
		BannerImage    string `json:"banner_image"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var profile models.UserProfile
	if err := config.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		// Create if not exists
		profile = models.UserProfile{
			UserID:         uint(userID.(float64)),
			Username:       input.Username,
			Bio:            input.Bio,
			ProfilePicture: input.ProfilePicture,
			BannerImage:    input.BannerImage,
		}
		config.DB.Create(&profile)
	} else {
		if input.Username != "" {
			profile.Username = input.Username
		}
		profile.Bio = input.Bio
		profile.ProfilePicture = input.ProfilePicture
		profile.BannerImage = input.BannerImage
		config.DB.Save(&profile)
	}

	c.JSON(http.StatusOK, gin.H{"data": profile})
}

// Delete Comment
func DeleteComment(c *gin.Context) {
	commentID := c.Param("id")
	userID, _ := c.Get("user_id")

	var comment models.PostComment
	if err := config.DB.Where("id = ?", commentID).First(&comment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	if comment.UserID != uint(userID.(float64)) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to delete this comment"})
		return
	}

	config.DB.Model(&models.Post{}).Where("id = ?", comment.PostID).UpdateColumn("total_comments", config.DB.Raw("total_comments - 1"))
	config.DB.Delete(&comment)

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
