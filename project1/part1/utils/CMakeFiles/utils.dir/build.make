# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.10

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
CMAKE_COMMAND = /usr/bin/cmake

# The command to remove a file.
RM = /usr/bin/cmake -E remove -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/lisiting/OSDS-2022Fall/project1

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/lisiting/OSDS-2022Fall/project1

# Include any dependencies generated for this target.
include part1/utils/CMakeFiles/utils.dir/depend.make

# Include the progress variables for this target.
include part1/utils/CMakeFiles/utils.dir/progress.make

# Include the compile flags for this target's objects.
include part1/utils/CMakeFiles/utils.dir/flags.make

part1/utils/CMakeFiles/utils.dir/utils.cc.o: part1/utils/CMakeFiles/utils.dir/flags.make
part1/utils/CMakeFiles/utils.dir/utils.cc.o: part1/utils/utils.cc
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/lisiting/OSDS-2022Fall/project1/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building CXX object part1/utils/CMakeFiles/utils.dir/utils.cc.o"
	cd /home/lisiting/OSDS-2022Fall/project1/part1/utils && /usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/utils.dir/utils.cc.o -c /home/lisiting/OSDS-2022Fall/project1/part1/utils/utils.cc

part1/utils/CMakeFiles/utils.dir/utils.cc.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/utils.dir/utils.cc.i"
	cd /home/lisiting/OSDS-2022Fall/project1/part1/utils && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/lisiting/OSDS-2022Fall/project1/part1/utils/utils.cc > CMakeFiles/utils.dir/utils.cc.i

part1/utils/CMakeFiles/utils.dir/utils.cc.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/utils.dir/utils.cc.s"
	cd /home/lisiting/OSDS-2022Fall/project1/part1/utils && /usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/lisiting/OSDS-2022Fall/project1/part1/utils/utils.cc -o CMakeFiles/utils.dir/utils.cc.s

part1/utils/CMakeFiles/utils.dir/utils.cc.o.requires:

.PHONY : part1/utils/CMakeFiles/utils.dir/utils.cc.o.requires

part1/utils/CMakeFiles/utils.dir/utils.cc.o.provides: part1/utils/CMakeFiles/utils.dir/utils.cc.o.requires
	$(MAKE) -f part1/utils/CMakeFiles/utils.dir/build.make part1/utils/CMakeFiles/utils.dir/utils.cc.o.provides.build
.PHONY : part1/utils/CMakeFiles/utils.dir/utils.cc.o.provides

part1/utils/CMakeFiles/utils.dir/utils.cc.o.provides.build: part1/utils/CMakeFiles/utils.dir/utils.cc.o


# Object files for target utils
utils_OBJECTS = \
"CMakeFiles/utils.dir/utils.cc.o"

# External object files for target utils
utils_EXTERNAL_OBJECTS =

part1/utils/libutils.a: part1/utils/CMakeFiles/utils.dir/utils.cc.o
part1/utils/libutils.a: part1/utils/CMakeFiles/utils.dir/build.make
part1/utils/libutils.a: part1/utils/CMakeFiles/utils.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/lisiting/OSDS-2022Fall/project1/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Linking CXX static library libutils.a"
	cd /home/lisiting/OSDS-2022Fall/project1/part1/utils && $(CMAKE_COMMAND) -P CMakeFiles/utils.dir/cmake_clean_target.cmake
	cd /home/lisiting/OSDS-2022Fall/project1/part1/utils && $(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/utils.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
part1/utils/CMakeFiles/utils.dir/build: part1/utils/libutils.a

.PHONY : part1/utils/CMakeFiles/utils.dir/build

part1/utils/CMakeFiles/utils.dir/requires: part1/utils/CMakeFiles/utils.dir/utils.cc.o.requires

.PHONY : part1/utils/CMakeFiles/utils.dir/requires

part1/utils/CMakeFiles/utils.dir/clean:
	cd /home/lisiting/OSDS-2022Fall/project1/part1/utils && $(CMAKE_COMMAND) -P CMakeFiles/utils.dir/cmake_clean.cmake
.PHONY : part1/utils/CMakeFiles/utils.dir/clean

part1/utils/CMakeFiles/utils.dir/depend:
	cd /home/lisiting/OSDS-2022Fall/project1 && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/lisiting/OSDS-2022Fall/project1 /home/lisiting/OSDS-2022Fall/project1/part1/utils /home/lisiting/OSDS-2022Fall/project1 /home/lisiting/OSDS-2022Fall/project1/part1/utils /home/lisiting/OSDS-2022Fall/project1/part1/utils/CMakeFiles/utils.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : part1/utils/CMakeFiles/utils.dir/depend
