CREATE TABLE `profiles` (
  id INTEGER UNSIGNED PRIMARY KEY UNIQUE NOT NULL AUTO_INCREMENT,
  account_id INTEGER UNSIGNED UNIQUE,
  bio VARCHAR(256) NOT NULL DEFAULT '',
  gender ENUM('m', 'f', 't', 'a', 'o'),
  country VARCHAR(2),
  created_at TIMESTAMP DEFAULT NOW()
) Engine=InnoDB;
