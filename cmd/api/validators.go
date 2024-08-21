package main

func serverIsValid(cfg config) (bool, map[string]interface{}) {
	errs := make(map[string]interface{})

	// Check if cfg map is correct since that's the only default set by environment variables
	// cfg
	if cfg.smtp.host == "" {
		errs["host"] = "host is required"
	}

	if cfg.smtp.port == 0 {
		errs["port"] = "port is required"
	}

	if cfg.smtp.username == "" {
		errs["username"] = "username is required"
	}

	if cfg.smtp.password == "" {
		errs["password"] = "password is required"
	}

	if cfg.db.dsn == "" {
		errs["dsn"] = "db dsn is required"
	}

	return len(errs) == 0, errs
}
