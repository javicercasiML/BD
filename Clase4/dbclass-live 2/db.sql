DROP DATABASE IF EXISTS warehouse_db;
CREATE DATABASE IF NOT EXISTS warehouse_db;

USE warehouse_db;

CREATE TABLE IF NOT EXISTS `warehouses` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `address` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO `warehouses` (`name`, `address`) VALUES ("Fresh", "Venecia, Italy");
INSERT INTO `warehouses` (`name`, `address`) VALUES ("DHL", "New York, US");

CREATE TABLE IF NOT EXISTS `products` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `type` varchar(100) NOT NULL,
  `count` int NOT NULL,
  `price` float NOT NULL,
  `warehouse_id` int NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses` (`id`)
);

INSERT INTO `products` (`name`, `type`, `count`, `price`, `warehouse_id`) VALUES ("Coca-Poca", "Bebidas", 5, 105.5, 1);
INSERT INTO `products` (`name`, `type`, `count`, `price`, `warehouse_id`) VALUES ("Freezo-Ice", "Helados", 7, 75, 1);
INSERT INTO `products` (`name`, `type`, `count`, `price`, `warehouse_id`) VALUES ("Papas-PLays", "Snacks", 3, 40, 2);