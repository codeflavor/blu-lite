build:
	@echo 'Compiling blulite API...'
	@cd cmd/servops && go build -o ../../_out/bin/servops -v .

clean:
	@echo 'Removing old _out dir'
	@rm -rf ./_out

start:
	@echo 'Starting server'
	@_out/bin/servops start

swagger:
	@echo 'Generating Swagger Specs...'
	@scripts/swagger.sh

clean:
	@echo 'Cleaning build folder'
	@scripts/build.sh clean
