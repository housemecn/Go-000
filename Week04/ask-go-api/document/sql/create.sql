-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `user_no`      bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '用户ID',
    `nickname`     varchar(64)         NOT NULL DEFAULT '' COMMENT '昵称',
    `user_name`    varchar(32)         NOT NULL DEFAULT '' COMMENT '用户名，用户唯一标识',
    `user_avatar`  varchar(255)        NOT NULL DEFAULT '' COMMENT '头像',
    `user_sex`     tinyint(4) unsigned NOT NULL DEFAULT 0 COMMENT '用户性别 10 男 20女 30 未知保密',
    `encrypt_code` varchar(64)         NOT NULL DEFAULT '' COMMENT '用户加密标识',
    `reg_from`     varchar(32)         NOT NULL DEFAULT '' COMMENT '邀请人的唯一标识 invite_code',
    `reg_ip`       varchar(64)                  DEFAULT '' COMMENT '注册IP',
    `reg_time`     datetime(6)                  DEFAULT current_timestamp(6) COMMENT '注册时间',
    `reg_type`     tinyint(4) unsigned NOT NULL DEFAULT 0 COMMENT '用户注册类型 10小程序 20微信 30app',
    `reg_os`       tinyint(4) unsigned NOT NULL DEFAULT 0 COMMENT '用户设备系统 10iOS、20Android、30Windows、40macOS、50其他',
    `user_level`   tinyint(4) unsigned NOT NULL DEFAULT 0 COMMENT '用户登记',
    `auth_salt`    varchar(64)         NOT NULL DEFAULT '' COMMENT '设置密码加盐值',
    `mobile`       varchar(32)                  DEFAULT '' COMMENT '用户绑定手机号',
    `mb_time`      bigint(20) unsigned          DEFAULT 0 COMMENT '用户绑定手机号时间',
    `mb_state`     tinyint(4) unsigned          DEFAULT 0 COMMENT '手机号绑定 0 默认 100已绑定',
    `state`        tinyint(4) unsigned NOT NULL DEFAULT 0 COMMENT '状态 100默认正常 90代审核 120禁用',
    `create_time`  datetime(6)         NOT NULL DEFAULT current_timestamp(6) COMMENT '创建时间',
    `modify_time`  datetime                     DEFAULT current_timestamp() COMMENT '修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uidx_uno` (`user_no`) USING BTREE,
    UNIQUE KEY `uidx_un` (`user_name`) USING BTREE,
    UNIQUE KEY `uidx_ec` (`encrypt_code`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户信息表';

-- ----------------------------
-- Table structure for user_auth
-- ----------------------------
DROP TABLE IF EXISTS `user_auth`;
CREATE TABLE `user_auth`
(
    `id`            bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `user_no`       bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '关联用户ID',
    `auth_account`  varchar(64)         NOT NULL DEFAULT '' COMMENT '登录ID',
    `auth_password` varchar(255)        NOT NULL DEFAULT '' COMMENT '登录密码',
    `auth_type`     varchar(24)         NOT NULL DEFAULT '' COMMENT '账号类型 wechat、welite、app',
    `login_ip`      varchar(64)                  DEFAULT '' COMMENT '最近登录IP',
    `login_time`    datetime(6)                  DEFAULT current_timestamp(6) COMMENT '最近登录时间',
    `auth_state`    tinyint(4) unsigned NOT NULL DEFAULT 100 COMMENT '状态 默认100，100正常，90禁用',
    `create_time`   datetime(6)         NOT NULL DEFAULT current_timestamp(6) COMMENT '创建时间',
    `modify_time`   datetime            NOT NULL DEFAULT current_timestamp() COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uidx_ut` (`user_no`, `auth_type`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 445
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户登录授权表';