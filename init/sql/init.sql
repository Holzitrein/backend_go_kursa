-- CREATE DATABASE mirea;
CREATE TABLE Groups
(
 idGroup SERIAL,
 nameGroup varchar(15) NOT NULL,
 PRIMARY KEY ( idGroup )
);

CREATE TABLE Subjects
(
    idSubject SERIAL,
    name varchar(50) NOT NULL,
    type int NOT NULL,
    PRIMARY KEY ( idSubject )
);

CREATE TABLE Teachers
(
    idTeacher SERIAL,
    name varchar(20) NOT NULL,
    secondName varchar(20) NOT NULL,
    department varchar(50) NOT NULL,
    PRIMARY KEY ( idTeacher )
);

CREATE TABLE Places
(
    idPlace SERIAL,
    namePlace varchar(20) NOT NULL,
    proector boolean NOT NULL,
    numSeats int NOT NULL,
    numComputers int NOT NULL,
    PRIMARY KEY ( idPlace )
);

CREATE TABLE Students
(
    idStudent SERIAL,
    name varchar(20) NOT NULL,
    secondName varchar(20) NOT NULL,
    grou int NOT NULL,
    PRIMARY KEY ( idStudent ),
    FOREIGN KEY ( grou ) REFERENCES Groups ( idGroup )
);

CREATE TABLE Retakes
(
    idRetake SERIAL,
    teacher int NOT NULL,
    subject int NOT NULL,
    timeOn timestamp NOT NULL,
    place int NOT NULL,
    PRIMARY KEY ( idRetake ),
    FOREIGN KEY ( teacher ) REFERENCES Teachers ( idTeacher ),
    FOREIGN KEY ( subject ) REFERENCES Subjects ( idSubject ),
    FOREIGN KEY ( place ) REFERENCES Places ( idPlace )
);
Create table Students_retake
(
    idRetakes SERIAL,
    student int not null,
    retake int not null,
    primary key (idRetakes),
    FOREIGN KEY ( student ) REFERENCES Students ( idStudent ),
    FOREIGN KEY ( retake ) REFERENCES Retakes ( idRetake )
);