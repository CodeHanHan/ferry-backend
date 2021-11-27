BEGIN;
DELETE FROM `casbin_rule` WHERE `ptype` = "p";
COMMIT;