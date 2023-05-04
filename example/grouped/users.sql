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

CREATE TABLE `wp_usermeta`(
    `umeta_id` BIGINT NOT NULL,
    `user_id` BIGINT NOT NULL,
    `meta_key` VARCHAR(255) NULL,
    `meta_value` LONGTEXT NULL
);

ALTER TABLE
    `wp_usermeta` ADD PRIMARY KEY(`umeta_id`);
ALTER TABLE
    `wp_usermeta` ADD CONSTRAINT `wp_usermeta_user_id_foreign` FOREIGN KEY(`user_id`) REFERENCES `wp_users`(`ID`);