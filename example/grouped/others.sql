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