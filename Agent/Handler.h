//
// Created by maaz on 9/25/19.
//

#ifndef AGENT_HANDLER_H
#define AGENT_HANDLER_H

#include <memory.h>
#include <stdint-gcc.h>
#include <iostream>
#include <sstream>
#include <cstdio>
#include <stdexcept>
#include <string>
#include <array>
#include "om.h"
#define __BUFFSIZE__ 128
#define _CPU_ 0
#define _MEM_ 1
#define _INIT_ 2
#define _SLICE_ 3
#define _CONNECT_ 4
#define _MEM_LIMIT_ 5
#define __NR_SYSCALL__ 336

using namespace std;

namespace ec {
    namespace agent {

        typedef struct ec_msg {
            om::net::ip4_addr client_ip;
            uint32_t cgroup_id;
            uint32_t req_type;        //0: cpu, 1: mem, 2: init, 3: slice, 4: create_cont
            uint64_t rsrc_amnt;      //amount of resources (cpu/mem)
            uint32_t request;        //1: request, 0: give back
            uint64_t runtime_remaining;
            uint64_t cont_name;

                void run(int64_t clifd);

        } ec_msg_t;

        class Handler {

            public:
                void run(int64_t clifd);
                
                static void* run_handler(void* server_args);


            private:
                ec_msg_t* handle_request(char* buff);
                
                uint64_t handle_mem_req(ec_msg_t* req);

                uint64_t connect_container(ec_msg_t* req);

                std::string exec(std::string cmd);

        };

        typedef struct serv_thread_args
        {
            int64_t clifd;

            Handler* req_handler;

        } serv_thread_args_t;  
    }
}
#endif //AGENT_HANDLER_H
