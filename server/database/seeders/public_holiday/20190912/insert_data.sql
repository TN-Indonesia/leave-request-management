-- Dont Input Public Holliday for weekend --
INSERT INTO public_holiday
VALUES
  	(1,'01-01-2019','01-01-2019','New Years Day'),
	(2,'05-02-2019','05-02-2019','Chinese New Year'),
	(3,'07-03-2019','07-03-2019','Bali Hindu New Year'),
	(4,'03-04-2019','03-04-2019','Isra Miraj'),
	(5,'19-04-2019','19-04-2019','Good Friday'),
	(6,'01-05-2019','01-05-2019','Labour Day'),
	(7,'19-05-2019','19-05-2019','Waisak Day'),
	(8,'30-05-2019','30-05-2019','Ascension Day of Jesus Christ'),
	(9,'01-06-2019','01-06-2019','Pancasila Day'),
	(10,'03-06-2019','04-06-2019','Lebaran Holiday'),
	(11,'05-06-2019','06-06-2019','Hari Raya Idul Fitri'),
	(12,'07-06-2019','07-06-2019','Lebaran Holiday'),
	(13,'11-08-2019','11-08-2019','Idul Adha'),
	(14,'17-08-2019','17-08-2019','Independence Day'),
	(15,'01-09-2019','01-09-2019','Islamic New Year'),
	(16,'09-11-2019','09-11-2019','Prophet Muhammads Birthday'),
	(17,'24-12-2019','24-12-2019','Christmas Holiday'),
	(18,'25-12-2019','25-12-2019','Christmas Day'),
	(19,'25-01-2020','25-01-2020','Chinese New Year 2571 Kongzili'),
	(20,'25-03-2020','25-03-2020','Bali Hindu New Year Saka 1941'),
	(21,'10-04-2020','10-04-2020','Wafat Isa Al Masih'),
	(22,'01-05-2020','01-05-2020','Labour Day​'),
	(23,'07-05-2020','07-05-2020','Waisak Day 2564'),
	(24,'21-05-2020','21-05-2020','Ascension Day of Jesus Christ​'),
	(25,'22-05-2020','22-05-2020','Hari Raya Idul Fitri 1441 Hijriyah'),
	(26,'25-05-2020','27-05-2020','Hari Raya Idul Fitri 1441 Hijriyah'),
	(27,'01-06-2020','01-06-2020','Pancasila Day​​'),
	(28,'31-07-2020','31-07-2020','Idul Adha 1441 Hijriyah​'),
	(29,'17-08-2020','17-08-2020','Independence Day'),
	(30,'20-08-2020','20-08-2020','Islamic New Year 1442 Hijriyah​'),
	(31,'29-10-2020','29-10-2020','Prophet Muhammads Birthday'),
	(32,'25-12-2020','25-12-2020','Christmas Holiday​')
ON CONFLICT (id)
DO NOTHING;