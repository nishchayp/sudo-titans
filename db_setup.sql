-- call admin route to initialize present ctf values
-- add records in PointsAndAccess and TotalPoints for all teams

-- create admin
insert into users(team_name, password, name1, name2, reg_no1, reg_no2) values("admin", "adminPassword", "nish", "lennon", "160905298", "160905299");
insert into admins(team_name) values("admin");

-- create teams
insert into users(team_name, password, name1, name2, reg_no1, reg_no2) values("t1", "t1Password", "nish1", "lennon1", "1609052981", "1609052991");
insert into users(team_name, password, name1, name2, reg_no1, reg_no2) values("t2", "t2Password", "nish1", "lennon1", "1609052982", "1609052992");
insert into users(team_name, password, name1, name2, reg_no1, reg_no2) values("t3", "t3Password", "nish1", "lennon1", "1609052983", "1609052993");
insert into users(team_name, password, name1, name2, reg_no1, reg_no2) values("t4", "t4Password", "nish1", "lennon1", "1609052984", "1609052994");
insert into users(team_name, password, name1, name2, reg_no1, reg_no2) values("t5", "t5Password", "nish1", "lennon1", "1609052985", "1609052995");

-- create entries of teams in PointsAndAccess
insert into points_and_accesses(team_name) values("t1");
insert into points_and_accesses(team_name) values("t2");
insert into points_and_accesses(team_name) values("t3");
insert into points_and_accesses(team_name) values("t4");
insert into points_and_accesses(team_name) values("t5");

-- add mcq questions
insert into mcq_details(
	question_id, 
	level, 
	question_1, 
	option_1_a, option_1_b, option_1_c, option_1_d, 
	question_2, 
	option_2_a, option_2_b, option_2_c, option_2_d, 
	question_3, 
	option_3_a, option_3_b, option_3_c, option_3_d, 
	question_4, 
	option_4_a, option_4_b, option_4_c, option_4_d) 
values(
	"MCQ_ ",
	1,
	"",
	"", "", "", "",
	"",
	"", "", "", "",
	"",
	"", "", "", "",
	"",
	"", "", "", ""
);

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

-- add ctf questions
insert into ctf_details(
	question_id, 
	level, 
	question, 
	hint) 
values(
	"CTF_ ", 
	1, 
	"question", 
	"hint"
); 
insert into ctf_details values(1, "CTF_1", 1, "Look beneath the surface", ""); 

-- add flags
insert into flags(
	question_id,
	flag) 
values( 
	"CTF_ ", 
	"flag"
);
insert into flags values(2, "CTF_1", "pwofyoda");
