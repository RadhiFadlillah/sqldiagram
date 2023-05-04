CREATE TABLE `wp_commentmeta`(
    `meta_id` BIGINT NOT NULL,
    `comment_id` BIGINT NOT NULL,
    `meta_key` VARCHAR(255) NULL,
    `meta_value` LONGTEXT NULL
);
ALTER TABLE
    `wp_commentmeta` ADD PRIMARY KEY(`meta_id`);
CREATE TABLE `wp_term_relationships`(
    `object_id` BIGINT NOT NULL,
    `term_taxonomy_id` BIGINT NOT NULL,
    `term_order` INT NOT NULL
);
CREATE TABLE `wp_sitemeta`(
    `meta_id` BIGINT NOT NULL,
    `site_id` BIGINT NOT NULL,
    `meta_key` VARCHAR(255) NULL,
    `meta_value` LONGTEXT NULL
);
ALTER TABLE
    `wp_sitemeta` ADD PRIMARY KEY(`meta_id`);
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
CREATE TABLE `wp_terms`(
    `term_id` BIGINT NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `slug` VARCHAR(255) NOT NULL,
    `term_group` BIGINT NOT NULL
);
ALTER TABLE
    `wp_terms` ADD PRIMARY KEY(`term_id`);
CREATE TABLE `wp_site`(
    `id` BIGINT NOT NULL,
    `domain` VARCHAR(255) NOT NULL,
    `path` VARCHAR(255) NOT NULL
);
ALTER TABLE
    `wp_site` ADD PRIMARY KEY(`id`);
CREATE TABLE `wp_termmeta`(
    `meta_id` BIGINT NOT NULL,
    `term_id` BIGINT NOT NULL,
    `meta_key` VARCHAR(255) NULL,
    `meta_value` LONGTEXT NULL
);
ALTER TABLE
    `wp_termmeta` ADD PRIMARY KEY(`meta_id`);
CREATE TABLE `wp_postmeta`(
    `meta_id` BIGINT NOT NULL,
    `post_id` BIGINT NOT NULL,
    `meta_key` VARCHAR(255) NULL,
    `meta_value` LONGTEXT NULL
);
ALTER TABLE
    `wp_postmeta` ADD PRIMARY KEY(`meta_id`);
CREATE TABLE `wp_signups`(
    `signup_id` BIGINT NOT NULL,
    `domain` VARCHAR(255) NOT NULL,
    `path` VARCHAR(255) NOT NULL,
    `title` LONGTEXT NOT NULL,
    `user_login` VARCHAR(255) NOT NULL,
    `user_email` VARCHAR(255) NOT NULL,
    `registered` DATETIME NOT NULL,
    `activated` DATETIME NOT NULL,
    `active` TINYINT NOT NULL,
    `activation_key` VARCHAR(255) NOT NULL,
    `meta` LONGTEXT NULL
);
ALTER TABLE
    `wp_signups` ADD PRIMARY KEY(`signup_id`);
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
CREATE TABLE `wp_registration_log`(
    `ID` BIGINT NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `IP` VARCHAR(255) NOT NULL,
    `blog_id` BIGINT NOT NULL,
    `date_registered` DATETIME NOT NULL
);
ALTER TABLE
    `wp_registration_log` ADD PRIMARY KEY(`ID`);
CREATE TABLE `wp_options`(
    `option_id` BIGINT NOT NULL,
    `option_name` VARCHAR(255) NOT NULL,
    `option_value` LONGTEXT NOT NULL,
    `autoload` VARCHAR(255) NOT NULL
);
ALTER TABLE
    `wp_options` ADD PRIMARY KEY(`option_id`);
CREATE TABLE `wp_users`(
    `ID` BIGINT NOT NULL,
    `user_login` VARCHAR(255) NOT NULL,
    `user_pass` VARCHAR(255) NOT NULL,
    `user_nicename` VARCHAR(255) NOT NULL,
    `user_email` VARCHAR(255) NOT NULL,
    `user_url` VARCHAR(255) NOT NULL,
    `user_registered` DATETIME NOT NULL,
    `user_activation_key` VARCHAR(255) NOT NULL,
    `user_status` INT NOT NULL,
    `display_name` VARCHAR(255) NOT NULL
);
ALTER TABLE
    `wp_users` ADD PRIMARY KEY(`ID`);
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
CREATE TABLE `wp_usermeta`(
    `umeta_id` BIGINT NOT NULL,
    `user_id` BIGINT NOT NULL,
    `meta_key` VARCHAR(255) NULL,
    `meta_value` LONGTEXT NULL
);
ALTER TABLE
    `wp_usermeta` ADD PRIMARY KEY(`umeta_id`);
ALTER TABLE
    `wp_posts` ADD CONSTRAINT `wp_posts_post_parent_foreign` FOREIGN KEY(`post_parent`) REFERENCES `wp_posts`(`ID`);
ALTER TABLE
    `wp_termmeta` ADD CONSTRAINT `wp_termmeta_term_id_foreign` FOREIGN KEY(`term_id`) REFERENCES `wp_terms`(`term_id`);
ALTER TABLE
    `wp_term_taxonomy` ADD CONSTRAINT `wp_term_taxonomy_term_id_foreign` FOREIGN KEY(`term_id`) REFERENCES `wp_terms`(`term_id`);
ALTER TABLE
    `wp_sitemeta` ADD CONSTRAINT `wp_sitemeta_site_id_foreign` FOREIGN KEY(`site_id`) REFERENCES `wp_site`(`id`);
ALTER TABLE
    `wp_postmeta` ADD CONSTRAINT `wp_postmeta_post_id_foreign` FOREIGN KEY(`post_id`) REFERENCES `wp_posts`(`ID`);
ALTER TABLE
    `wp_commentmeta` ADD CONSTRAINT `wp_commentmeta_comment_id_foreign` FOREIGN KEY(`comment_id`) REFERENCES `wp_comments`(`comment_ID`);
ALTER TABLE
    `wp_term_relationships` ADD CONSTRAINT `wp_term_relationships_object_id_foreign` FOREIGN KEY(`object_id`) REFERENCES `wp_posts`(`ID`);
ALTER TABLE
    `wp_comments` ADD CONSTRAINT `wp_comments_comment_post_id_foreign` FOREIGN KEY(`comment_post_ID`) REFERENCES `wp_posts`(`ID`);
ALTER TABLE
    `wp_links` ADD CONSTRAINT `wp_links_link_owner_foreign` FOREIGN KEY(`link_owner`) REFERENCES `wp_users`(`ID`);
ALTER TABLE
    `wp_posts` ADD CONSTRAINT `wp_posts_post_author_foreign` FOREIGN KEY(`post_author`) REFERENCES `wp_users`(`ID`);
ALTER TABLE
    `wp_comments` ADD CONSTRAINT `wp_comments_user_id_foreign` FOREIGN KEY(`user_id`) REFERENCES `wp_users`(`ID`);
ALTER TABLE
    `wp_term_relationships` ADD CONSTRAINT `wp_term_relationships_term_taxonomy_id_foreign` FOREIGN KEY(`term_taxonomy_id`) REFERENCES `wp_term_taxonomy`(`term_taxonomy_id`);
ALTER TABLE
    `wp_usermeta` ADD CONSTRAINT `wp_usermeta_user_id_foreign` FOREIGN KEY(`user_id`) REFERENCES `wp_users`(`ID`);