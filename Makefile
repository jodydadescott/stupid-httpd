default:
	@echo "Make what? (all, clean)"
	exit 2

all:
	cd build && $(MAKE) all
	cd docker && $(MAKE) all

clean:
	cd build && $(MAKE) clean
	cd docker && $(MAKE) clean