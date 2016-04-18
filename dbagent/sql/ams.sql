CREATE TABLE IF NOT EXISTS t_ams_use_id (
    f_index int NOT NULL AUTO_INCREMENT,
    f_user_id bigint unsigned NOT NULL, 
    f_create_time datetime NOT NULL,
    PRIMARY KEY(f_user_id),
    INDEX USING BTREE (f_user_id)
) ENGINE = InnoDB;


CREATE TABLE IF NOT EXISTS t_ams_open_id (
    f_index int NOT NULL AUTO_INCREMENT,
    f_platform tinyint unsigned NOT NULL,
    f_app_id char(128) ,
    f_open_id char(128) NOT NULL,
    f_user_id bigint unsigned NOT NULL, 
    f_master_user_id bigint unsigned NOT NULL, 
    PRIMARY KEY (f_platform, f_app_id, f_open_id),
    INDEX USING BTREE (f_platform, f_app_id, f_open_id)
);
