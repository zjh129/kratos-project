include ../../app_makefile

.PHONY: entimport
entimport:
	cd internal/data/ && \
    go run -mod=mod ariga.io/entimport/cmd/entimport -dsn "mysql://root:root@tcp(127.0.0.1:3306)/kratos_project"