.PHONY: all ios ios-sim ios-bind bind clear copy-assets gosec

BIND_PACKAGES=./src ./src/engine-widgets ./src/engine-widgets/downloadableimagewidget ./src/engine-widgets/feedwidget ./src/engine-widgets/settingswidget
ASSETS=assets
GO_FLAGS=

IOS_API=-iosversion 15.0

OUT_IOS=../ios/GomobilePresentation/Frameworks/Engine.xcframework
OUT_IOS_ASSETS=../ios/GomobilePresentation/Resources/EngineAssets

TARGET=
OUT=
OUT_ASSETS=
TARGET_VER=

all:
	@echo "Use \"ios\" or \"ios-sim\" parameters"

ios: TARGET=ios
ios: ios-bind

ios-sim: TARGET=iossimulator
ios-sim: ios-bind

ios-bind: OUT=$(OUT_IOS)
ios-bind: OUT_ASSETS=$(OUT_IOS_ASSETS)
ios-bind: TARGET_VER=$(IOS_API)
ios-bind: bind

bind: clear
bind: copy-assets
bind: gosec
bind:
	@echo "Building..."
	gomobile bind -target $(TARGET) $(TARGET_VER) -o=$(OUT) $(GO_FLAGS) $(BIND_PACKAGES)
	@echo "Done!"

clear:
	@echo "Clearing..."
	@rm -drf $(OUT) && rm -drf $(OUT_ASSETS)

copy-assets:
	@echo "Copying assets..."
	@cp -R $(ASSETS)/ $(OUT_ASSETS)

gosec:
	gosec ./...