server:
	$(MAKE) -C mqtt_test
	./go_build.bsh mqtt_to_db.go
	./docker_build.bsh

all: server
