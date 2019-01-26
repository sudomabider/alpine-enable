```
$ enable -h

Supported arguments: [git mysql postgres xdebug crux soap zip swoole pm2]

Options:
  -d    Print the full command but not execute

$ enable -d zip soap postgres swoole

Command [apk add --no-cache libxml2-dev libzip-dev postgresql-dev zlib-dev && apk add --no-cache --virtual .build $PHPIZE_DEPS && pecl install swoole && docker-php-ext-enable swoole && docker-php-ext-install pcntl pdo_pgsql soap zip && apk del .build]
```
