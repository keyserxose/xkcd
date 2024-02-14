### Queries

INSERT INTO files (filename, executed, executionDate)
VALUES("thicket", 1, "2024-02-11 12:00:00");

DELETE FROM files;
DELETE FROM SQLITE_SEQUENCE WHERE NAME = 'files';

select id,filename,executionDate from files where executed = 1 order by id asc limit 1