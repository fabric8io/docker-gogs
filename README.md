Gogs (Go Git Service) Docker Image
==================================

This image runs Gogs with SSH & web access. You can configure Gogs via the following environment variables:

-	`DB_TYPE` - `mysql` or `postgresql`
-	`DB_HOST` - the database server to use, e.g. 127.0.0.1:3306
-	`DB_NAME` - the name of the database to use
-	`DB_USER` - the user to connect to the specified database as
-	`DB_PASSWD` - the password to connect to the database as
-	`DOMAIN` - the domain of the server (default: `gogs.fabric8.local`)
-	`TASK_INTERVAL` - the interval in minutes between webhooks being invoked (default: 0)

-	For development, you might want to disable certifcate validation for webhooks. To do this use:

-	`SKIP_TLS_VERIFY` - set to `true` to disable webhook certificate validation (default: `false`)

