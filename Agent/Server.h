//
// Created by maaz on 9/24/19.
//

#ifndef AGENT_SERVER_H
#define AGENT_SERVER_H


#include <stdint-gcc.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <iostream>
#include <netinet/in.h>
#include <pthread.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include "Handler.h"
#include <sys/select.h>


#define _MAXGCMNO_ 30

using namespace std;

namespace ec {
    namespace agent {


        class Server {

        public:
            explicit Server(uint16_t _port);

            struct ec_agent {
                ec_agent() = default;
                int64_t sockfd          = 0;
                struct sockaddr_in addr;
            };

            void init_agent_server();
            void run();

        private:
            uint16_t port;

            ec_agent *_ec_agent;

            bool is_avail;

            //Handler* req_handler;
        };

    }
}


#endif //AGENT_SERVER_H
