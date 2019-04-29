package recipe

// Recipe defines how a module should be enabled
type Recipe map[string]dep

var recipe = Recipe{
	"git": dep{
		system: []string{"git", "openssh-client"},
	},
	"crux": dep{
		php: []string{"bcmath", "sockets"},
	},
	"soap": dep{
		system: []string{"libxml2-dev"},
		php:    []string{"soap"},
	},
	"zip": dep{
		system: []string{"zlib-dev", "libzip-dev"},
		php:    []string{"zip"},
	},
	"mysql": dep{
		system: []string{"mysql-client"},
		php:    []string{"pdo_mysql", "mysql"},
	},
	"postgres": dep{
		system: []string{"postgresql-dev"},
		php:    []string{"pgsql", "pdo_pgsql"},
	},
	"swoole": dep{
		build: []string{"$PHPIZE_DEPS"},
		php:   []string{"pcntl"},
		pecl:  []string{"swoole"},
	},
	"xdebug": dep{
		build: []string{"$PHPIZE_DEPS"},
		pecl:  []string{"xdebug"},
		post: `xdebug_ini=$PHP_INI_DIR/conf.d/docker-php-ext-xdebug.ini && \
echo "xdebug.remote_port=9000" >> ${xdebug_ini} && \
echo "xdebug.remote_enable=1" >> ${xdebug_ini} && \
echo "xdebug.remote_autostart=1" >> ${xdebug_ini} && \
echo "xdebug.remote_host=${XDEBUG_REMOTE_HOST:-host.docker.internal}" >> ${xdebug_ini}`,
	},
	"pcov": dep{
		build: []string{"$PHPIZE_DEPS"},
		pecl:  []string{"pcov"},
	},
	"pm2": dep{
		system: []string{"nodejs", "npm"},
		npm:    []string{"pm2"},
	},
	"xsl": dep{
		system: []string{"libxslt-dev"},
		php:    []string{"xsl"},
	},
	"mcrypt": dep{
		system: []string{"libmcrypt-dev"},
		php:    []string{"mcrypt"},
	},
}
