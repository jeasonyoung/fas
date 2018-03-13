

#--------------------------------------------------------------------------------
#-- 家庭账务系统 视图创建脚本
use fas;
#--------------------------------------------------------------------------------
-- 1.账本查询
drop view if exists vw_fas_accounts;
create view vw_fas_accounts
as
  select a.`id`,a.`code`,a.`name`,a.`abbr`,a.`type`,a.`status`,b.`userId`,b.`role`
  from tbl_fas_accounts a
  inner join tbl_fas_account_users b on b.`accountId` = a.`id`;

#--------------------------------------------------------------------------------
