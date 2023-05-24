#!/bin/sh

PROTOS_DIR=../protos
OUT_IOS_PROTOS_DIR=../../ios/GomobilePresentation/Protos
OUT_GO_ENGINE_PROTOS_DIR=../src/internal

echo "Cleaning..."
rm -drf $OUT_IOS_PROTOS_DIR
mkdir -p $OUT_IOS_PROTOS_DIR
rm -drf $OUT_GO_ENGINE_PROTOS_DIR/protos
mkdir -p $OUT_GO_ENGINE_PROTOS_DIR

echo "Building swift files..."
protoc --proto_path=$PROTOS_DIR --swift_opt=FileNaming=DropPath --swift_out=$OUT_IOS_PROTOS_DIR ./$(find $PROTOS_DIR -iname "*.proto")

echo "Building go files..."
protoc --proto_path=$PROTOS_DIR --go_out=$OUT_GO_ENGINE_PROTOS_DIR ./$(find $PROTOS_DIR -iname "*.proto")

echo "Done!"