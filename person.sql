create table person
(
	id bigserial primary key,
	first_name name not null,
	sur_name name not null,
	sex char(1)
)