create table vulnerabilities(
    id int not null primary key,
    vulnerability varchar(332) not null,
    year varchar(4) not null,
    attackComplexity varchar(10) not null,
    baseScore varchar(10) not null
)