-- hotgo自动生成菜单权限SQL 通常情况下只在首次生成代码时自动执行一次
-- 如需再次执行请先手动删除生成的菜单权限和在SQL文件：D:\workspace\go\go-quick-hotgo\server\storage\data\generate\addons\vendor_menu.sql
-- Version: 2.11.5
-- Date: 2023-12-15 19:43:39
-- Link https://github.com/bufanyun/hotgo

SET
SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET
AUTOCOMMIT = 0;
START TRANSACTION;

--
-- 数据库： `hotgo`
--

-- --------------------------------------------------------

--
-- 插入表中的数据 `hg_admin_menu`
--


SET
@now := now();


-- 菜单目录
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`,
                             `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`,
                             `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`,
                             `created_at`, `updated_at`)
VALUES (NULL, '2321', '供应商检索', 'vendor', '/vendor', 'MenuOutlined', '1', '', '', '', 'ParentLayout', '1', '', '0',
        '0', '', '0', '0', '0', '1', '', '0', '', '1', @now, @now);


SET
@dirId = LAST_INSERT_ID();


-- 菜单页面
-- 列表
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`,
                             `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`,
                             `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`,
                             `created_at`, `updated_at`)
VALUES (NULL, @dirId, '供应商检索列表', 'vendorIndex', 'index', '', '2', '', '/supplier_search/vendor/list', '',
        '/addons/supplier_search/vendor/index', '1', '', '0', '0', '', '0', '0', '0', '2', '', '10', '', '1', @now,
        @now);


SET
@listId = LAST_INSERT_ID();

-- 详情
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`,
                             `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`,
                             `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`,
                             `created_at`, `updated_at`)
VALUES (NULL, @dirId, '供应商检索详情', 'vendorView', 'view/:id?', '', '2', '', '/supplier_search/vendor/view', '',
        '/addons/supplier_search/vendor/view', '0', 'vendorIndex', '0', '0', '', '0', '1', '0', '2', '', '20', '', '1',
        @now, @now);


-- 菜单按钮

-- 编辑
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`,
                             `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`,
                             `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`,
                             `created_at`, `updated_at`)
VALUES (NULL, @listId, '编辑/新增供应商检索', 'vendorEdit', '', '', '3', '', '/supplier_search/vendor/edit', '', '',
        '1', '', '0', '0', '', '0', '1', '0', '3', '', '10', '', '1', @now, @now);


SET
@editId = LAST_INSERT_ID();


-- 删除
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`,
                             `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`,
                             `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`,
                             `created_at`, `updated_at`)
VALUES (NULL, @listId, '删除供应商检索', 'vendorDelete', '', '', '3', '', '/supplier_search/vendor/delete', '', '', '1',
        '', '0', '0', '', '0', '0', '0', '3', '', '10', '', '1', @now, @now);


-- 导出
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`,
                             `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`,
                             `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`,
                             `created_at`, `updated_at`)
VALUES (NULL, @listId, '导出供应商检索', 'vendorExport', '', '', '3', '', '/supplier_search/vendor/export', '', '', '1',
        '', '0', '0', '', '0', '0', '0', '3', '', '10', '', '1', @now, @now);


COMMIT;