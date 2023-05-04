CREATE TABLE `wp_sitemeta`(
    `meta_id` BIGINT NOT NULL,
    `site_id` BIGINT NOT NULL,
    `meta_key` VARCHAR(255) NULL,
    `meta_value` LONGTEXT NULL
);
ALTER TABLE
    `wp_sitemeta` ADD PRIMARY KEY(`meta_id`);

CREATE TABLE `wp_site`(
    `id` BIGINT NOT NULL,
    `domain` VARCHAR(255) NOT NULL,
    `path` VARCHAR(255) NOT NULL
);
ALTER TABLE
    `wp_site` ADD PRIMARY KEY(`id`);

ALTER TABLE
    `wp_sitemeta` ADD CONSTRAINT `wp_sitemeta_site_id_foreign` FOREIGN KEY(`site_id`) REFERENCES `wp_site`(`id`);