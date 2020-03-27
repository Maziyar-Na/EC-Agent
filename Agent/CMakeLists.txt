cmake_minimum_required(VERSION 3.15)
project(Agent)

set(CMAKE_CXX_STANDARD 17)
set(PTHREAD_LIB pthread)
include(FindProtobuf)
find_package(Protobuf REQUIRED)
add_subdirectory(proto)
include_directories(${PROTOBUF_INCLUDE_DIR})

add_custom_command(
        OUTPUT
            ${CMAKE_CURRENT_LIST_DIR}/proto/msg.pb.h
            ${CMAKE_CURRENT_LIST_DIR}/proto/msg.pb.cc
        MAIN_DEPENDENCY ${CMAKE_CURRENT_LIST_DIR}/proto/msg.proto
        COMMAND ${PROTOBUF_PROTOC}
        ARGS --proto_path=. --cpp_out=. msg.proto
        WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR}/proto)

set(AGENT_PROTO_SRC
    proto/msg.pb.h
    proto/msg.pb.cc)

add_executable(Agent
        Handler.cpp Handler.h
        Server.cpp Server.h
        om.h main.cpp  ${AGENT_PROTO_SRC})

target_link_libraries(Agent ${PTHREAD_LIB} proto ${PROTOBUF_LIBRARY})

