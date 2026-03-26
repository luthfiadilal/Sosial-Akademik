-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               8.0.30 - MySQL Community Server - GPL
-- Server OS:                    Win64
-- HeidiSQL Version:             12.1.0.6537
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- Dumping structure for table akademik.assignments
CREATE TABLE IF NOT EXISTS `assignments` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `class_id` bigint NOT NULL,
  `title` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` text COLLATE utf8mb4_unicode_ci,
  `due_date` datetime DEFAULT NULL,
  `max_score` decimal(5,2) DEFAULT '100.00',
  `created_by` bigint DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_assignment_class` (`class_id`),
  KEY `fk_assignment_lecturer` (`created_by`),
  CONSTRAINT `fk_assignment_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_assignment_lecturer` FOREIGN KEY (`created_by`) REFERENCES `lecturers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.assignments: ~0 rows (approximately)

-- Dumping structure for table akademik.assignment_submissions
CREATE TABLE IF NOT EXISTS `assignment_submissions` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `assignment_id` bigint NOT NULL,
  `student_id` bigint NOT NULL,
  `file_url` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `note` text COLLATE utf8mb4_unicode_ci,
  `submitted_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `score` decimal(5,2) DEFAULT NULL,
  `feedback` text COLLATE utf8mb4_unicode_ci,
  `graded_by` bigint DEFAULT NULL,
  `graded_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_submission_assignment` (`assignment_id`),
  KEY `fk_submission_student` (`student_id`),
  KEY `fk_submission_lecturer` (`graded_by`),
  CONSTRAINT `fk_submission_assignment` FOREIGN KEY (`assignment_id`) REFERENCES `assignments` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_submission_lecturer` FOREIGN KEY (`graded_by`) REFERENCES `lecturers` (`id`),
  CONSTRAINT `fk_submission_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.assignment_submissions: ~0 rows (approximately)

-- Dumping structure for table akademik.attendances
CREATE TABLE IF NOT EXISTS `attendances` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `student_id` bigint NOT NULL,
  `class_id` bigint NOT NULL,
  `meeting_number` int NOT NULL,
  `date` date DEFAULT NULL,
  `status` enum('hadir','izin','sakit','alfa') COLLATE utf8mb4_unicode_ci DEFAULT 'hadir',
  PRIMARY KEY (`id`),
  KEY `fk_att_student` (`student_id`),
  KEY `fk_att_class` (`class_id`),
  CONSTRAINT `fk_att_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`),
  CONSTRAINT `fk_att_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.attendances: ~0 rows (approximately)

-- Dumping structure for table akademik.billings
CREATE TABLE IF NOT EXISTS `billings` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `student_id` bigint NOT NULL,
  `semester` int NOT NULL,
  `academic_year` varchar(9) COLLATE utf8mb4_unicode_ci NOT NULL,
  `amount` decimal(12,2) NOT NULL,
  `status` enum('unpaid','partial','paid') COLLATE utf8mb4_unicode_ci DEFAULT 'unpaid',
  `due_date` date DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_bill_student` (`student_id`),
  CONSTRAINT `fk_bill_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.billings: ~0 rows (approximately)

-- Dumping structure for table akademik.classes
CREATE TABLE IF NOT EXISTS `classes` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `course_id` bigint NOT NULL,
  `class_name` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `capacity` int DEFAULT '30',
  `semester` int NOT NULL,
  `academic_year` varchar(9) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_class_course` (`course_id`),
  CONSTRAINT `fk_class_course` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.classes: ~0 rows (approximately)

-- Dumping structure for table akademik.class_lecturers
CREATE TABLE IF NOT EXISTS `class_lecturers` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `class_id` bigint NOT NULL,
  `lecturer_id` bigint NOT NULL,
  `role` enum('utama','asisten') COLLATE utf8mb4_unicode_ci DEFAULT 'utama',
  PRIMARY KEY (`id`),
  KEY `fk_cl_class` (`class_id`),
  KEY `fk_cl_lecturer` (`lecturer_id`),
  CONSTRAINT `fk_cl_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_cl_lecturer` FOREIGN KEY (`lecturer_id`) REFERENCES `lecturers` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.class_lecturers: ~0 rows (approximately)

