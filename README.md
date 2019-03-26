```
$ enable -h

Usage: enable OPTIONS [args...]

Options:
    -version|v         Print version
    -help|h            Print usage
    -dry|d             Print the command but not execute

args: [git crux soap mysql swoole xdebug pm2 xsl mcrypt zip postgres pcov]

$ enable -d zip soap postgres swoole

[Command]
apk add --no-cache libxml2-dev libzip-dev postgresql-dev zlib-dev && \
apk add --no-cache --virtual .build $PHPIZE_DEPS && \
pecl install swoole && \
docker-php-ext-enable swoole && \
docker-php-ext-install pcntl pdo_pgsql pgsql soap zip && \
apk del .build
```
