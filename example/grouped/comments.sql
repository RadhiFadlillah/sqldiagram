CREATE TABLE `wp_commentmeta`(
    `meta_id` BIGINT NOT NULL,
    `comment_id` BIGINT NOT NULL,
    `meta_key` VARCHAR(255) NULL,
    `meta_value` LONGTEXT NULL
);
ALTER TABLE
    `wp_commentmeta` ADD PRIMARY KEY(`meta_id`);

CREATE TABLE `wp_comments`(
    `comment_ID` BIGINT NOT NULL,
    `comment_post_ID` BIGINT NOT NULL,
    `comment_author` TEXT NOT NULL,
    `comment_author_email` VARCHAR(255) NOT NULL,
    `comment_author_url` VARCHAR(255) NOT NULL,
    `comment_author_IP` VARCHAR(255) NOT NULL,
    `comment_date` DATETIME NOT NULL,
    `comment_date_gmt` DATETIME NOT NULL,
    `comment_content` TEXT NOT NULL,
    `comment_karma` INT NOT NULL,
    `comment_approved` VARCHAR(255) NOT NULL,
    `comment_agent` VARCHAR(255) NOT NULL,
    `comment_type` VARCHAR(255) NOT NULL,
    `comment_parent` BIGINT NOT NULL,
    `user_id` BIGINT NOT NULL
);
ALTER TABLE
    `wp_comments` ADD PRIMARY KEY(`comment_ID`);

ALTER TABLE
    `wp_commentmeta` ADD CONSTRAINT `wp_commentmeta_comment_id_foreign` FOREIGN KEY(`comment_id`) REFERENCES `wp_comments`(`comment_ID`);
ALTER TABLE
    `wp_comments` ADD CONSTRAINT `wp_comments_comment_post_id_foreign` FOREIGN KEY(`comment_post_ID`) REFERENCES `wp_posts`(`ID`);
ALTER TABLE
    `wp_comments` ADD CONSTRAINT `wp_comments_user_id_foreign` FOREIGN KEY(`user_id`) REFERENCES `wp_users`(`ID`);