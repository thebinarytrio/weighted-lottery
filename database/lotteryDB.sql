CREATE DATABASE lottery;

USE lottery; 

CREATE TABLE users (
    userid INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY, 
    username VARCHAR(100) NOT NULL,
    user_organization VARCHAR(100),
    user_email VARCHAR(100) NOT NULL,
)

CREATE TABLE events ( 
    eventid INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    eventname VARCHAR(100) NOT NULL,
    event_description VARCHAR(100) NOT NULL,
    event_total_attend VARCHAR(100) NOT NULL,
    event_start_date VARCHAR(100) NOT NULL,
    event_end_date VARCHAR(100) NOT NULL
)

CREATE TABLE reservation (
    reservationid INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    reservationdate VARCHAR(100) NOT NULL, 
    event_id INT(10) NOT NULL,
    userID INT(10) NOT NULL,
    FOREIGN KEY (event_id) REFERENCES events(eventid) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (userID) REFERENCES users(userid) ON UPDATE CASCADE ON DELETE CASCADE
)

CREATE TABLE attendee (
    attendeeid INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    reservation_id INT(10) NOT NULL, 
    attendee_name VARCHAR(100) NOT NULL,
    attendee_email VARCHAR(100) NOT NULL,
    FOREIGN KEY (reservation_id) REFERENCES reservation(reservation_id) ON UPDATE CASCADE ON DELETE CASCADE
)

CREATE TABLE eventAttendance(
    eattendanceID INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    attendee INT(10) NOT NULL,
    event_id INT(10) NOT NULL,
    FOREIGN KEY (attendee) REFERENCES attendee(attendee_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (event_id) REFERENCES events(eventid) ON UPDATE CASCADE ON DELETE CASCADE
)