-- Dumping structure for table akademik.conversations
CREATE TABLE IF NOT EXISTS `conversations` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `is_group` tinyint(1) DEFAULT '0',
  `created_by` bigint DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_conv_user` (`created_by`),
  CONSTRAINT `fk_conv_user` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.conversations: ~0 rows (approximately)

-- Dumping structure for table akademik.conversation_members
CREATE TABLE IF NOT EXISTS `conversation_members` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `conversation_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `role` enum('admin','member') COLLATE utf8mb4_unicode_ci DEFAULT 'member',
  `joined_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `conversation_id` (`conversation_id`,`user_id`),
  KEY `fk_cm_user` (`user_id`),
  CONSTRAINT `fk_cm_conv` FOREIGN KEY (`conversation_id`) REFERENCES `conversations` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_cm_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.conversation_members: ~0 rows (approximately)

-- Dumping structure for table akademik.courses
CREATE TABLE IF NOT EXISTS `courses` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `code` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL,
  `sks` int NOT NULL,
  `semester` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.courses: ~0 rows (approximately)

-- Dumping structure for table akademik.course_materials
CREATE TABLE IF NOT EXISTS `course_materials` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `class_id` bigint NOT NULL,
  `title` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` text COLLATE utf8mb4_unicode_ci,
  `file_url` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `material_type` enum('file','link','video') COLLATE utf8mb4_unicode_ci DEFAULT 'file',
  `created_by` bigint DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_material_class` (`class_id`),
  KEY `fk_material_lecturer` (`created_by`),
  CONSTRAINT `fk_material_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_material_lecturer` FOREIGN KEY (`created_by`) REFERENCES `lecturers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.course_materials: ~0 rows (approximately)

-- Dumping structure for table akademik.curriculums
CREATE TABLE IF NOT EXISTS `curriculums` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `study_program_id` bigint NOT NULL,
  `year` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `is_active` tinyint(1) DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_curriculum_prodi` (`study_program_id`),
  CONSTRAINT `fk_curriculum_prodi` FOREIGN KEY (`study_program_id`) REFERENCES `study_programs` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.curriculums: ~0 rows (approximately)

-- Dumping structure for table akademik.curriculum_details
CREATE TABLE IF NOT EXISTS `curriculum_details` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `curriculum_id` bigint NOT NULL,
  `course_id` bigint NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_cd_curriculum` (`curriculum_id`),
  KEY `fk_cd_course` (`course_id`),
  CONSTRAINT `fk_cd_course` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_cd_curriculum` FOREIGN KEY (`curriculum_id`) REFERENCES `curriculums` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.curriculum_details: ~0 rows (approximately)

-- Dumping structure for table akademik.faculties
CREATE TABLE IF NOT EXISTS `faculties` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `code` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.faculties: ~0 rows (approximately)

-- Dumping structure for table akademik.follows
CREATE TABLE IF NOT EXISTS `follows` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `follower_id` bigint NOT NULL,
  `following_id` bigint NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `follower_id` (`follower_id`,`following_id`),
  KEY `fk_follow_following` (`following_id`),
  CONSTRAINT `fk_follow_follower` FOREIGN KEY (`follower_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_follow_following` FOREIGN KEY (`following_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.follows: ~0 rows (approximately)

-- Dumping structure for table akademik.forums
CREATE TABLE IF NOT EXISTS `forums` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `class_id` bigint NOT NULL,
  `title` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_by` bigint DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_forum_class` (`class_id`),
  KEY `fk_forum_user` (`created_by`),
  CONSTRAINT `fk_forum_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_forum_user` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.forums: ~0 rows (approximately)

-- Dumping structure for table akademik.forum_comments
CREATE TABLE IF NOT EXISTS `forum_comments` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `thread_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `comment` text COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_comment_thread` (`thread_id`),
  KEY `fk_comment_user` (`user_id`),
  CONSTRAINT `fk_comment_thread` FOREIGN KEY (`thread_id`) REFERENCES `forum_threads` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_comment_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.forum_comments: ~0 rows (approximately)

