The `enable` binary allows one to easily install the most common functionalities via `enable`able modules, e.g. `mysql` and `zip`

## Usage

```bash
$ enable -h

Usage: enable OPTIONS [args...]

Options:
  --all|-a          Enable all available modules
  --version|-v      Print version
  --help|-h         Print usage
  --dry|-d          Print the command but not execute

args: [crux gd git imap mysql pcov pm2 postgres soap swoole xdebug xmlrpc xsl zip]
```

```bash
$ enable -d zip soap postgres swoole

[Command]
apk add --no-cache libxml2-dev libzip-dev postgresql-dev zlib-dev && \
apk add --no-cache --virtual .build $PHPIZE_DEPS && \
pecl install swoole && \
docker-php-ext-enable swoole && \
docker-php-ext-install pcntl pdo_pgsql pgsql soap zip && \
apk del .build
```

## Manage modules

The module recipes are defined in `recipe/recipes`. Note there are different recipes for different environments.
