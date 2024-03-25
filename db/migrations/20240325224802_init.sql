-- Create "databases" table
CREATE TABLE `databases` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `organization_id` text NOT NULL, `name` text NOT NULL, `geo` text NULL, `dsn` text NOT NULL, `token` text NULL, `status` text NOT NULL DEFAULT ('CREATING'), `provider` text NOT NULL DEFAULT ('LOCAL'), `group_id` text NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `databases_groups_databases` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE NO ACTION);
-- Create index "database_organization_id" to table: "databases"
CREATE UNIQUE INDEX `database_organization_id` ON `databases` (`organization_id`) WHERE deleted_at is NULL;
-- Create "groups" table
CREATE TABLE `groups` (`id` text NOT NULL, `created_at` datetime NULL, `updated_at` datetime NULL, `created_by` text NULL, `updated_by` text NULL, `deleted_at` datetime NULL, `deleted_by` text NULL, `name` text NOT NULL, `description` text NULL, `primary_location` text NULL, `locations` json NULL, `token` text NULL, `region` text NOT NULL DEFAULT ('AMER'), PRIMARY KEY (`id`));
