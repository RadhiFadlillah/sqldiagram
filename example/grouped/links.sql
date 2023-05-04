CREATE TABLE `wp_links`(
    `link_id` BIGINT NOT NULL,
    `link_url` VARCHAR(255) NOT NULL,
    `link_name` VARCHAR(255) NOT NULL,
    `link_image` VARCHAR(255) NOT NULL,
    `link_target` VARCHAR(255) NOT NULL,
    `link_description` VARCHAR(255) NOT NULL,
    `link_visible` VARCHAR(255) NOT NULL,
    `link_owner` BIGINT NOT NULL,
    `link_rating` INT NOT NULL,
    `link_updated` DATETIME NOT NULL,
    `link_rel` VARCHAR(255) NOT NULL,
    `link_notes` MEDIUMTEXT NOT NULL,
    `link_rss` VARCHAR(255) NOT NULL
);
ALTER TABLE
    `wp_links` ADD PRIMARY KEY(`link_id`);

ALTER TABLE
    `wp_links` ADD CONSTRAINT `wp_links_link_owner_foreign` FOREIGN KEY(`link_owner`) REFERENCES `wp_users`(`ID`);