

#--------------------------------------------------------------------------------
#-- 家庭账务系统 数据库创建脚本
create database if not exists fas default charset utf8 collate utf8_general_ci;
use fas;
#--------------------------------------------------------------------------------
--
-- 12.表-账本明细表
drop table if exists tbl_fas_account_items;

-- 11.表-账本关联用户表
drop table if exists tbl_fas_account_users;

-- 10.表-账本表
drop table if exists tbl_fas_accounts;

-- 9.表-通讯录与用户关联表
drop table if exists tbl_fas_user_books;

-- 8.表-通讯录
drop table if exists tbl_fas_books;

-- 6.表-用户登录流水表
drop table if exists tbl_fas_user_logins;

-- 5.表-用户第三方登录关联表
drop table if exists tbl_fas_user_oauths;

-- 2.表-系统-版本管理表
drop table if exists tbl_fas_sys_versions;
#--------------------------------------------------------------------------------
-- 1.表-系统-渠道表
drop table if exists tbl_fas_sys_channels;
create table tbl_fas_sys_channels(
  `id`      varchar(64) not null      comment '渠道ID',
  `code`    int(2) unsigned not null  comment '渠道代码',
  `name`    varchar(32) not null      comment '渠道名称',

  `status`  int(1) unsigned default 1 comment '状态(1:正常,0:停用)',

  `createTime` timestamp default current_timestamp  comment '创建时间',
  `lastTime`   timestamp default current_timestamp on update current_timestamp comment '更新时间',

  constraint pk_tbl_fas_sys_channels primary key(`id`),-- 主键约束
  constraint uk_tbl_fas_sys_channels_code unique key(`code`)-- 唯一约束
) engine=InnoDB default charset=utf8 comment '系统-渠道表';

-- 初始化
insert into tbl_fas_sys_channels(`id`,`code`,`name`) values(uuid(), 100, '管理后台');
insert into tbl_fas_sys_channels(`id`,`code`,`name`) values(uuid(), 200, 'android');
insert into tbl_fas_sys_channels(`id`,`code`,`name`) values(uuid(), 300, 'ios');
insert into tbl_fas_sys_channels(`id`,`code`,`name`) values(uuid(), 400, 'h5');
insert into tbl_fas_sys_channels(`id`,`code`,`name`) values(uuid(), 500, 'wechat');

-- 2.表-系统-版本管理表
drop table if exists tbl_fas_sys_versions;
create table tbl_fas_sys_versions(
  `id`           varchar(64) not null       comment '版本ID',
  `name`         varchar(32) not null       comment '版本名称',
  `version`      int(2) unsigned not null   comment '版本号',

  `checkCode`    varchar(64) default null   comment '校验码',
  `status`       int(1) unsigned default 1  comment '状态(1:有效,0:无效)',

  `startTime`    timestamp default current_timestamp comment '生效时间',
  `url`          varchar(1024) default null comment '下载地址',

  `description`  varchar(1024) default null comment '描述',

  `channelId`    varchar(64) not null       comment '所属渠道ID',

  `createTime`   timestamp default current_timestamp  comment '创建时间',
  `lastTime`     timestamp default current_timestamp on update current_timestamp  comment '更新时间',

  constraint pk_tbl_fas_sys_versions primary key(`id`),-- 主键约束
  constraint uk_tbl_fas_sys_versions_version unique key(`channelId`,`version`), -- 联合唯一约束
  constraint fk_tbl_fas_sys_versions_channelId foreign key(`channelId`) references tbl_fas_sys_channels(`id`)-- 外键约束
) engine=InnoDB default charset=utf8 comment '系统-版本管理表';

-- 3.表-系统-终端日志表
drop table if exists tbl_fas_client_logs;
create table tbl_fas_client_logs(
  `id`      varchar(64) not null       comment '日志ID',
  `mac`     varchar(128) default null  comment '设备标识',
  `userId`  varchar(64) default null   comment '用户ID',

  `type`    int(1) unsigned default 0  comment '类型(0:normal,1:warn, 2:error)',
  `path`    varchar(1024) default null comment '日志文件路径',


  `createTime`   timestamp default current_timestamp  comment '创建时间',
  `lastTime`     timestamp default current_timestamp on update current_timestamp  comment '更新时间',

  constraint pk_tbl_fas_client_logs primary key(`id`)
) engine=InnoDB default charset=utf8 comment '系统-终端日志表';

