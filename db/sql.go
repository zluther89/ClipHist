package db

var initDBSQL string = `
CREATE TABLE IF NOT EXISTS 
clip(content string unique, timestamp datetime default current_timestamp);
CREATE TRIGGER IF NOT EXISTS limiter AFTER INSERT ON clip
    BEGIN
      delete from clip where 
        timestamp =(select min(timestamp) from clip ) 
        and (select count(*) from clip )>100;
    END;
`

var insertSQL string = `INSERT OR REPLACE INTO clip(content) VALUES(?)`

var selectTopSQL string = `SELECT * FROM clip ORDER BY rowid desc`

var findSQL string = `SELECT content FROM clip WHERE timestamp = %v`
