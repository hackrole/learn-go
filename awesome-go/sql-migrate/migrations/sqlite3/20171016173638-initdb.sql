
-- +migrate Up
create table people (id int);


-- +migrate Down
drop table people;
