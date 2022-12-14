//
// Created by loona on 31/12/22.
//

#pragma once

#include <sys/socket.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <string>
#include <cstring>
#include <iostream>

namespace net {

    class UDPConn {
    private:
        int sock;
        sockaddr_in addr;
        socklen_t addr_len;
    public:
        UDPConn(char *host, uint port);

        ~UDPConn();

        // sends data
        ssize_t send(std::string req);

        // recieves data
        ssize_t recieve(char buffer[], int buffer_len) const;

        // allows a user to send binary data
        template <typename T>
        ssize_t send_bin(T data);
    };
}
