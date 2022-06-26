package models

//CREATE TABLE `im_users` (
//`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
//`name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
//`email` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
//`email_verified_at` timestamp NULL DEFAULT NULL,
//`password` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
//`remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
//`created_at` timestamp NULL DEFAULT NULL,
//`updated_at` timestamp NULL DEFAULT NULL,
//`avatar` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '头像',
//`oauth_id` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '第三方id',
//`bound_oauth` tinyint(1) DEFAULT '0' COMMENT '1\\github 2\\gitee',
//`deleted_at` timestamp NULL DEFAULT NULL,
//`oauth_type` tinyint(1) DEFAULT NULL COMMENT '1.微博 2.github',
//`status` tinyint(1) DEFAULT '0' COMMENT '0 离线 1 在线',
//`bio` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户简介',
//`sex` tinyint(1) DEFAULT '0' COMMENT '0 未知 1.男 2.女',
//`client_type` tinyint(1) DEFAULT NULL COMMENT '1.web 2.pc 3.app',
//`age` int(3) DEFAULT NULL,
//`last_login_time` timestamp NULL DEFAULT NULL COMMENT '最后登录时间',
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

type User struct {
	Model
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
}

// TableName 会将 User 的表名重写为 `profiles`
func (User) TableName() string {
	return "im_users"
}
