CREATE TABLE `rinha`.`clients` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `limit_account` INT NOT NULL,
  `balance` INT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
);

CREATE TABLE `rinha`.`transactions` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `client_id` INT UNSIGNED NOT NULL,
  `type` VARCHAR(1) NOT NULL,
  `description` VARCHAR(10) NOT NULL,
  `value` INT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  CONSTRAINT fk_transactions_client_id
        FOREIGN KEY (client_id) REFERENCES clients (id)
);

DO $$
BEGIN
  INSERT INTO clients (name, limit_account)
  VALUES
    ('o barato sai caro', 1000 * 100),
    ('zan corp ltda', 800 * 100),
    ('les cruders', 10000 * 100),
    ('padaria joia de cocaia', 100000 * 100),
    ('kid mais', 5000 * 100);
END; $$