cmake_minimum_required(VERSION 3.15)
project(Agent)

set(CMAKE_CXX_STANDARD 17)
set(PTHREAD_LIB pthread)

include_directories(.)

add_executable(Agent
        CMakeLists.txt
        Handler.cpp
        Handler.h
        main.cpp
        Server.cpp
        Server.h)

target_link_libraries(Agent PRIVATE ${PTHREAD_LIB})