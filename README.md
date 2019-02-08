```
$ enable -h

Supported arguments: [git mysql postgres xdebug crux soap zip swoole pm2]

Options:
  -d    Print the full command but not execute
  -v    Specific PECL package version

$ enable -d zip soap postgres swoole

Command [apk add --no-cache libxml2-dev libzip-dev postgresql-dev zlib-dev && apk add --no-cache --virtual .build $PHPIZE_DEPS && pecl install swoole && docker-php-ext-enable swoole && docker-php-ext-install pcntl pdo_pgsql soap zip && apk del .build]

# enable -d -version 2.7.0RC1 xdebug

Command [apk add --no-cache --virtual .build $PHPIZE_DEPS && pecl install xdebug-2.7.0RC1 && docker-php-ext-enable xdebug && apk del .build]

```
