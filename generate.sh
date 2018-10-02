#!/usr/bin/env bash

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
protoc primefactor/primefactorpb/primefactor.proto --go_out=plugins=grpc:.
protoc findmax/findmaxpb/findmax.proto --go_out=plugins=grpc:.
