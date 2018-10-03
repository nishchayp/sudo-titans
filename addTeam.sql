-- insert into users(
-- 	team_name, 
-- 	password, 
-- 	name1, 
-- 	name2, 
-- 	reg_no1, 
-- 	reg_no2) 
-- values(
-- 	"t1", 
-- 	"t1Password", 
-- 	"nish1", 
-- 	"lennon1", 
-- 	"1609052981", 
-- 	"1609052991"
-- );

-- insert into points_and_accesses(team_name) values("t1");

-- create teams
insert into users(team_name, password, name1, name2, reg_no1, reg_no2) values("t1", "t1Password", "nish1", "lennon1", "1609052981", "1609052991");
insert into users(team_name, password, name1, name2, reg_no1, reg_no2) values("t2", "t2Password", "nish1", "lennon1", "1609052982", "1609052992");
insert into users(team_name, password, name1, name2, reg_no1, reg_no2) values("t3", "t3Password", "nish1", "lennon1", "1609052983", "1609052993");
insert into users(team_name, password, name1, name2, reg_no1, reg_no2) values("t4", "t4Password", "nish1", "lennon1", "1609052984", "1609052994");
insert into users(team_name, password, name1, name2, reg_no1, reg_no2) values("t5", "t5Password", "nish1", "lennon1", "1609052985", "1609052995");
insert into users(team_name, password, name1, name2, reg_no1, reg_no2) values("t6", "t6Password", "nish1", "lennon1", "1609052986", "1609052996");
insert into users(team_name, password, name1, name2, reg_no1, reg_no2) values("t7", "t7Password", "nish1", "lennon1", "1609052987", "1609052997");
insert into users(team_name, password, name1, name2, reg_no1, reg_no2) values("t8", "t8Password", "nish1", "lennon1", "1609052988", "1609052998");

-- create entries of teams in PointsAndAccess
insert into points_and_accesses(team_name) values("t1");
insert into points_and_accesses(team_name) values("t2");
insert into points_and_accesses(team_name) values("t3");
insert into points_and_accesses(team_name) values("t4");
insert into points_and_accesses(team_name) values("t5");
insert into points_and_accesses(team_name) values("t6");
insert into points_and_accesses(team_name) values("t7");
insert into points_and_accesses(team_name) values("t8");