-- 4.表-用户表
drop table if exists tbl_fas_users;
create table tbl_fas_users(
  `id`       varchar(64)  not null        comment '用户ID',
  `account`  varchar(128) not null        comment '用户账号',
  `password` varchar(64)  default null    comment '用户密码',

  `nickName` varchar(32)   default null   comment '用户昵称',
  `iconUrl`  varchar(1024) default null   comment '头像url',

  `mobile`   varchar(20)   default null   comment '手机号码',
  `email`    varchar(255)  default null   comment '邮件地址',

  `status`   int(1) unsigned default 1    comment '状态(1:启用,0:停用)',

  `createTime`   timestamp default current_timestamp  comment '创建时间',
  `lastTime`     timestamp default current_timestamp on update current_timestamp  comment '更新时间',

  constraint pk_tbl_fas_users primary key(`id`), -- 主键约束
  constraint uk_tbl_fas_users_account unique key(`account`)-- 账号唯一约束
) engine=InnoDB default charset=utf8 comment '用户表';

-- 5.表-用户第三方登录关联表
drop table if exists tbl_fas_user_oauths;
create table tbl_fas_user_oauths(
  `id`       varchar(64)  not null      comment '关联ID',
  `userId`   varchar(64)  not null      comment '用户ID',
  `type`     int(2) unsigned default 1  comment '类型(1:wechat,2:alipay)',
  `authCode` varchar(128) not null      comment '第三方授权码',

  constraint pk_tbl_fas_user_oauths primary key(`id`),-- 主键约束
  constraint uk_tbl_fas_user_oauths_all unique key(`type`,`authCode`),-- 联合唯一约束
  constraint fk_tbl_fas_user_oauths_userId foreign key(`userId`) references tbl_fas_users(`id`)-- 外键约束
) engine=InnoDB default charset=utf8 comment '用户第三方登录表';

-- 6.表-用户登录流水表
drop table if exists tbl_fas_user_logins;
create table tbl_fas_user_logins(
  `id`        varchar(64) not null    comment '登录ID',
  `userId`    varchar(64) not null    comment '用户ID',
  `channelId` varchar(64) not null    comment '登录渠道ID',
  `method`    int(1) unsigned default 0 comment '登录方式(0:本地登录,1:微信,2:支付宝)',

  `token`   varchar(64) not null      comment '用户登录令牌',
  `ipAddr`  varchar(32) default null  comment '登录IP地址',
  `mac`     varchar(64) default null  comment '登录设备标识',

  `expiredTime` int(13) default 0    comment '过期时间戳',
  `status`  int(1) unsigned default 1 comment '状态(1:有效,0:无效)',

  `createTime`   timestamp default current_timestamp  comment '创建时间',
  `lastTime`     timestamp default current_timestamp on update current_timestamp  comment '更新时间',

  constraint pk_tbl_fas_user_logins primary key(`id`), -- 主键约束
  constraint uk_tbl_fas_user_logins_token unique key(`token`),-- 唯一约束
  constraint fk_tbl_fas_user_logins_userId foreign key(`userId`) references tbl_fas_users(`id`),-- 外键约束
  constraint fk_tbl_fas_user_logins_channelId foreign key(`channelId`) references tbl_fas_sys_channels(`id`)-- 外键约束
) engine=InnoDB default charset=utf8 comment '用户登录流水表';

-- 7.表-用户登录流水历史记录表
drop table if exists tbl_fas_user_login_histories;
create table tbl_fas_user_login_histories(
  `id`        varchar(64) not null    comment '登录ID',
  `userId`    varchar(64) not null    comment '用户ID',
  `channelId` varchar(64) not null    comment '登录渠道ID',
  `method`    int(1) unsigned default 0 comment '登录方式(0:本地登录,1:微信,2:支付宝)',

  `token`   varchar(64) not null      comment '用户登录令牌',
  `ipAddr`  varchar(32) default null  comment '登录IP地址',
  `mac`     varchar(64) default null  comment '登录设备标识',

  `createTime`   timestamp default current_timestamp  comment '创建时间',
  `lastTime`     timestamp default current_timestamp on update current_timestamp  comment '更新时间',

  constraint pk_tbl_fas_user_login_histories primary key(`id`) -- 主键约束
) engine=InnoDB default charset=utf8 comment '用户登录流水历史记录表';

-- 8.表-通讯录
drop table if exists tbl_fas_books;
create table tbl_fas_books(
  `id`           varchar(64) not null      comment '通讯录ID',
  `userId`       varchar(64) not null      comment '所属用户ID',

  `friendName`   varchar(32) default null  comment '朋友姓名',
  `friendMobile` varchar(32) default null  comment '朋友手机',

  `createTime`   timestamp default current_timestamp  comment '创建时间',
  `lastTime`     timestamp default current_timestamp on update current_timestamp  comment '更新时间',

  constraint pk_tbl_fas_books primary key(`id`),-- 主键约束
  constraint uk_tbl_fas_books_all unique key(`userId`,`friendMobile`),-- 联合唯一主键
  constraint fk_tbl_fas_books_userId foreign key(`userId`) references tbl_fas_users(`id`)-- 外键约束
) engine=InnoDB default charset=utf8 comment '通讯录';

