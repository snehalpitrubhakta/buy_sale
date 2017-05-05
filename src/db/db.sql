drop table cart;
drop table solditems;
drop table paymentdetails;
drop table userdetails;
drop table itemdetails;

create table if not exists userdetails (
	username varchar(20), 
	email varchar(20) not null,
	password text not null,
	phonenumber text not null,
	address text not null,
	primary key(username)
);

create table if not exists itemdetails(
	itemid varchar(20),
	itemname text,
	unit varchar(2) check(unit in('kg')),
	price float not null,
	quality text check (quality in ('high','medium','low')),
	isavailable boolean default 1,
	primary key(itemid)
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
	orderid int,
	itemid varchar(20) not null,
	username varchar(20) not null,
	foreign key(itemid) references itemdetails(itemid),
	foreign key(username) references userdetails(username),
	primary key(orderid)
	
);
	