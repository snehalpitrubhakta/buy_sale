drop table cart;
drop table solditems;
/*
drop table paymentdetails;
drop table itemdetails;
*/
/*
create table if not exists itemdetails(
	itemname text,
	price float not null,
	primary key(itemname)
);
*/
create table if not exists cart(
	srNo int not null AUTO_INCREMENT,
	username varchar(20),
	itemname varchar(20),
	quantity varchar(10),
	quality text check (quality in ('high','medium','low')),
	price float not null,
	primary key(srNo)
);
/*
create table if not exists paymentdetails(
	username varchar(20),
	itemid varchar(20),
	paymenttype text check (paymenttype in('cod','card')),
	foreign key(itemid) references itemdetails(itemid),
	foreign key(username) references userdetails(username)	
);
*/
create table if not exists solditems(
	orderid int not null AUTO_INCREMENT,
	username varchar(20) not null,
	itemname varchar(20) not null,
	quantity varchar(10),
	quality text check (quality in ('high','medium','low')),
	address varchar(40) not null,
	primary key(orderid)
);	
