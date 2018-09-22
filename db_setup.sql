-- create admin
insert into users(team_name, password, name1, name2, reg_no1, reg_no2) values("admin", "adminPassword", "nish", "lennon", "160905298", "160905299");
insert into admins(team_name) values("admin");

-- add mcq questions
insert into mcq_details values(
	1,
	"MCQ_1",
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
