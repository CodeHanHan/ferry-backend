BEGIN;
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/api/v1/ping/list', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/api/v1/ping/create', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/api/v1/ping/delete', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/api/v1/ping/update', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/api/v1/user/me', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/api/v1/user', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/api/v1/user/login', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/api/v1/user', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/api/v1/user/logintest', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/api/v1/user/adminupdateuser', 'PATCH', NULL, NULL, NULL);
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/api/v1/role', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/api/v1/role/*', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/api/v1/role', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/api/v1/role/*', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/api/v1/role', 'PUT', NULL, NULL, NULL);
COMMIT;

BEGIN;
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'common', '/api/v1/ping/delete', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'common', '/api/v1/ping/update', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'common', '/api/v1/user/login', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'common', '/api/v1/user/logintest', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'common', '/api/v1/user/updateuser', 'PATCH', NULL, NULL, NULL);
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'common', '/api/v1/user/changepassword', 'PATCH', NULL, NULL, NULL);
COMMIT;