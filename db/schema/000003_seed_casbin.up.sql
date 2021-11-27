BEGIN;
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/api/v1/ping/list', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/api/v1/ping/create', 'POST', NULL, NULL, NULL);
COMMIT;