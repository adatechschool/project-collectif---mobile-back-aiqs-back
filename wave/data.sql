mysqldump: [Warning] Using a password on the command line interface can be insecure.
-- MySQL dump 10.13  Distrib 8.0.29, for Linux (x86_64)
--
-- Host: docker.for.mac.localhost    Database: surfdatabase
-- ------------------------------------------------------
-- Server version	8.0.29

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `surves`
--

DROP TABLE IF EXISTS `surves`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `surves` (
  `id` int NOT NULL AUTO_INCREMENT,
  `surf_break` longtext COLLATE utf8mb4_unicode_ci,
  `difficulty_level` bigint DEFAULT NULL,
  `destination` longtext COLLATE utf8mb4_unicode_ci,
  `magic_seaweed_link` longtext COLLATE utf8mb4_unicode_ci,
  `photos` longtext COLLATE utf8mb4_unicode_ci,
  `peak_surf_season_begins` longtext COLLATE utf8mb4_unicode_ci,
  `destination_state_country` longtext COLLATE utf8mb4_unicode_ci,
  `peak_surf_season_ends` longtext COLLATE utf8mb4_unicode_ci,
  `address` longtext COLLATE utf8mb4_unicode_ci,
  `lat` double DEFAULT NULL,
  `lng` double DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_surves_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `surves`
--

LOCK TABLES `surves` WRITE;
/*!40000 ALTER TABLE `surves` DISABLE KEYS */;
INSERT INTO `surves` VALUES (15,'Reef Break',4,'Pipeline','https://magicseaweed.com/Pipeline-Backdoor-Surf-Report/616/','https://dl.airtable.com/ZuXJZ2NnTF40kCdBfTld_thomas-ashlock-64485-unsplash.jpg?ts=1653308030&userId=usrVSPQAdslUijuds&cs=5624cbe57ad2986c','2018-07-22','Oahu, Hawaii','2018-08-31','Pipeline, Oahu, Hawaii',21.6650562,-158.05120469999997,NULL,NULL,NULL),(16,'Point Break',5,'Supertubes','https://magicseaweed.com/Jeffreys-Bay-J-Bay-Surf-Report/88/','https://dl.airtable.com/e3QoP3cFSyykZJOvWGIy_cesar-couto-477018-unsplash%20(1).jpg?ts=1653308030&userId=usrVSPQAdslUijuds&cs=8de6a520818bb80d','2018-08-01','Jeffreys Bay, South Africa','2018-10-09','Supertubes, Jeffreys Bay, South Africa',-34.031783,24.93159400000002,NULL,NULL,NULL),(17,'Point Break',2,'Manu Bay','https://magicseaweed.com/Raglan-Surf-Report/91/','https://dl.airtable.com/jF7Y6RpGQueQok9TRQ9I_holger-link-707884-unsplash.jpg?ts=1653308030&userId=usrVSPQAdslUijuds&cs=7903db9ad8314844','2018-12-01','Raglan, New Zealand','2019-01-31','Manu Bay, Raglan, New Zealand',-37.8214592,174.81221619999997,NULL,NULL,NULL),(18,'Point Break',4,'Superbank','https://magicseaweed.com/Surfers-Paradise-Gold-Coast-Surf-Report/1012/','https://dl.airtable.com/I4E4xZeQbO2g814udQDm_jeremy-bishop-80371-unsplash.jpg?ts=1653308030&userId=usrVSPQAdslUijuds&cs=fd002047a0be2ce7','2018-11-28','Gold Coast, Australia','2019-02-01','Superbank, Gold Coast, Australia',-28.162557,153.55002880000006,NULL,NULL,NULL),(19,'Outer Banks',3,'Kitty Hawk','https://magicseaweed.com/North-Carolina-Outer-Banks-North-Surfing/283/','https://dl.airtable.com/c9xSCxlS8a6YITIYnQjE_jens-theess-511-unsplash.jpg?ts=1653308030&userId=usrVSPQAdslUijuds&cs=7bf70e5165c8c631','2018-08-09','Outer Banks, North Carolina','2018-10-18','Kitty Hawk, Outer Banks, North Carolina',36.1005777,-75.71195569999998,NULL,NULL,NULL);
/*!40000 ALTER TABLE `surves` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `user_id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext COLLATE utf8mb4_unicode_ci,
  `email` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `password` longtext COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (3,'TEST2','quentinyuehyulee@gmail.com','$2a$10$GTLLLHpLGaIdhVO5v8QZFenfzdH54GKWBuoM8ZBKEJcVics41UNHq'),(4,'Quentin','quentinyuehyulee@gmail.com','$2a$10$ioTmoQNQO/HBtQd25dpIjO/Jt2x2Hu0/lJ8Xiz03BTEOz/w6G05ai'),(5,'Isciane','quentinyuehyulee@gmail.com','$2a$10$gsUcWUScizDAXk24H7mwUuqwLpG/YJ.sh7SG1i4KqEAL7Loz8tZUC');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-05-24 14:04:42
