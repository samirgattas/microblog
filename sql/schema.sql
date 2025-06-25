-- THIS SCRIPT IS TO CREATE TABLES WHEN RUN THE MYSQL DOCKER CONTAINER

USE microblog;

CREATE TABLE User (
  user_id INT NOT NULL,
  nickname VARCHAR(45) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id),
  UNIQUE INDEX `nickname_unique_idx` (nickname)
);

CREATE TABLE Followed (
  id INT NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  followed_user_id INT NOT NULL,
  enabled TINYINT NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  update_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE INDEX `user_id_followed_user_id_unique_idx` (user_id, followed_user_id)
);