-- 9.表-通讯录与用户关联表
drop table if exists tbl_fas_user_books;
create table tbl_fas_user_books(
  `id`      varchar(64) not null    comment '关联ID',

  `userId`  varchar(64) not null    comment '用户ID',
  `bookId`  varchar(64) not null    comment '通讯录ID',

  `createTime`   timestamp default current_timestamp  comment '创建时间',

  constraint pk_tbl_fas_user_books primary key(`id`),-- 主键约束
  constraint uk_tbl_fas_user_books_all unique key(`userId`,`bookId`),-- 联合唯一主键
  constraint fk_tbl_fas_user_books_userId foreign key(`userId`) references tbl_fas_users(`id`),-- 外键约束
  constraint fk_tbl_fas_user_books_bookId foreign key(`bookId`) references tbl_fas_books(`id`)-- 外键约束
) engine=InnoDB default charset=utf8 comment '通讯录与用户关联表';

-- 10.表-账本表
drop table if exists tbl_fas_accounts;
create table tbl_fas_accounts(
  `id`           varchar(64)  not null                     comment '账本ID',
  `code`         int(11) unsigned not null auto_increment  comment '账本代码(排序)',
  `name`         varchar(128) not null                     comment '账本名称',
  `abbr`         varchar(32) default null                  comment '账本简称',
  `type`         int(1) unsigned default 0                 comment '账本类型(0:私密账本,1:只读共享,2:公共账本)',
  `status`       int(1) unsigned default 1                 comment '状态(0:封账,1:启用,2:删除)',

  `createUserId` varchar(64) default null                  comment '创建用户',

  `createTime`   timestamp default current_timestamp       comment '创建时间',
  `lastTime`     timestamp default current_timestamp on update current_timestamp  comment '更新时间',

  constraint pk_tbl_fas_accounts primary key(`id`),-- 主键约束
  constraint uk_tbl_fas_accounts_code unique key(`code`),-- 唯一约束
  constraint uk_tbl_fas_accounts_all unique key(`createUserId`,`name`),-- 联合唯一约束
  constraint fk_tbl_fas_accounts_createUserId foreign key(`createUserId`) references tbl_fas_users(`id`)-- 外键约束
) engine=InnoDB default charset=utf8 comment '账本表';

-- 11.表-账本关联用户表
drop table if exists tbl_fas_account_users;
create table tbl_fas_account_users(
  `id`          varchar(64) not null       comment '关联ID',
  `accountId`   varchar(64) not null       comment '账本ID',
  `userId`      varchar(64) not null       comment '用户ID',
  `role`        int(1) unsigned default 0  comment '共享角色(0:所有者,1:参与者)',

  `createTime`   timestamp default current_timestamp  comment '创建时间',

  constraint pk_tbl_fas_account_users primary key(`id`),-- 主键约束
  constraint uk_tbl_fas_account_users_all unique key(`accountId`,`userId`),-- 联合唯一约束
  constraint fk_tbl_fas_account_users_accountId foreign key(`accountId`) references tbl_fas_accounts(`id`),-- 外键约束
  constraint fk_tbl_fas_account_users_userId foreign key(`userId`) references tbl_fas_users(`id`)--  外键约束
) engine=InnoDB default charset=utf8 comment '账本共享用户表';

-- 12.表-账本明细表
drop table if exists tbl_fas_account_items;
create table tbl_fas_account_items(
  `id`          varchar(64) not null        comment '明细ID',
  `accountId`   varchar(64) not null        comment '所属账本ID',
  `code`        int(11) unsigned default 0  comment '账单序号',

  `userId`      varchar(64) not null      comment '所属用户ID',

  `title`       varchar(32) not null      comment '名目',
  `money`       decimal(8,2) default 0.0  comment '金额',
  `time`        date default null         comment '时间',

  `createTime`   timestamp default current_timestamp       comment '创建时间',
  `lastTime`     timestamp default current_timestamp on update current_timestamp  comment '更新时间',

  constraint pk_tbl_fas_account_items primary key(`id`), -- 主键约束
  constraint uk_tbl_fas_account_items_code unique key(`accountId`,`code`),-- 联合唯一约束
  constraint fk_tbl_fas_account_items_accountId foreign key(`accountId`) references tbl_fas_accounts(`id`),-- 外键约束
  constraint fk_tbl_fas_account_items_userId foreign key(`userId`) references tbl_fas_users(`id`)-- 外键约束
) engine=InnoDB default charset=utf8 comment '账本明细表';

#--------------------------------------------------------------------------------
