create table users(
   name VARCHAR(100) NOT NULL,
   id integer NOT NULL AUTO_INCREMENT,
   surname VARCHAR(100) NOT NULL,
   email VARCHAR(100) NOT NULL,
   token CHAR(36) NOT NULL UNIQUE,
   PRIMARY KEY ( id )
);

create table users_credentials(
   user VARCHAR(100) NOT NULL UNIQUE,
   id integer NOT NULL AUTO_INCREMENT,
   password VARCHAR(100) NOT NULL,
   id_user integer NOT NULL,
   PRIMARY KEY (id),
   FOREIGN KEY (id_user) REFERENCES users(id)
);


INSERT INTO users (name, surname, email, token) VALUES ("prueba", "prueba", "prueba@gmail", UUID());
INSERT INTO users_credentials (user, password, id_user) VALUES ("user", "12345", 1);
