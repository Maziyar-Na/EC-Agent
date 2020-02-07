cmake_minimum_required(VERSION 3.14)
project(Agent)

set(CMAKE_CXX_STANDARD 17)
set(PTHREAD_LIB pthread)

include(FindProtobuf)
find_package(Protobuf REQUIRED)
add_subdirectory(proto)
include_directories(${PROTOBUF_INCLUDE_DIR})

add_executable(Agent Handler.cpp Handler.h Server.cpp Server.h om.h main.cpp)

target_link_libraries(Agent ${PTHREAD_LIB} proto ${PROTOBUF_LIBRARY})
