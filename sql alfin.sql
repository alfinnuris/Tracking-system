
CREATE TABLE IF NOT EXISTS `customers` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `username` longtext,
  `address` longtext,
  `user_id` bigint DEFAULT NULL,
  `license_id` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_customers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- Pengeluaran data tidak dipilih.

CREATE TABLE IF NOT EXISTS `drivers` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `user_id` bigint DEFAULT NULL,
  `license_id` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_drivers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- Pengeluaran data tidak dipilih.

CREATE TABLE IF NOT EXISTS `shipments` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `po_number` longtext,
  `driver` longtext,
  `driver_name` longtext,
  `customer` longtext,
  `customer_name` longtext,
  `delivery_schedule` longtext,
  `image_barcode` longtext,
  `shipment_status` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_shipments_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- Pengeluaran data tidak dipilih.

CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` longtext,
  `email` longtext,
  `password` longtext,
  `role` longtext,
  `name` longtext,
  `phone` longtext,
  `address` longtext,
  `image` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4;

-- Pengeluaran data tidak dipilih.

