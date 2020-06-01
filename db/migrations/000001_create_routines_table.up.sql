CREATE TABLE IF NOT EXISTS routines(
   routine_id serial PRIMARY KEY,
   name VARCHAR (50) NOT NULL,
   description VARCHAR (300),
   level VARCHAR (20)
);