cmake_minimum_required(VERSION 3.22.1)


project(CLIENT)

add_executable(${PROJECT_NAME} main.cpp net/requests/requests.cpp net/requests/requests.h)
add_executable(test_client tests/test-client.cpp)