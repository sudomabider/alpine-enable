package main

var reqs = map[string]Dep{
	"git": Dep{
		sys: []string{"git", "openssh-client"},
	},
	"crux": Dep{
		php: []string{"bcmath", "sockets"},
	},
	"soap": Dep{
		sys: []string{"libxml2-dev"},
		php: []string{"soap"},
	},
	"zip": Dep{
		sys: []string{"zlib-dev libzip-dev"},
		php: []string{"zip"},
	},
	"mysql": Dep{
		sys: []string{"mysql-client"},
		php: []string{"pdo_mysql"},
	},
	"postgres": Dep{
		sys: []string{"postgresql-dev"},
		php: []string{"pdo_pgsql"},
	},
	"swoole": Dep{
		build: []string{"$PHPIZE_DEPS"},
		pecl:  []string{"swoole"},
	},
	"xdebug": Dep{
		build: []string{"$PHPIZE_DEPS"},
		pecl:  []string{"xdebug"},
	},
	"pm2": Dep{
		sys: []string{"nodejs"},
		npm: []string{"pm2"},
	},
}

// SupportedModules returns a list of supported modules
func SupportedModules() []string {
	keys := make([]string, 0, len(reqs))
	for k := range reqs {
		keys = append(keys, k)
	}
	return keys
}
