# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.17

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Disable VCS-based implicit rules.
% : %,v


# Disable VCS-based implicit rules.
% : RCS/%


# Disable VCS-based implicit rules.
% : RCS/%,v


# Disable VCS-based implicit rules.
% : SCCS/s.%


# Disable VCS-based implicit rules.
% : s.%


.SUFFIXES: .hpux_make_needs_suffix_list


# Command-line flag to silence nested $(MAKE).
$(VERBOSE)MAKESILENT = -s

# Suppress display of executed commands.
$(VERBOSE).SILENT:


# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /usr/local/bin/cmake

# The command to remove a file.
RM = /usr/local/bin/cmake -E rm -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent

# Include any dependencies generated for this target.
include proto/CMakeFiles/proto.dir/depend.make

# Include the progress variables for this target.
include proto/CMakeFiles/proto.dir/progress.make

# Include the compile flags for this target's objects.
include proto/CMakeFiles/proto.dir/flags.make

proto/msg.pb.h: proto/msg.proto
proto/msg.pb.h: /usr/local/bin/protoc
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --blue --bold --progress-dir=/home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Running cpp protocol buffer compiler on msg.proto"
	cd /home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent/proto && /usr/local/bin/protoc --cpp_out /home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent/proto -I /home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent/proto /home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent/proto/msg.proto

proto/msg.pb.cc: proto/msg.pb.h
	@$(CMAKE_COMMAND) -E touch_nocreate proto/msg.pb.cc

proto/CMakeFiles/proto.dir/msg.pb.cc.o: proto/CMakeFiles/proto.dir/flags.make
proto/CMakeFiles/proto.dir/msg.pb.cc.o: proto/msg.pb.cc
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Building CXX object proto/CMakeFiles/proto.dir/msg.pb.cc.o"
	cd /home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent/proto && /usr/bin/g++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/proto.dir/msg.pb.cc.o -c /home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent/proto/msg.pb.cc

proto/CMakeFiles/proto.dir/msg.pb.cc.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/proto.dir/msg.pb.cc.i"
	cd /home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent/proto && /usr/bin/g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent/proto/msg.pb.cc > CMakeFiles/proto.dir/msg.pb.cc.i

proto/CMakeFiles/proto.dir/msg.pb.cc.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/proto.dir/msg.pb.cc.s"
	cd /home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent/proto && /usr/bin/g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent/proto/msg.pb.cc -o CMakeFiles/proto.dir/msg.pb.cc.s

# Object files for target proto
proto_OBJECTS = \
"CMakeFiles/proto.dir/msg.pb.cc.o"

# External object files for target proto
proto_EXTERNAL_OBJECTS =

proto/libproto.so: proto/CMakeFiles/proto.dir/msg.pb.cc.o
proto/libproto.so: proto/CMakeFiles/proto.dir/build.make
proto/libproto.so: proto/CMakeFiles/proto.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent/CMakeFiles --progress-num=$(CMAKE_PROGRESS_3) "Linking CXX shared library libproto.so"
	cd /home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent/proto && $(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/proto.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
proto/CMakeFiles/proto.dir/build: proto/libproto.so

.PHONY : proto/CMakeFiles/proto.dir/build

proto/CMakeFiles/proto.dir/clean:
	cd /home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent/proto && $(CMAKE_COMMAND) -P CMakeFiles/proto.dir/cmake_clean.cmake
.PHONY : proto/CMakeFiles/proto.dir/clean

proto/CMakeFiles/proto.dir/depend: proto/msg.pb.h
proto/CMakeFiles/proto.dir/depend: proto/msg.pb.cc
	cd /home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent /home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent/proto /home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent /home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent/proto /home/maaz/Desktop/EC-4.20.16/EC-Agent/Agent/proto/CMakeFiles/proto.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : proto/CMakeFiles/proto.dir/depend
