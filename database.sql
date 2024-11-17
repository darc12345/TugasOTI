-- MySQL dump 10.13  Distrib 8.0.40, for Linux (x86_64)
--
-- Host: localhost    Database: projectdb
-- ------------------------------------------------------
-- Server version	8.0.40

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
-- Table structure for table `OrderDetails`
--

DROP TABLE IF EXISTS `OrderDetails`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `OrderDetails` (
  `ORDERID` varchar(255) NOT NULL,
  `PRODUCTID` varchar(255) NOT NULL,
  `QUANTITY` int unsigned NOT NULL,
  `PRICE` bigint unsigned DEFAULT NULL,
  `TOTAL` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`ORDERID`,`PRODUCTID`),
  KEY `PRODUCTID` (`PRODUCTID`),
  CONSTRAINT `OrderDetails_ibfk_1` FOREIGN KEY (`PRODUCTID`) REFERENCES `Products` (`PRODUCTID`),
  CONSTRAINT `OrderDetails_ibfk_2` FOREIGN KEY (`ORDERID`) REFERENCES `Orders` (`ORDERID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `OrderDetails`
--

LOCK TABLES `OrderDetails` WRITE;
/*!40000 ALTER TABLE `OrderDetails` DISABLE KEYS */;
/*!40000 ALTER TABLE `OrderDetails` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Orders`
--

DROP TABLE IF EXISTS `Orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Orders` (
  `USERID` varchar(255) NOT NULL,
  `ORDERID` varchar(255) NOT NULL,
  PRIMARY KEY (`ORDERID`),
  KEY `USERID` (`USERID`),
  CONSTRAINT `Orders_ibfk_1` FOREIGN KEY (`USERID`) REFERENCES `Users` (`USERID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Orders`
--

LOCK TABLES `Orders` WRITE;
/*!40000 ALTER TABLE `Orders` DISABLE KEYS */;
/*!40000 ALTER TABLE `Orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Products`
--

DROP TABLE IF EXISTS `Products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Products` (
  `PRODUCTID` varchar(255) NOT NULL,
  `PRODUCTDESC` varchar(255) DEFAULT NULL,
  `PRODUCTNAME` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`PRODUCTID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Products`
--

LOCK TABLES `Products` WRITE;
/*!40000 ALTER TABLE `Products` DISABLE KEYS */;
INSERT INTO `Products` VALUES ('277f0b3d-7a5c-4dec-86e6-1beae7f7b15f','This is a good product','adudugoat'),('5a1428af-8992-4537-80f8-9f9150bd3061','This is a good product','adudugoat'),('7ecd7ff6-05b7-46a6-8bbd-03777874c605','This is an updated product','Yagami'),('c011db5c-1b25-4adf-9f9b-14137a155ae9','This is a good product','adudugoat');
/*!40000 ALTER TABLE `Products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Users`
--

DROP TABLE IF EXISTS `Users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Users` (
  `USERID` varchar(255) NOT NULL,
  `EMAIL` varchar(255) NOT NULL,
  `ADDRESS` varchar(255) NOT NULL,
  `PASSWORDHASH` varchar(255) NOT NULL,
  `ROLE` varchar(16) NOT NULL DEFAULT 'USER',
  PRIMARY KEY (`USERID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Users`
--

LOCK TABLES `Users` WRITE;
/*!40000 ALTER TABLE `Users` DISABLE KEYS */;
INSERT INTO `Users` VALUES ('00b267a7-809a-4c78-9cae-6a9710b1dbf6','','','47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=','USER'),('0557354a-e91b-48bd-9665-538d6a7680a4','johndoe@mail.ugm.ac.id','jkt','10/w7o2juYBrGMh32/KbveULW9jk2tejpyUAD+uC6PE=','USER'),('6056b8cc-2d80-45ff-a424-703e9f74cd9e','johndoe@mail.ugm.ac.id','new york','10/w7o2juYBrGMh32/KbveULW9jk2tejpyUAD+uC6PE=','USER'),('84621199-31f6-47a0-944a-43b12f79d2d8','admin@mail.ugm.ac.id','Milky Way','10/w7o2juYBrGMh32/KbveULW9jk2tejpyUAD+uC6PE=','ADMIN'),('f3056a76-d18b-4c0e-8099-de148bcf562f','johndoe@mail.ugm.ac.id','new york','10/w7o2juYBrGMh32/KbveULW9jk2tejpyUAD+uC6PE=','USER'),('f8ca4ac2-2343-4e99-8174-61c100deb0b0','','','47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=','USER');
/*!40000 ALTER TABLE `Users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-11-17 21:27:51
