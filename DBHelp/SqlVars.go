package DBHelp

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
