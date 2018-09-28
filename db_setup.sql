-- call admin route to initialize present ctf values
-- add records in PointsAndAccess and TotalPoints for all teams

-- create admin
insert into users(team_name, password, name1, name2, reg_no1, reg_no2) values("admin", "adminPassword", "nish", "lennon", "160905298", "160905299");
insert into users(team_name, password, name1, name2, reg_no1, reg_no2) values("t1", "t1Password", "nish1", "lennon1", "1609052981", "1609052991");
insert into admins(team_name) values("admin");

-- create entries in PointsAndAccess
insert into points_and_accesses(team_name) values("t1");

-- add mcq questions
insert into mcq_details values(
	1,
	"MCQ_1",
	1,
	"1 lorem ipsum bksfbs fknsk abfsk fhls fanlflnf lnafln albfalba alafn aflbslnaf lbafl,
	lorem ipsum bksfbs fknsk abfsk fhls fanlflnf lnafln albfalba alafn aflbslnaf lbafl
	lorem ipsum bksfbs fknsk abfsk fhls fanlflnf lnafln albfalba alafn aflbslnaf lbafl
	lorem ipsum bksfbs fknsk abfsk fhls fanlflnf lnafln albfalba alafn aflbslnaf lbafl",
	"op1A", "op1B", "op1C", "op1D",
	"2 lorem ipsum bksfbs fknsk abfsk fhls fanlflnf lnafln albfalba alafn aflbslnaf lbafl,
	lorem ipsum bksfbs fknsk abfsk fhls fanlflnf lnafln albfalba alafn aflbslnaf lbafl
	lorem ipsum bksfbs fknsk abfsk fhls fanlflnf lnafln albfalba alafn aflbslnaf lbafl
	lorem ipsum bksfbs fknsk abfsk fhls fanlflnf lnafln albfalba alafn aflbslnaf lbafl",
	"op1A", "op1B", "op1C", "op1D",
	"3 lorem ipsum bksfbs fknsk abfsk fhls fanlflnf lnafln albfalba alafn aflbslnaf lbafl,
	lorem ipsum bksfbs fknsk abfsk fhls fanlflnf lnafln albfalba alafn aflbslnaf lbafl
	lorem ipsum bksfbs fknsk abfsk fhls fanlflnf lnafln albfalba alafn aflbslnaf lbafl
	lorem ipsum bksfbs fknsk abfsk fhls fanlflnf lnafln albfalba alafn aflbslnaf lbafl",
	"op1A", "op1B", "op1C", "op1D",
	"4 lorem ipsum bksfbs fknsk abfsk fhls fanlflnf lnafln albfalba alafn aflbslnaf lbafl,
	lorem ipsum bksfbs fknsk abfsk fhls fanlflnf lnafln albfalba alafn aflbslnaf lbafl
	lorem ipsum bksfbs fknsk abfsk fhls fanlflnf lnafln albfalba alafn aflbslnaf lbafl
	lorem ipsum bksfbs fknsk abfsk fhls fanlflnf lnafln albfalba alafn aflbslnaf lbafl",
	"op1A", "op1B", "op1C", "op1D"
	);

-- add mcq questions
	insert into ctf_details values(1, "CTF_1", 1, "Look beneath the surface", ""); 

-- add flags
insert into flags values(2, "CTF_1", "pwofyoda");
