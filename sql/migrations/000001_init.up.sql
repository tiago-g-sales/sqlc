CREATE TABLE categories (
	id	            varchar(36) NOT NULL PRIMARY KEY,
	name	        TEXT        NOT NULL,
	description	    TEXT
);

CREATE TABLE courses (
	id	            varchar(36)     NOT NULL PRIMARY KEY,
	category_id	    varchar(36)     NOT NULL,
	name 	        TEXT            NOT NULL,
	description 	TEXT,
    price           decimal(10,2)   NOT NULL,
    FOREIGN KEY     (category_id) REFERENCES categories(id)
);