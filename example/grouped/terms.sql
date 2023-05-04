CREATE TABLE `wp_term_relationships`(
    `object_id` BIGINT NOT NULL,
    `term_taxonomy_id` BIGINT NOT NULL,
    `term_order` INT NOT NULL
);

CREATE TABLE `wp_terms`(
    `term_id` BIGINT NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `slug` VARCHAR(255) NOT NULL,
    `term_group` BIGINT NOT NULL
);
ALTER TABLE
    `wp_terms` ADD PRIMARY KEY(`term_id`);

CREATE TABLE `wp_termmeta`(
    `meta_id` BIGINT NOT NULL,
    `term_id` BIGINT NOT NULL,
    `meta_key` VARCHAR(255) NULL,
    `meta_value` LONGTEXT NULL
);
ALTER TABLE
    `wp_termmeta` ADD PRIMARY KEY(`meta_id`);

CREATE TABLE `wp_term_taxonomy`(
    `term_taxonomy_id` BIGINT NOT NULL,
    `term_id` BIGINT NOT NULL,
    `taxonomy` VARCHAR(255) NOT NULL,
    `description` LONGTEXT NOT NULL,
    `parent` BIGINT NOT NULL,
    `count` BIGINT NOT NULL
);
ALTER TABLE
    `wp_term_taxonomy` ADD PRIMARY KEY(`term_taxonomy_id`);

ALTER TABLE
    `wp_termmeta` ADD CONSTRAINT `wp_termmeta_term_id_foreign` FOREIGN KEY(`term_id`) REFERENCES `wp_terms`(`term_id`);
ALTER TABLE
    `wp_term_taxonomy` ADD CONSTRAINT `wp_term_taxonomy_term_id_foreign` FOREIGN KEY(`term_id`) REFERENCES `wp_terms`(`term_id`);
ALTER TABLE
    `wp_term_relationships` ADD CONSTRAINT `wp_term_relationships_object_id_foreign` FOREIGN KEY(`object_id`) REFERENCES `wp_posts`(`ID`);
ALTER TABLE
    `wp_term_relationships` ADD CONSTRAINT `wp_term_relationships_term_taxonomy_id_foreign` FOREIGN KEY(`term_taxonomy_id`) REFERENCES `wp_term_taxonomy`(`term_taxonomy_id`);