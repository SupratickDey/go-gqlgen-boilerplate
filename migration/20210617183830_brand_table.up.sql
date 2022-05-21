CREATE TABLE brands (
    id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    created_at timestamp NOT NULL DEFAULT (now()),
    updated_at timestamp NOT NULL DEFAULT (now())
);

INSERT INTO `brands` (`id`, `name`, `created_at`, `updated_at`)
VALUES ('1', 'Logitech', '2021-06-16 16:05:48', '2021-06-16 16:05:48')