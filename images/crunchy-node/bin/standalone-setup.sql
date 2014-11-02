
create user testdb login superuser password 'testdb';

create database testdb;

grant all privileges on database testdb to testdb;

\c testdb;

create extension adminpack;

