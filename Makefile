build:
	@echo 'Compiling blulite API...'
	@scripts/build.sh build

swagger:
	@echo 'Generating Swagger Specs...'
	@scripts/swagger.sh

clean:
	@echo 'Cleaning build folder'
	@scripts/build.sh clean
