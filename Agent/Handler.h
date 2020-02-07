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
#include "proto/msg.pb.h"

#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/io/zero_copy_stream_impl.h>

#define __BUFFSIZE__ 2048
#define _CPU_ 0
#define _MEM_ 1
#define _INIT_ 2
#define _SLICE_ 3
#define _CONNECT_ 4
#define __NR_SYSCALL__ 336

using namespace std;
using namespace google::protobuf::io;


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
                char* handle_request(char* buff);
                
                uint64_t handle_mem_req(uint64_t cgroup_id);

                uint64_t connect_container(string server_ip, string container_name);

                std::string exec(string command);

                google::protobuf::uint32 readHdr(char *buf);

        };

        typedef struct serv_thread_args
        {
            int64_t clifd;

            Handler* req_handler;

        } serv_thread_args_t;  
    }
}
#endif //AGENT_HANDLER_H
