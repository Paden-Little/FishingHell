#!/bin/bash

# Set output directory for compiled files
OUTPUT_DIR="./"
DASHBOARD_DESTINATION="../dashboard/protobufs"
FLEXIBLE_MQ_DESTINATION="../flexible-mq/protobufs"
FISHER_DESTINATION="../fisher/protobufs"
ACCOUNT_DESTINATION="../account/protobufs"
FRONTEND_DESTINATION="../frontend/FishingHellFrontend/FishingHellFrontend/Protobufs"
COMPILIED="./protobufs"

mkdir -p $OUTPUT_DIR/
mkdir -p $DASHBOARD_DESTINATION
mkdir -p $FLEXIBLE_MQ_DESTINATION
mkdir -p $FISHER_DESTINATION
mkdir -p $ACCOUNT_DESTINATION

protoc --go_out=$OUTPUT_DIR/ proto-files/fisher.proto
protoc --go_out=$OUTPUT_DIR/ proto-files/account.proto
protoc --csharp_out=$OUTPUT_DIR/ proto-files/account.proto
protoc --csharp_out=$OUTPUT_DIR/ proto-files/fisher.proto

cp $COMPILIED/fisher.pb.go $DASHBOARD_DESTINATION
cp $COMPILIED/fisher.pb.go $FLEXIBLE_MQ_DESTINATION
cp $COMPILIED/fisher.pb.go $FISHER_DESTINATION

cp $COMPILIED/account.pb.go $FLEXIBLE_MQ_DESTINATION
cp $COMPILIED/account.pb.go $ACCOUNT_DESTINATION
cp $COMPILIED/account.pb.go $DASHBOARD_DESTINATION

cp Account.cs $FRONTEND_DESTINATION
cp Fisher.cs $FRONTEND_DESTINATION
echo "Protobufs compiled successfully!"
