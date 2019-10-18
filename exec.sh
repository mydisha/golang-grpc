#!/bin/bash

protoc greet/model/greet.proto --go_out=plugins=grpc:.