cmake_minimum_required(VERSION 3.16)

project(fsmd VERSION 1.0.0)

set(CMAKE_EXPORT_COMPILE_COMMANDS ON)
set (CMAKE_CXX_STANDARD 17)
set (CXX_STANDARD_REQUIRED ON)
set (CMAKE_CXX_EXTENSIONS ON)

set(CMAKE_CXX_FLAGS_DEBUG_INIT "-Wall")
set(CMAKE_CXX_FLAGS_RELEASE_INIT "-Wall -O3")

include_directories(fsmd include)

add_executable(finslice_demo
  src/udplisten.cpp
  src/finslice_demo.cpp
)
