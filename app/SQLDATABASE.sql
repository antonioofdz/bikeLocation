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

create table bike(
   id integer not null AUTO_INCREMENT,
   model varchar(300) not null,
   PRIMARY KEY (id)
);

create table bike_location(
   id integer not null AUTO_INCREMENT,
   lat float,
   lon float,
   address varchar(300),
   id_bike integer not null,
   PRIMARY KEY (id),
   FOREIGN KEY (id_bike) REFERENCES bike(id)
);

create table user_bike(
   id integer not null AUTO_INCREMENT,
   id_bike integer not null,
   id_user integer not null,
   booked boolean not null,
   dateRent datetime, 
   dateReturn datetime,
   PRIMARY KEY (id),
   FOREIGN KEY (id_user) REFERENCES users(id),
   FOREIGN KEY (id_bike) REFERENCES bike(id)
);

INSERT INTO users (name, surname, email, token) VALUES ("prueba", "prueba", "prueba@gmail", UUID());
INSERT INTO users_credentials (user, password, id_user) VALUES ("user", "12345", 1);

INSERT INTO bike (model) VALUES ("model1");
INSERT INTO bike (model) VALUES ("model2");
INSERT INTO bike (model) VALUES ("model3");

INSERT INTO bike_location (lat, lon, address, id_bike) VALUES (1, 2, "address1", 1);
INSERT INTO bike_location (lat, lon, address,id_bike) VALUES (1, 2, "address2", 2);
INSERT INTO bike_location (lat, lon, address,id_bike) VALUES (1, 2, "address3", 3);

INSERT INTO user_bike (id_bike, id_user,booked) VALUES (1, 1, 0);