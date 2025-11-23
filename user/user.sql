CREATE TABLE user_test (
    id bigint AUTO_INCREMENT,
    name varchar(255) NULL COMMENT 'The username',
    password varchar(255) NOT NULL DEFAULT '' COMMENT 'The user password',
    phone varchar(255) NOT NULL DEFAULT '' COMMENT 'The mobile phone number',
    gender char(10) NOT NULL DEFAULT 'male' COMMENT 'gender,malelfemalelunknown',
    nickname varchar(255) NULL DEFAULT '' COMMENT 'The nickname',
    type tinyint(1) NULL DEFAULT 0 COMMENT 'The user type,O:normal,1:vip, for test golang keyword',
    create_at timestamp NULL,
    update_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE phone_index (phone),
    UNIQUE name_index (name),
    PRIMARY KEY (id)
)ENGINE=InnoDB COLLATE utf8mb4_general_ci COMMENT '测试表';