CREATE DATABASE lottery;

USE lottery;

/*all peopple that created an account*/
CREATE TABLE `users` (
    userid INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY, 
    username VARCHAR(100) NOT NULL
    /* user_organization VARCHAR(100),
    user_email VARCHAR(100) NOT NULL */ 
);

/*event table*/ 
CREATE TABLE `events` ( 
    event_id INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    event_name VARCHAR(100) NOT NULL,
    event_description VARCHAR(100) NOT NULL,
    event_total_attend VARCHAR(100) NOT NULL,
    event_start_date VARCHAR(100) NOT NULL,
    event_end_date VARCHAR(100) NOT NULL
);

/*all users that registered for the event*/
CREATE TABLE registration (
    registrationid INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    registrationDate VARCHAR(100) NOT NULL, 
    event_id INT(10) NOT NULL,
    userID INT(10) NOT NULL,
    FOREIGN KEY (event_id) REFERENCES events(event_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (userID) REFERENCES users(userid) ON UPDATE CASCADE ON DELETE CASCADE
);

/*all users that are accepted after the lottery*/ 
CREATE TABLE attendee (
    attendeeid INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    registration_id INT(10) NOT NULL, 
    attendee_name VARCHAR(100) NOT NULL,
    attendee_email VARCHAR(100) NOT NULL,
    FOREIGN KEY (registration_id) REFERENCES registration(registrationid) ON UPDATE CASCADE ON DELETE CASCADE
);

/*all users that were accepted to the event and actually attended event*/
CREATE TABLE eventAttendance(
    eventAttendanceID INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    attendee INT(10) NOT NULL,
    event_id INT(10) NOT NULL,
    FOREIGN KEY (attendee) REFERENCES attendee(attendeeid) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (event_id) REFERENCES events(event_id) ON UPDATE CASCADE ON DELETE CASCADE
);

/* Test data */
/*INSERT INTO users username VALUES ('marianne');*/