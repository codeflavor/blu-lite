build :
	@echo 'Compiling blulite API...'
	@cd cmd/servops && go build -o ../../_out/bin/servops -v .

.PHONY: clean
clean :
	@echo 'Removing old _out dir'
	@rm -rf _out/

start :
	@echo 'Starting server'
	@_out/bin/servops start --workdir=.

swagger :
	@echo 'Generating Swagger Specs...'
	@scripts/swagger.sh
