package main

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
		php:    []string{"pdo_mysql"},
	},
	"postgres": dep{
		system: []string{"postgresql-dev"},
		php:    []string{"pdo_pgsql"},
	},
	"swoole": dep{
		build: []string{"$PHPIZE_DEPS"},
		php:   []string{"pcntl"},
		pecl:  []string{"swoole"},
	},
	"xdebug": dep{
		build: []string{"$PHPIZE_DEPS"},
		pecl:  []string{"xdebug"},
	},
	"pcov": dep{
		build: []string{"$PHPIZE_DEPS"},
		pecl:  []string{"pcov"},
	},
	"pm2": dep{
		system: []string{"nodejs", "npm"},
		npm:    []string{"pm2"},
	},
}
