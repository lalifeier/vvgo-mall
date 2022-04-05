#!/bin/sh

protoc -I=proto/ --ent_out=. --ent_opt=schemadir=./schema proto/entpb/user.proto
