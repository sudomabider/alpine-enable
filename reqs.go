package main

var reqs = map[string]Dep{
	"git": Dep{
		sys: "git openssh-client",
	},
	"crux": Dep{
		php: "bcmath sockets",
	},
	"soap": Dep{
		sys: "libxml2-dev",
		php: "soap",
	},
	"zip": Dep{
		sys: "zlib-dev libzip-dev",
		php: "zip",
	},
	"mysql": Dep{
		sys: "mysql-client",
		php: "pdo_mysql",
	},
	"postgres": Dep{
		sys: "postgresql-dev",
		php: "pdo_pgsql",
	},
	"swoole": Dep{
		build: "$PHPIZE_DEPS",
		pecl:  "swoole",
	},
	"pm2": Dep{
		sys: "nodejs",
		npm: "pm2",
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
