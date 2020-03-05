
-- +migrate Up
CREATE TABLE `articles` (
    `id` INT NOT NULL,
    `name` VARCHAR (255),
    `description` text null,
    `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`)
);

-- +migrate Down
DROP TABLE articles;