package recipe

import (
    "fmt"
)

var baseRecipe = Recipe{
	"git": Def{
		System: []string{"git", "openssh-client"},
	},
	"crux": Def{
		Php: []string{"bcmath", "sockets"},
	},
	"soap": Def{
		System: []string{"libxml2-dev"},
		Php:    []string{"soap"},
	},
	"zip": Def{
		System: []string{"zlib-dev", "libzip-dev"},
		Php:    []string{"zip"},
	},
	"xdebug": Def{
		Build: []string{"$PHPIZE_DEPS"},
		Pecl:  []string{"xdebug"},
		Post: `xdebug_ini=$PHP_INI_DIR/conf.d/docker-php-ext-xdebug.ini && \
echo "xdebug.remote_port=9000" >> ${xdebug_ini} && \
echo "xdebug.remote_enable=1" >> ${xdebug_ini} && \
echo "xdebug.remote_autostart=1" >> ${xdebug_ini} && \
echo "xdebug.remote_host=${XDEBUG_REMOTE_HOST:-host.docker.internal}" >> ${xdebug_ini}`,
	},
	"pm2": Def{
		System: []string{"nodejs", "npm"},
		Npm:    []string{"pm2"},
	},
	"xsl": Def{
		System: []string{"libxslt-dev"},
		Php:    []string{"xsl"},
	},
}

var php5Recipe = Recipe{
    "mysql": Def{
        System: []string{"mysql-client"},
        Php:    []string{"pdo_mysql", "mysql", "mysqli"},
    },
    "postgres": Def{
        System: []string{"postgresql-dev"},
        Php:    []string{"pgsql", "pdo_pgsql"},
    },
    "mcrypt": Def{
        System: []string{"libmcrypt-dev"},
        Php:    []string{"mcrypt"},
    },
    "xdebug": Def{
        Build: []string{"$PHPIZE_DEPS"},
        Pecl:  []string{"xdebug"},
        Post: `xdebug_ini=$PHP_INI_DIR/conf.d/docker-php-ext-xdebug.ini && \
echo "xdebug.remote_port=9000" >> ${xdebug_ini} && \
echo "xdebug.remote_enable=1" >> ${xdebug_ini} && \
echo "xdebug.remote_autostart=1" >> ${xdebug_ini} && \
echo "xdebug.remote_host=${XDEBUG_REMOTE_HOST:-host.docker.internal}" >> ${xdebug_ini}`,
        Version: "2.5.5",
    },
}

var php7Recipe = Recipe{
    "mysql": Def{
        System: []string{"mysql-client"},
        Php:    []string{"pdo_mysql"},
    },
    "postgres": Def{
        System: []string{"postgresql-dev"},
        Php:    []string{"pdo_pgsql"},
    },
    "pcov": Def{
        Build: []string{"$PHPIZE_DEPS"},
        Pecl:  []string{"pcov"},
    },
    "gd": Def{
        System: []string{"libpng-dev"},
        Php:    []string{"gd"},
    },
    "imap": Def{
        System: []string{"imap-dev"},
        Php:    []string{"imap"},
    },
    "swoole": Def{
        Build: []string{"$PHPIZE_DEPS"},
        Php:   []string{"pcntl"},
        Pecl:  []string{"swoole"},
    },
    "xmlrpc": Def{
        System: []string{"libxml2-dev"},
        Php:    []string{"xmlrpc"},
    },
}

func combine(rs ...Recipe) Recipe {
    combined := Recipe{}

    for _, r := range rs {
        for k, d := range r {
            combined[k] = d
        }
    }

    return combined
}

func GetPHPRecipe(ver string) (Recipe, error) {
    if ver == "7" {
        return combine(baseRecipe, php7Recipe), nil
    }

    if ver == "5" {
        return combine(baseRecipe, php5Recipe), nil
    }

    return Recipe{}, fmt.Errorf("no recipes are found for PHP version [%s]", ver)
}
