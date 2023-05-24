.PHONY: all build-protos engine-ios engine-ios-sim

all:
	@echo "Use \"build-protos\", \"engine-ios\" or \"engine-ios-sim\" parameters"

build-protos:
	@cd ./engine/scripts && ./build-protos.sh

engine-ios:
	@cd ./engine && make ios

engine-ios-sim:
	@cd ./engine && make ios-sim
