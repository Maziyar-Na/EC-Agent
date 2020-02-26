cmake_minimum_required(VERSION 3.15)
project(Agent_proto)

set(CMAKE_STANDARD 17)

find_program(PROTOBUF_PROTOC protoc)

include(FindProtobuf)
find_package(Protobuf REQUIRED)
include_directories(${PROTOBUF_INCLUDE_DIR})
protobuf_generate_cpp(PROTO_SRC PROTO_HEADER msg.proto)
add_library(proto SHARED ${PROTO_HEADER} ${PROTO_SRC})