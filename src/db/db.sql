drop table cart;
drop table solditems;
drop table itemquantity;
drop table paymentdetails;
drop table userdetails;
drop table itemdetails;

create table if not exists userdetails (
	username varchar(20), 
	email varchar(20),
	password text,
	phonenumber text,
	primary key(username)
);

create table if not exists itemdetails(
	itemid varchar(20),
	itemname text,
	quality text check (quality in ('high','medium','low')),
	isavailable boolean default 1,
	primary key(itemid)
);

create table  if not exists itemquantity(
	itemid varchar(20),
	unit varchar(5) check(unit in('kg','mg','tn')),
	price float,
	foreign key(itemid) references itemdetails(itemid),
	primary key(unit,itemid)
);

create table if not exists cart(
	username varchar(20),
	itemid varchar(20),
	foreign key(itemid) references itemdetails(itemid),
	foreign key(username) references userdetails(username)
);

create table if not exists paymentdetails(
	username varchar(20),
	itemid varchar(20),
	paymenttype text check (paymenttype in('cod','card')),
	foreign key(itemid) references itemdetails(itemid),
	foreign key(username) references userdetails(username)	
);

create table if not exists solditems(
	itemid varchar(20),
	username varchar(20),
	unit varchar(5),
	foreign key(itemid) references itemdetails(itemid),
	foreign key(username) references userdetails(username),
	foreign key(unit) references itemquantity(unit)
);
