install:
	@echo 'Compiling blulite API'
	@/bin/python3 scripts/install.py install

clean:
	@echo 'Removing blulite binary'
	@/bin/python3 scripts/install.py clean