-- Dumping structure for table akademik.forum_threads
CREATE TABLE IF NOT EXISTS `forum_threads` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `forum_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `content` text COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_thread_forum` (`forum_id`),
  KEY `fk_thread_user` (`user_id`),
  CONSTRAINT `fk_thread_forum` FOREIGN KEY (`forum_id`) REFERENCES `forums` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_thread_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.forum_threads: ~0 rows (approximately)

-- Dumping structure for table akademik.grades
CREATE TABLE IF NOT EXISTS `grades` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `student_id` bigint NOT NULL,
  `class_id` bigint NOT NULL,
  `nilai_angka` decimal(5,2) DEFAULT NULL,
  `nilai_huruf` varchar(2) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `input_by` bigint DEFAULT NULL,
  `input_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_grade_student` (`student_id`),
  KEY `fk_grade_class` (`class_id`),
  KEY `fk_grade_lecturer` (`input_by`),
  CONSTRAINT `fk_grade_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`),
  CONSTRAINT `fk_grade_lecturer` FOREIGN KEY (`input_by`) REFERENCES `lecturers` (`id`),
  CONSTRAINT `fk_grade_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.grades: ~0 rows (approximately)

-- Dumping structure for table akademik.krs
CREATE TABLE IF NOT EXISTS `krs` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `student_id` bigint NOT NULL,
  `semester` int NOT NULL,
  `academic_year` varchar(9) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` enum('draft','diajukan','disetujui','ditolak') COLLATE utf8mb4_unicode_ci DEFAULT 'draft',
  `approved_by` bigint DEFAULT NULL,
  `approved_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_krs_student` (`student_id`),
  KEY `fk_krs_approver` (`approved_by`),
  CONSTRAINT `fk_krs_approver` FOREIGN KEY (`approved_by`) REFERENCES `lecturers` (`id`),
  CONSTRAINT `fk_krs_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.krs: ~0 rows (approximately)

-- Dumping structure for table akademik.krs_approval_logs
CREATE TABLE IF NOT EXISTS `krs_approval_logs` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `krs_id` bigint NOT NULL,
  `lecturer_id` bigint NOT NULL,
  `status` enum('diajukan','disetujui','ditolak') COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `note` text COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_krslog_krs` (`krs_id`),
  KEY `fk_krslog_lecturer` (`lecturer_id`),
  CONSTRAINT `fk_krslog_krs` FOREIGN KEY (`krs_id`) REFERENCES `krs` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_krslog_lecturer` FOREIGN KEY (`lecturer_id`) REFERENCES `lecturers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.krs_approval_logs: ~0 rows (approximately)

-- Dumping structure for table akademik.krs_details
CREATE TABLE IF NOT EXISTS `krs_details` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `krs_id` bigint NOT NULL,
  `class_id` bigint NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_krsd_krs` (`krs_id`),
  KEY `fk_krsd_class` (`class_id`),
  CONSTRAINT `fk_krsd_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`),
  CONSTRAINT `fk_krsd_krs` FOREIGN KEY (`krs_id`) REFERENCES `krs` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.krs_details: ~0 rows (approximately)

-- Dumping structure for table akademik.lecturers
CREATE TABLE IF NOT EXISTS `lecturers` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `nidn` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `study_program_id` bigint DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `nidn` (`nidn`),
  KEY `fk_lecturer_user` (`user_id`),
  KEY `fk_lecturer_prodi` (`study_program_id`),
  CONSTRAINT `fk_lecturer_prodi` FOREIGN KEY (`study_program_id`) REFERENCES `study_programs` (`id`) ON DELETE SET NULL,
  CONSTRAINT `fk_lecturer_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.lecturers: ~0 rows (approximately)

