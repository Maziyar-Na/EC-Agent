cmake_minimum_required(VERSION 3.12)
project(Agent)

set(CMAKE_CXX_STANDARD 17)
set(PTHREAD_LIB pthread)

add_executable(Agent Handler.cpp Handler.h Server.cpp Server.h om.h main.cpp)

target_link_libraries(Agent ${PTHREAD_LIB})
