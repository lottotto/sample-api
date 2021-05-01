CREATE DATABASE sample;
\c sample;

CREATE TABLE HORSE (
    id  VARCHAR PRIMARY KEY,
    name VARCHAR NOT NULL,
    trainer VARCHAR NOT NULL,
    owner VARCHAR NOT NULL,
    breeder VARCHAR NOT NULL,
    sire VARCHAR NOT NULL,
    broodmare VARCHAR NOT NULL 
);

CREATE TABLE RACE (
    id VARCHAR PRIMARY KEY,
    name VARCHAR NOT NULL,
    distance INTEGER NOT NULL,
    field VARCHAR NOT NULL,
    rotation VARCHAR NOT NULL,
    condition VARCHAR NOT NULL,
    cource VARCHAR NOT NULL
);

CREATE TABLE HORSE_RACE_RESULT (
    raceid VARCHAR NOT NULL,
    horseid VARCHAR NOT NULL,
    arrival VARCHAR NOT NULL,
    bracket INTEGER NOT NULL,
    horse_number INTEGER NOT NULL,
    weight numeric NOT NULL,
    jockey VARCHAR NOT NULL,
    time VARCHAR,
    diff VARCHAR,
    position VARCHAR,
    last3F numeric,
    odds VARCHAR,
    favorite numeric,
    horse_weight VARCHAR,
    trainer VARCHAR NOT NULL,
    owner VARCHAR NOT NULL
)