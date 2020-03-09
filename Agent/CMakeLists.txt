cmake_minimum_required(VERSION 3.12)
project(Agent)

set(CMAKE_CXX_STANDARD 17)
set(PTHREAD_LIB pthread)

include(FindProtobuf)
find_package(Protobuf REQUIRED)
add_subdirectory(proto)
include_directories(${PROTOBUF_INCLUDE_DIR})

add_executable(Agent main.cpp 
            Handler.cpp                     Handler.h 
            Server.cpp                      Server.h 
            om.h
)

target_link_libraries(Agent PRIVATE ${PTHREAD_LIB} proto ${PROTOBUF_LIBRARY})
