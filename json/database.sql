create table vulnerabilities(
    id int not null primary key,
    vulnerability varchar(308) not null,
    year int not null,
    attackComplexity varchar(4) not null,
    baseScore int not null
)