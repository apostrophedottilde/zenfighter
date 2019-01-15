CREATE TABLE IF NOT EXISTS `zenfighter`.`fighters` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(16) NOT NULL,
  `strength` INT(11) NOT NULL,
  `weaponpower` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `name_UNIQUE` (`name` ASC),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC))
ENGINE = InnoDB
AUTO_INCREMENT = 159
DEFAULT CHARACTER SET = latin1;

CREATE SCHEMA IF NOT EXISTS `zenfightertest` DEFAULT CHARACTER SET latin1 ;

GRANT ALL PRIVILEGES ON zenfightertest.* TO '*'@'*';

CREATE TABLE IF NOT EXISTS `zenfightertest`.`fighters` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(16) NOT NULL,
  `strength` INT(11) NOT NULL,
  `weaponpower` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `name_UNIQUE` (`name` ASC),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC))
ENGINE = InnoDB
AUTO_INCREMENT = 159
DEFAULT CHARACTER SET = latin1