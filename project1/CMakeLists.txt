# cmake needs this line
cmake_minimum_required(VERSION 3.1)

# Define project name
project(proj1part2)

# Find OpenCV, you may need to set OpenCV_DIR variable
# to the absolute path to the directory containing OpenCVConfig.cmake file
# via the command line or GUI
find_package(OpenCV REQUIRED)

# If the package has been found, several variables will
# be set, you can find the full list with descriptions
# in the OpenCVConfig.cmake file.
# Print some message showing some of them
message(STATUS "OpenCV library status:")
message(STATUS "    config: ${OpenCV_DIR}")
message(STATUS "    version: ${OpenCV_VERSION}")
message(STATUS "    libraries: ${OpenCV_LIBS}")
message(STATUS "    include path: ${OpenCV_INCLUDE_DIRS}")

set(THREADS_PREFER_PTHREAD_FLAG ON)
find_package(Threads REQUIRED)

add_subdirectory(part1/concurrent_queue)
add_subdirectory(part1/utils)

# Declare the executable target built from your sources
add_executable(proj1part2 part2/threads.cc part1/concurrent_queue/concurrent_queue.cc part1/utils/utils.cc)

# Link your application with OpenCV libraries
target_link_libraries(proj1part2 PRIVATE ${OpenCV_LIBS} pthread)

target_include_directories(${PROJECT_NAME} PRIVATE
                           ${PROJECT_BINARY_DIR}
                           ${PROJECT_SOURCE_DIR}/part1/concurrent_queue
                           ${PROJECT_SOURCE_DIR}/part1/utils
                           )