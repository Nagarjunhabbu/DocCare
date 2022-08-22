Create Table users(
    id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name varchar(255) NOT null,
    place varchar(255) default null,
    email varchar(255) default null,
    password varchar(255) default null
);

insert into users(name,place,email,password)values("Nagarjun","Banglore","mrnags14@gmail.com","hello21");
insert into users(name,place,email,password)values("Yash","Manglore","yash114@gmail.com","wassup21");
insert into users(name,place,email,password)values("Ganesh","Sirsi","ganesh66@gmail.com","rock45");


