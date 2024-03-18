// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = "{\"Schema\":\"github.com/datumforge/geodetic/internal/ent/schema\",\"Package\":\"github.com/datumforge/geodetic/internal/ent/generated\",\"Schemas\":[{\"name\":\"Database\",\"config\":{\"Table\":\"\"},\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntOAS\":{\"Create\":{\"Groups\":null,\"Policy\":0},\"Delete\":{\"Groups\":null,\"Policy\":0},\"Example\":null,\"Groups\":null,\"List\":{\"Groups\":null,\"Policy\":0},\"Read\":{\"Groups\":null,\"Policy\":0},\"ReadOnly\":true,\"Schema\":null,\"Skip\":false,\"Update\":{\"Groups\":null,\"Policy\":0}}}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntOAS\":{\"Create\":{\"Groups\":null,\"Policy\":0},\"Delete\":{\"Groups\":null,\"Policy\":0},\"Example\":null,\"Groups\":null,\"List\":{\"Groups\":null,\"Policy\":0},\"Read\":{\"Groups\":null,\"Policy\":0},\"ReadOnly\":true,\"Schema\":null,\"Skip\":false,\"Update\":{\"Groups\":null,\"Policy\":0}}}},{\"name\":\"created_by\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"immutable\":true,\"position\":{\"Index\":2,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntOAS\":{\"Create\":{\"Groups\":null,\"Policy\":0},\"Delete\":{\"Groups\":null,\"Policy\":0},\"Example\":null,\"Groups\":null,\"List\":{\"Groups\":null,\"Policy\":0},\"Read\":{\"Groups\":null,\"Policy\":0},\"ReadOnly\":true,\"Schema\":null,\"Skip\":false,\"Update\":{\"Groups\":null,\"Policy\":0}}}},{\"name\":\"updated_by\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":3,\"MixedIn\":true,\"MixinIndex\":0},\"annotations\":{\"EntOAS\":{\"Create\":{\"Groups\":null,\"Policy\":0},\"Delete\":{\"Groups\":null,\"Policy\":0},\"Example\":null,\"Groups\":null,\"List\":{\"Groups\":null,\"Policy\":0},\"Read\":{\"Groups\":null,\"Policy\":0},\"ReadOnly\":true,\"Schema\":null,\"Skip\":false,\"Update\":{\"Groups\":null,\"Policy\":0}}}},{\"name\":\"deleted_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1},\"annotations\":{\"EntGQL\":{\"Skip\":48},\"EntOAS\":{\"Create\":{\"Groups\":null,\"Policy\":0},\"Delete\":{\"Groups\":null,\"Policy\":0},\"Example\":null,\"Groups\":null,\"List\":{\"Groups\":null,\"Policy\":0},\"Read\":{\"Groups\":null,\"Policy\":0},\"ReadOnly\":true,\"Schema\":null,\"Skip\":false,\"Update\":{\"Groups\":null,\"Policy\":0}}}},{\"name\":\"deleted_by\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1},\"annotations\":{\"EntGQL\":{\"Skip\":48},\"EntOAS\":{\"Create\":{\"Groups\":null,\"Policy\":0},\"Delete\":{\"Groups\":null,\"Policy\":0},\"Example\":null,\"Groups\":null,\"List\":{\"Groups\":null,\"Policy\":0},\"Read\":{\"Groups\":null,\"Policy\":0},\"ReadOnly\":true,\"Schema\":null,\"Skip\":false,\"Update\":{\"Groups\":null,\"Policy\":0}}}},{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2},\"annotations\":{\"EntOAS\":{\"Create\":{\"Groups\":null,\"Policy\":0},\"Delete\":{\"Groups\":null,\"Policy\":0},\"Example\":null,\"Groups\":null,\"List\":{\"Groups\":null,\"Policy\":0},\"Read\":{\"Groups\":null,\"Policy\":0},\"ReadOnly\":true,\"Schema\":null,\"Skip\":false,\"Update\":{\"Groups\":null,\"Policy\":0}}}},{\"name\":\"organization_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"the ID of the organization\"},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"the name to the database\"},{\"name\":\"geo\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"the geo location of the database\"},{\"name\":\"dsn\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"the DSN to the database\"}],\"indexes\":[{\"unique\":true,\"fields\":[\"organization_id\"],\"annotations\":{\"EntSQLIndexes\":{\"Desc\":false,\"DescColumns\":null,\"IncludeColumns\":null,\"OpClass\":\"\",\"OpClassColumns\":null,\"Prefix\":0,\"PrefixColumns\":null,\"Type\":\"\",\"Types\":null,\"Where\":\"deleted_at is NULL\"}}}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}],\"annotations\":{\"EntGQL\":{\"MutationInputs\":[{\"IsCreate\":true},{}]}}}],\"Features\":[\"sql/versioned-migration\",\"privacy\",\"schema/snapshot\",\"entql\",\"namedges\",\"sql/schemaconfig\",\"intercept\",\"namedges\"]}"
