Gogs (Go Git Service) Docker Image
==================================

Latest version: 0.6.1 ([fabric8/gogs:0.6.1](https://registry.hub.docker.com/u/fabric8/gogs/))

This image runs Gogs with SSH & web access.

To configure Gogs, you can either mount a config file at `/opt/gogs/custom/conf/app.ini`
or use environment variables to configure all options. Specifying a config file
will prevent environment variables from taking effect: options are not merged.

All options in the [config file](http://gogs.io/docs/advanced/configuration_cheat_sheet.html)
are configurable via environment variables using a special format for the name:

```
GOGS_SECTION_NAME__KEY_NAME
```

All environment variables prefixed with `GOGS_` will be used to create the
`/opt/gogs/custom/conf/app.ini` file. The section name is optional, but the
key name is required. Notice the double underscore between `SECTION_NAME` &
`KEY_NAME`. For example, to override the user to run Gogs as
(`RUN_USER`) you can specify:

```
GOGS_RUN_USER=git
```

Or to override the database type (`DB_TYPE` key in `[database]` section) you can
specify:

```
GOGS_DATABASE__DB_TYPE=mysql
```

If a section contains a `.` (invalid env var for bash) then replace `.` with an `_`,
e.g. use `GOGS_OAUTH_GOOGLE__ENABLED` for the `oauth.google` section `ENABLED` key.

Reference config links:

- No section (prefix: `GOGS_*`): http://gogs.io/docs/advanced/configuration_cheat_sheet.html#overall
- Repository section (prefix: `GOGS_REPOSITORY__*`): http://gogs.io/docs/advanced/configuration_cheat_sheet.html#repository
- Server section (prefix: `GOGS_SERVER__*`): http://gogs.io/docs/advanced/configuration_cheat_sheet.html#server
- Database section (prefix: `GOGS_DATABASE__*`): http://gogs.io/docs/advanced/configuration_cheat_sheet.html#database
- Security section (prefix: `GOGS_SECURITY__*`): http://gogs.io/docs/advanced/configuration_cheat_sheet.html#security
- Service section (prefix: `GOGS_SERVICE__*`): http://gogs.io/docs/advanced/configuration_cheat_sheet.html#service
- Webhook section (prefix: `GOGS_WEBHOOK__*`): http://gogs.io/docs/advanced/configuration_cheat_sheet.html#webhook
- Mailer section (prefix: `GOGS_MAILER__*`): http://gogs.io/docs/advanced/configuration_cheat_sheet.html#mailer
- OAuth section (prefix: `GOGS_OAUTH__*`): http://gogs.io/docs/advanced/configuration_cheat_sheet.html#oauth
- Cache section (prefix: `GOGS_CACHE__*`): http://gogs.io/docs/advanced/configuration_cheat_sheet.html#cache
- Session section (prefix: `GOGS_SESSION__*`): http://gogs.io/docs/advanced/configuration_cheat_sheet.html#session
- Picture section (prefix: `GOGS_PICTURE__*`): http://gogs.io/docs/advanced/configuration_cheat_sheet.html#picture
- Log section (prefix: `GOGS_LOG__*`): http://gogs.io/docs/advanced/configuration_cheat_sheet.html#log
