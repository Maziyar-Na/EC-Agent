# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.14

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Remove some rules from gmake that .SUFFIXES does not remove.
SUFFIXES =

.SUFFIXES: .hpux_make_needs_suffix_list


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
CMAKE_COMMAND = /opt/clion-2019.2.2/bin/cmake/linux/bin/cmake

# The command to remove a file.
RM = /opt/clion-2019.2.2/bin/cmake/linux/bin/cmake -E remove -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/maaz/CLionProjects/Agent

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/maaz/CLionProjects/Agent/cmake-build-debug

# Include any dependencies generated for this target.
include CMakeFiles/Agent.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/Agent.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/Agent.dir/flags.make

CMakeFiles/Agent.dir/Handler.cpp.o: CMakeFiles/Agent.dir/flags.make
CMakeFiles/Agent.dir/Handler.cpp.o: ../Handler.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/maaz/CLionProjects/Agent/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building CXX object CMakeFiles/Agent.dir/Handler.cpp.o"
	/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/Agent.dir/Handler.cpp.o -c /home/maaz/CLionProjects/Agent/Handler.cpp

CMakeFiles/Agent.dir/Handler.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/Agent.dir/Handler.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/maaz/CLionProjects/Agent/Handler.cpp > CMakeFiles/Agent.dir/Handler.cpp.i

CMakeFiles/Agent.dir/Handler.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/Agent.dir/Handler.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/maaz/CLionProjects/Agent/Handler.cpp -o CMakeFiles/Agent.dir/Handler.cpp.s

CMakeFiles/Agent.dir/Server.cpp.o: CMakeFiles/Agent.dir/flags.make
CMakeFiles/Agent.dir/Server.cpp.o: ../Server.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/maaz/CLionProjects/Agent/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Building CXX object CMakeFiles/Agent.dir/Server.cpp.o"
	/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/Agent.dir/Server.cpp.o -c /home/maaz/CLionProjects/Agent/Server.cpp

CMakeFiles/Agent.dir/Server.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/Agent.dir/Server.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/maaz/CLionProjects/Agent/Server.cpp > CMakeFiles/Agent.dir/Server.cpp.i

CMakeFiles/Agent.dir/Server.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/Agent.dir/Server.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/maaz/CLionProjects/Agent/Server.cpp -o CMakeFiles/Agent.dir/Server.cpp.s

CMakeFiles/Agent.dir/main.cpp.o: CMakeFiles/Agent.dir/flags.make
CMakeFiles/Agent.dir/main.cpp.o: ../main.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/maaz/CLionProjects/Agent/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_3) "Building CXX object CMakeFiles/Agent.dir/main.cpp.o"
	/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/Agent.dir/main.cpp.o -c /home/maaz/CLionProjects/Agent/main.cpp

CMakeFiles/Agent.dir/main.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/Agent.dir/main.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/maaz/CLionProjects/Agent/main.cpp > CMakeFiles/Agent.dir/main.cpp.i

CMakeFiles/Agent.dir/main.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/Agent.dir/main.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/maaz/CLionProjects/Agent/main.cpp -o CMakeFiles/Agent.dir/main.cpp.s

# Object files for target Agent
Agent_OBJECTS = \
"CMakeFiles/Agent.dir/Handler.cpp.o" \
"CMakeFiles/Agent.dir/Server.cpp.o" \
"CMakeFiles/Agent.dir/main.cpp.o"

# External object files for target Agent
Agent_EXTERNAL_OBJECTS =

Agent: CMakeFiles/Agent.dir/Handler.cpp.o
Agent: CMakeFiles/Agent.dir/Server.cpp.o
Agent: CMakeFiles/Agent.dir/main.cpp.o
Agent: CMakeFiles/Agent.dir/build.make
Agent: CMakeFiles/Agent.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/maaz/CLionProjects/Agent/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_4) "Linking CXX executable Agent"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/Agent.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/Agent.dir/build: Agent

.PHONY : CMakeFiles/Agent.dir/build

CMakeFiles/Agent.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/Agent.dir/cmake_clean.cmake
.PHONY : CMakeFiles/Agent.dir/clean

CMakeFiles/Agent.dir/depend:
	cd /home/maaz/CLionProjects/Agent/cmake-build-debug && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/maaz/CLionProjects/Agent /home/maaz/CLionProjects/Agent /home/maaz/CLionProjects/Agent/cmake-build-debug /home/maaz/CLionProjects/Agent/cmake-build-debug /home/maaz/CLionProjects/Agent/cmake-build-debug/CMakeFiles/Agent.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/Agent.dir/depend

