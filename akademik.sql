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

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
