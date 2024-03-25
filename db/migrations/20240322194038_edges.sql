-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_databases" table
CREATE TABLE `new_databases` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `organization_id` text NOT NULL, `name` text NOT NULL, `geo` text NULL, `dsn` text NOT NULL, `token` text NULL, `status` text NOT NULL DEFAULT ('CREATING'), `provider` text NOT NULL DEFAULT ('LOCAL'), `group_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `databases_groups_databases` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE NO ACTION);
-- Copy rows from old table "databases" to new temporary table "new_databases"
INSERT INTO `new_databases` (`id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `organization_id`, `name`, `geo`, `dsn`, `token`, `status`, `provider`) SELECT `id`, `created_at`, `updated_at`, `created_by`, `updated_by`, `deleted_at`, `deleted_by`, `organization_id`, `name`, `geo`, `dsn`, `token`, `status`, `provider` FROM `databases`;
-- Drop "databases" table after copying rows
DROP TABLE `databases`;
-- Rename temporary table "new_databases" to "databases"
ALTER TABLE `new_databases` RENAME TO `databases`;
-- Create index "database_organization_id" to table: "databases"
CREATE UNIQUE INDEX `database_organization_id` ON `databases` (`organization_id`) WHERE deleted_at is NULL;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