-- Dumping structure for table akademik.messages
CREATE TABLE IF NOT EXISTS `messages` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `conversation_id` bigint NOT NULL,
  `sender_id` bigint NOT NULL,
  `message` text COLLATE utf8mb4_unicode_ci,
  `message_type` enum('text','image','video','file') COLLATE utf8mb4_unicode_ci DEFAULT 'text',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_msg_conv` (`conversation_id`),
  KEY `fk_msg_user` (`sender_id`),
  CONSTRAINT `fk_msg_conv` FOREIGN KEY (`conversation_id`) REFERENCES `conversations` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_msg_user` FOREIGN KEY (`sender_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.messages: ~0 rows (approximately)

-- Dumping structure for table akademik.message_media
CREATE TABLE IF NOT EXISTS `message_media` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `message_id` bigint NOT NULL,
  `media_url` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `media_type` enum('image','video','file') COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_mm_message` (`message_id`),
  CONSTRAINT `fk_mm_message` FOREIGN KEY (`message_id`) REFERENCES `messages` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.message_media: ~0 rows (approximately)

-- Dumping structure for table akademik.message_reads
CREATE TABLE IF NOT EXISTS `message_reads` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `message_id` bigint DEFAULT NULL,
  `user_id` bigint DEFAULT NULL,
  `read_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `message_id` (`message_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.message_reads: ~0 rows (approximately)

-- Dumping structure for table akademik.notifications
CREATE TABLE IF NOT EXISTS `notifications` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint DEFAULT NULL,
  `title` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `message` text COLLATE utf8mb4_unicode_ci,
  `is_read` tinyint(1) DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_notif_user` (`user_id`),
  CONSTRAINT `fk_notif_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.notifications: ~0 rows (approximately)

-- Dumping structure for table akademik.payments
CREATE TABLE IF NOT EXISTS `payments` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `billing_id` bigint NOT NULL,
  `amount` decimal(12,2) NOT NULL,
  `payment_date` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `method` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `reference_no` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `status` enum('pending','success','failed') COLLATE utf8mb4_unicode_ci DEFAULT 'pending',
  PRIMARY KEY (`id`),
  KEY `fk_payment_billing` (`billing_id`),
  CONSTRAINT `fk_payment_billing` FOREIGN KEY (`billing_id`) REFERENCES `billings` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.payments: ~0 rows (approximately)

-- Dumping structure for table akademik.posts
CREATE TABLE IF NOT EXISTS `posts` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `content` text COLLATE utf8mb4_unicode_ci,
  `visibility` enum('public','private','followers') COLLATE utf8mb4_unicode_ci DEFAULT 'public',
  `total_likes` int DEFAULT '0',
  `total_comments` int DEFAULT '0',
  `total_shares` int DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_post_user` (`user_id`),
  CONSTRAINT `fk_post_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.posts: ~0 rows (approximately)

-- Dumping structure for table akademik.post_comments
CREATE TABLE IF NOT EXISTS `post_comments` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `post_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `comment` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `parent_id` bigint DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_post_comments_post` (`post_id`),
  KEY `fk_post_comments_user` (`user_id`),
  KEY `fk_post_comments_parent` (`parent_id`),
  CONSTRAINT `fk_post_comments_parent` FOREIGN KEY (`parent_id`) REFERENCES `post_comments` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_post_comments_post` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_post_comments_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.post_comments: ~0 rows (approximately)

-- Dumping structure for table akademik.post_likes
CREATE TABLE IF NOT EXISTS `post_likes` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `post_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `post_id` (`post_id`,`user_id`),
  KEY `fk_like_user` (`user_id`),
  CONSTRAINT `fk_like_post` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_like_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.post_likes: ~0 rows (approximately)

-- Dumping structure for table akademik.post_media
CREATE TABLE IF NOT EXISTS `post_media` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `post_id` bigint NOT NULL,
  `media_url` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `media_type` enum('image','video') COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_media_post` (`post_id`),
  CONSTRAINT `fk_media_post` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.post_media: ~0 rows (approximately)

-- Dumping structure for table akademik.post_shares
CREATE TABLE IF NOT EXISTS `post_shares` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `post_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `caption` text COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_share_post` (`post_id`),
  KEY `fk_share_user` (`user_id`),
  CONSTRAINT `fk_share_post` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_share_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.post_shares: ~0 rows (approximately)

-- Dumping structure for table akademik.post_views
CREATE TABLE IF NOT EXISTS `post_views` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `post_id` bigint DEFAULT NULL,
  `user_id` bigint DEFAULT NULL,
  `viewed_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.post_views: ~0 rows (approximately)

-- Dumping structure for table akademik.quizzes
CREATE TABLE IF NOT EXISTS `quizzes` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `class_id` bigint NOT NULL,
  `title` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `description` text COLLATE utf8mb4_unicode_ci,
  `start_time` datetime DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  `duration_minutes` int DEFAULT NULL,
  `created_by` bigint DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_quiz_class` (`class_id`),
  KEY `fk_quiz_lecturer` (`created_by`),
  CONSTRAINT `fk_quiz_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_quiz_lecturer` FOREIGN KEY (`created_by`) REFERENCES `lecturers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.quizzes: ~0 rows (approximately)

-- Dumping structure for table akademik.quiz_answers
CREATE TABLE IF NOT EXISTS `quiz_answers` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `question_id` bigint NOT NULL,
  `student_id` bigint NOT NULL,
  `selected_option_id` bigint DEFAULT NULL,
  `answer_text` text COLLATE utf8mb4_unicode_ci,
  `is_correct` tinyint(1) DEFAULT NULL,
  `score` decimal(5,2) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_answer_question` (`question_id`),
  KEY `fk_answer_student` (`student_id`),
  KEY `fk_answer_option` (`selected_option_id`),
  CONSTRAINT `fk_answer_option` FOREIGN KEY (`selected_option_id`) REFERENCES `quiz_options` (`id`),
  CONSTRAINT `fk_answer_question` FOREIGN KEY (`question_id`) REFERENCES `quiz_questions` (`id`),
  CONSTRAINT `fk_answer_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.quiz_answers: ~0 rows (approximately)

-- Dumping structure for table akademik.quiz_options
CREATE TABLE IF NOT EXISTS `quiz_options` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `question_id` bigint NOT NULL,
  `option_text` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `is_correct` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `fk_option_question` (`question_id`),
  CONSTRAINT `fk_option_question` FOREIGN KEY (`question_id`) REFERENCES `quiz_questions` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.quiz_options: ~0 rows (approximately)

-- Dumping structure for table akademik.quiz_questions
CREATE TABLE IF NOT EXISTS `quiz_questions` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `quiz_id` bigint NOT NULL,
  `question` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `type` enum('multiple_choice','essay') COLLATE utf8mb4_unicode_ci DEFAULT 'multiple_choice',
  PRIMARY KEY (`id`),
  KEY `fk_question_quiz` (`quiz_id`),
  CONSTRAINT `fk_question_quiz` FOREIGN KEY (`quiz_id`) REFERENCES `quizzes` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.quiz_questions: ~0 rows (approximately)

-- Dumping structure for table akademik.quiz_results
CREATE TABLE IF NOT EXISTS `quiz_results` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `quiz_id` bigint NOT NULL,
  `student_id` bigint NOT NULL,
  `total_score` decimal(5,2) DEFAULT NULL,
  `submitted_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_result_quiz` (`quiz_id`),
  KEY `fk_result_student` (`student_id`),
  CONSTRAINT `fk_result_quiz` FOREIGN KEY (`quiz_id`) REFERENCES `quizzes` (`id`),
  CONSTRAINT `fk_result_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.quiz_results: ~0 rows (approximately)

-- Dumping structure for table akademik.schedules
CREATE TABLE IF NOT EXISTS `schedules` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `class_id` bigint NOT NULL,
  `day` enum('senin','selasa','rabu','kamis','jumat','sabtu') COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `start_time` time DEFAULT NULL,
  `end_time` time DEFAULT NULL,
  `room` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_schedule_class` (`class_id`),
  CONSTRAINT `fk_schedule_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.schedules: ~0 rows (approximately)

-- Dumping structure for table akademik.scholarships
CREATE TABLE IF NOT EXISTS `scholarships` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `type` enum('percentage','fixed') COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `value` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.scholarships: ~0 rows (approximately)

-- Dumping structure for table akademik.students
CREATE TABLE IF NOT EXISTS `students` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `npm` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `study_program_id` bigint NOT NULL,
  `angkatan` year NOT NULL,
  `advisor_id` bigint DEFAULT NULL,
  `status` enum('aktif','cuti','lulus','dropout') COLLATE utf8mb4_unicode_ci DEFAULT 'aktif',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `npm` (`npm`),
  KEY `fk_student_user` (`user_id`),
  KEY `fk_student_prodi` (`study_program_id`),
  KEY `fk_student_advisor` (`advisor_id`),
  CONSTRAINT `fk_student_advisor` FOREIGN KEY (`advisor_id`) REFERENCES `lecturers` (`id`) ON DELETE SET NULL,
  CONSTRAINT `fk_student_prodi` FOREIGN KEY (`study_program_id`) REFERENCES `study_programs` (`id`),
  CONSTRAINT `fk_student_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.students: ~0 rows (approximately)

-- Dumping structure for table akademik.student_gpa
CREATE TABLE IF NOT EXISTS `student_gpa` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `student_id` bigint NOT NULL,
  `semester` int DEFAULT NULL,
  `academic_year` varchar(9) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `ips` decimal(3,2) DEFAULT NULL,
  `sks` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_gpa_student` (`student_id`),
  CONSTRAINT `fk_gpa_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.student_gpa: ~0 rows (approximately)

-- Dumping structure for table akademik.student_scholarships
CREATE TABLE IF NOT EXISTS `student_scholarships` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `student_id` bigint DEFAULT NULL,
  `scholarship_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_ss_student` (`student_id`),
  KEY `fk_ss_sch` (`scholarship_id`),
  CONSTRAINT `fk_ss_sch` FOREIGN KEY (`scholarship_id`) REFERENCES `scholarships` (`id`),
  CONSTRAINT `fk_ss_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.student_scholarships: ~0 rows (approximately)

-- Dumping structure for table akademik.student_status_logs
CREATE TABLE IF NOT EXISTS `student_status_logs` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `student_id` bigint DEFAULT NULL,
  `status` enum('aktif','cuti','lulus','dropout') COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `start_date` date DEFAULT NULL,
  `end_date` date DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_status_student` (`student_id`),
  CONSTRAINT `fk_status_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.student_status_logs: ~0 rows (approximately)

-- Dumping structure for table akademik.study_programs
CREATE TABLE IF NOT EXISTS `study_programs` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `faculty_id` bigint NOT NULL,
  `code` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `degree` enum('D3','S1','S2','S3') COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`),
  KEY `fk_prodi_faculty` (`faculty_id`),
  CONSTRAINT `fk_prodi_faculty` FOREIGN KEY (`faculty_id`) REFERENCES `faculties` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.study_programs: ~0 rows (approximately)

