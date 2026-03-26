package routes

import (
	"akademik-backend/handlers"
	"akademik-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	
	// Authentication Routes (Public)
	auth := api.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	// Akademik Routes (Currently Public, can be guarded later)
	akademik := api.Group("/akademik")
	{
		akademik.GET("/courses", handlers.GetCourses)
		akademik.GET("/classes", handlers.GetClasses)
	}

	// Protected Routes (Requires JWT Token)
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// Profil Routes
		profile := protected.Group("/profile")
		{
			profile.GET("/", handlers.GetProfile)
			profile.PUT("/", handlers.EditProfile)
		}

		// Krs Routes (Requires Student profile)
		krs := protected.Group("/krs")
		{
			krs.GET("/student/:student_id", handlers.GetKrsByStudent)
			krs.POST("/", handlers.CreateKrs)
		}

		// Admin Master Data & Users Management (Requires JWT and ideally Admin Role checking)
		admin := protected.Group("/admin")
		{
			// User Management
			admin.GET("/users", handlers.GetAllUsers)
			admin.GET("/users/:id", handlers.GetUserByID)
			admin.POST("/users", handlers.CreateUser)
			admin.PUT("/users/:id", handlers.UpdateUser)
			admin.DELETE("/users/:id", handlers.DeleteUser)

			// Course Management
			admin.POST("/courses", handlers.CreateCourse)
			admin.PUT("/courses/:id", handlers.UpdateCourse)
			admin.DELETE("/courses/:id", handlers.DeleteCourse)

			// Class Management
			admin.POST("/classes", handlers.CreateClass)
			admin.PUT("/classes/:id", handlers.UpdateClass)
			admin.DELETE("/classes/:id", handlers.DeleteClass)
		}

		// Keuangan Routes
		keuangan := protected.Group("/keuangan")
		{
			keuangan.GET("/billings/student/:student_id", handlers.GetBillingsByStudent)
			keuangan.GET("/payments/billing/:billing_id", handlers.GetPaymentsByBilling)
		}

		// Administratif Routes
		adminLogs := protected.Group("/administratif")
		{
			adminLogs.GET("/notifications", handlers.GetNotifications)
			adminLogs.PUT("/notifications/:id/read", handlers.MarkNotificationRead)
		}

		// LMS (Learning Management System) Routes
		lms := protected.Group("/lms")
		{
			lms.GET("/class/:class_id/materials", handlers.GetCourseMaterials)
			lms.GET("/class/:class_id/assignments", handlers.GetAssignments)
			lms.GET("/class/:class_id/quizzes", handlers.GetQuizzes)
			lms.GET("/class/:class_id/forums", handlers.GetForums)
		}

		// Sosial Media Routes
		sosial := protected.Group("/sosial")
		{
			// User Profiles
			sosial.GET("/users/:username/profile", handlers.GetSocialProfile)
			sosial.PUT("/users/profile", handlers.UpdateSocialProfile)

			// Posts
			sosial.GET("/posts", handlers.GetPosts)
			sosial.POST("/posts", handlers.CreatePost)
			sosial.DELETE("/posts/:id", handlers.DeletePost)
			
			// Interactions
			sosial.POST("/posts/:id/like", handlers.LikePost)
			sosial.POST("/posts/:id/comment", handlers.CommentOnPost)
			sosial.GET("/posts/:id/comments", handlers.GetPostComments)
			sosial.DELETE("/posts/comments/:id", handlers.DeleteComment)
			sosial.POST("/posts/:id/share", handlers.SharePost)
			
			// Network
			sosial.POST("/users/:id/follow", handlers.FollowUser)
		}

		// Chat & Messaging Routes
		chat := protected.Group("/chat")
		{
			chat.GET("/conversations", handlers.GetConversations)
			chat.POST("/conversations", handlers.CreateConversation)
			chat.GET("/conversations/:id/messages", handlers.GetConversationMessages)
			chat.POST("/conversations/:id/messages", handlers.SendMessage)
		}

		// WebSockets Route
		protected.GET("/ws", handlers.ServeWS)
	}
}
