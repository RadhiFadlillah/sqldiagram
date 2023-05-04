CREATE TABLE `wp_posts`(
    `ID` BIGINT NOT NULL,
    `post_author` BIGINT NOT NULL,
    `post_date` DATETIME NOT NULL,
    `post_date_gmt` DATETIME NOT NULL,
    `post_content` LONGTEXT NOT NULL,
    `post_title` TEXT NOT NULL,
    `post_excerpt` TEXT NOT NULL,
    `post_status` VARCHAR(255) NOT NULL,
    `comment_status` VARCHAR(255) NOT NULL,
    `ping_status` VARCHAR(255) NOT NULL,
    `post_password` VARCHAR(255) NOT NULL,
    `post_name` VARCHAR(255) NOT NULL,
    `to_ping` TEXT NOT NULL,
    `pinged` TEXT NOT NULL,
    `post_modified` DATETIME NOT NULL,
    `post_modified_gmt` DATETIME NOT NULL,
    `post_content_filtered` LONGTEXT NOT NULL,
    `post_parent` BIGINT NOT NULL,
    `guid` VARCHAR(255) NOT NULL,
    `menu_order` INT NOT NULL,
    `post_type` VARCHAR(255) NOT NULL,
    `post_mime_type` VARCHAR(255) NOT NULL,
    `comment_count` BIGINT NOT NULL
);
ALTER TABLE
    `wp_posts` ADD PRIMARY KEY(`ID`);

CREATE TABLE `wp_postmeta`(
    `meta_id` BIGINT NOT NULL,
    `post_id` BIGINT NOT NULL,
    `meta_key` VARCHAR(255) NULL,
    `meta_value` LONGTEXT NULL
);
ALTER TABLE
    `wp_postmeta` ADD PRIMARY KEY(`meta_id`);

ALTER TABLE
    `wp_posts` ADD CONSTRAINT `wp_posts_post_parent_foreign` FOREIGN KEY(`post_parent`) REFERENCES `wp_posts`(`ID`);
ALTER TABLE
    `wp_postmeta` ADD CONSTRAINT `wp_postmeta_post_id_foreign` FOREIGN KEY(`post_id`) REFERENCES `wp_posts`(`ID`);
ALTER TABLE
    `wp_posts` ADD CONSTRAINT `wp_posts_post_author_foreign` FOREIGN KEY(`post_author`) REFERENCES `wp_users`(`ID`);