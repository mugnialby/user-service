SELECT * FROM USERS;

DROP TABLE USERS;

CREATE TABLE users (
  id SERIAL PRIMARY KEY, 
  username VARCHAR (64) UNIQUE NOT NULL, 
  password VARCHAR (256) NOT NULL,  
  first_name VARCHAR (128) NOT NULL,  
  last_name VARCHAR (128) NOT NULL, 
  created_by VARCHAR (64), 
  created_at TIMESTAMP
);