-- Dumping structure for table akademik.transcripts
CREATE TABLE IF NOT EXISTS `transcripts` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `student_id` bigint NOT NULL,
  `total_sks` int DEFAULT '0',
  `ipk` decimal(3,2) DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_trans_student` (`student_id`),
  CONSTRAINT `fk_trans_student` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.transcripts: ~0 rows (approximately)

-- Dumping structure for table akademik.transcript_details
CREATE TABLE IF NOT EXISTS `transcript_details` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `transcript_id` bigint NOT NULL,
  `course_id` bigint NOT NULL,
  `class_id` bigint DEFAULT NULL,
  `nilai_huruf` varchar(2) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `nilai_angka` decimal(5,2) DEFAULT NULL,
  `sks` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_td_transcript` (`transcript_id`),
  KEY `fk_td_course` (`course_id`),
  CONSTRAINT `fk_td_course` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`),
  CONSTRAINT `fk_td_transcript` FOREIGN KEY (`transcript_id`) REFERENCES `transcripts` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.transcript_details: ~0 rows (approximately)

-- Dumping structure for table akademik.users
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `role` enum('admin','dosen','mahasiswa') COLLATE utf8mb4_unicode_ci NOT NULL,
  `is_active` tinyint(1) DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.users: ~0 rows (approximately)

-- Dumping structure for table akademik.user_profiles
CREATE TABLE IF NOT EXISTS `user_profiles` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `username` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `bio` text COLLATE utf8mb4_unicode_ci,
  `profile_picture` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `banner_image` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `total_followers` int DEFAULT '0',
  `total_following` int DEFAULT '0',
  `total_posts` int DEFAULT '0',
  `total_likes` int DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `fk_profile_user` (`user_id`),
  CONSTRAINT `fk_profile_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table akademik.user_profiles: ~0 rows (approximately)

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
