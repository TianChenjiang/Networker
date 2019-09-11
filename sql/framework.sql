-- MySQL dump 10.13  Distrib 5.7.21, for Linux (x86_64)
--
-- Host: localhost    Database: citicup
-- ------------------------------------------------------
-- Server version	5.7.21

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `company`
--

DROP TABLE IF EXISTS `company`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `company` (
  `id` bigint(20) DEFAULT NULL,
  `ts_code` text COLLATE utf8mb4_unicode_ci,
  `sub_code` text COLLATE utf8mb4_unicode_ci,
  `name` text COLLATE utf8mb4_unicode_ci,
  `ipo_date` text COLLATE utf8mb4_unicode_ci,
  `issue_date` text COLLATE utf8mb4_unicode_ci,
  `amount` double DEFAULT NULL,
  `market_amount` double DEFAULT NULL,
  `price` double DEFAULT NULL,
  `pe` double DEFAULT NULL,
  `limit_amount` double DEFAULT NULL,
  `funds` double DEFAULT NULL,
  `ballot` double DEFAULT NULL,
  KEY `ix_company_id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `company`
--

LOCK TABLES `company` WRITE;
/*!40000 ALTER TABLE `company` DISABLE KEYS */;
INSERT INTO `company` VALUES (0,'002939.SZ','002939','长城证券','20181017','20181026',31034,27931,6.31,22.98,9.3,19.582,0.16),(1,'002940.SZ','002940','昂利康','20181011','20181023',2250,2025,23.07,22.99,0.9,5.191,0.03),(2,'601162.SH','780162','天风证券','20181009','20181019',51800,46620,1.79,22.86,15.5,0,0.25),(3,'300760.SZ','300760','迈瑞医疗','20180927','20181016',12160,10944,48.8,22.99,3.6,59.341,0.08),(4,'300694.SZ','300694','蠡湖股份','20180927','20181015',5383,4845,9.89,22.98,2.15,5.324,0.04),(5,'300749.SZ','300749','顶固集创','20180913','20180925',2850,2565,12.22,22.99,1.1,3.483,0.03),(6,'002937.SZ','002937','兴瑞科技','20180912','20180926',4600,4140,9.94,22.99,1.8,4.572,0.04),(7,'601577.SH','780577','长沙银行','20180912','20180926',34216,30794,7.99,6.97,10.2,27.338,0.17),(8,'603583.SH','732583','捷昌驱动','20180911','20180921',3020,2718,29.17,22.99,1.2,8.809,0.03),(9,'002936.SZ','002936','郑州银行','20180907','20180919',60000,54000,4.59,6.5,18,27.54,0.25),(10,'603810.SH','732810','丰山集团','20180906','20180917',2000,2000,25.43,20.39,2,5.086,0.02),(11,'300748.SZ','300748','金力永磁','20180906','20180921',4160,3744,5.39,22.98,1.2,2.242,0.05),(12,'002938.SZ','002938','鹏鼎控股','20180905','20180918',23114,20803,16.07,22.99,6.9,37.145,0.12);
/*!40000 ALTER TABLE `company` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-09-11 12:22:14
