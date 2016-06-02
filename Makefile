build:
	@echo 'Compiling blulite API...'
	@/bin/python3 scripts/install.py build

swagger:
	@echo 'Generating Swagger Specs...'
	@/bin/python3 scripts/swagger.py

clean:
	@echo 'Cleaning build folder'
	@/bin/python3 scripts/install.py clean
