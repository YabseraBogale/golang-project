create table Software(
	id int not null primary key,
	message text not null,
	date datetime not null
);

create table Compaine(
	id int not null primary key,
	numberOfRequest int not null,
	listOfStack text not null
);