\c sample;
-- LOAD DATA LOCAL INFILE "/docker-entrypoint-initdb.d/horse.csv" INTO TABLE HORSE FIELDS TERMINATED BY ',';
copy horse from '/docker-entrypoint-initdb.d/horse.csv' ( delimiter ',', format csv, header true );
copy race from '/docker-entrypoint-initdb.d/race.csv'( delimiter ',', format csv, header true );
copy horse_race_result from '/docker-entrypoint-initdb.d/horse_race_result.csv'( delimiter ',', format csv, header true );