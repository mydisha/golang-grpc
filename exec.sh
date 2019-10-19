#!/bin/bash

protoc greet/model/greet.proto --go_out=plugins=grpc:.
protoc greet/model/sum.proto --go_out=plugins=grpc:.