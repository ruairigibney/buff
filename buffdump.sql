--
-- Dumping data for table `video_stream`
--

LOCK TABLES `video_stream` WRITE;
/*!40000 ALTER TABLE `video_stream` DISABLE KEYS */;
INSERT INTO `video_stream` VALUES (1,'Buff Example Stream 1',1591728525,1591728525),(2,'Buff Example Stream 2',1591732125,1591732125),(3,'Buff Example Stream 3',1591735725,1591735725);
/*!40000 ALTER TABLE `video_stream` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `video_stream_question`
--

LOCK TABLES `video_stream_question` WRITE;
/*!40000 ALTER TABLE `video_stream_question` DISABLE KEYS */;
INSERT INTO `video_stream_question` VALUES (1,1,'Is this an example question?',1),(2,1,'Is this also an example question?',NULL);
/*!40000 ALTER TABLE `video_stream_question` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `video_stream_answer`
--

LOCK TABLES `video_stream_answer` WRITE;
/*!40000 ALTER TABLE `video_stream_answer` DISABLE KEYS */;
INSERT INTO `video_stream_answer` VALUES (1,1,'This is a correct answer'),(2,1,'This is an incorrect answer');
/*!40000 ALTER TABLE `video_stream_answer` ENABLE KEYS */;
UNLOCK TABLES;
