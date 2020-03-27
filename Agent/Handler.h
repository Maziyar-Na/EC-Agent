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
#include <mutex>
#include <chrono>
#include <ctime>

#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/io/zero_copy_stream_impl.h>

#define __BUFFSIZE__ 2048 //TODO: may need to change
#define _CPU_ 0
#define _MEM_ 1
#define _INIT_ 2
#define _SLICE_ 3
#define _CONNECT_ 4
#define _MEM_LIMIT_ 5
#define _SET_MAX_MEM_ 6
#define _READ_QUOTA_ 7

#define __RESIZE_MAX_MEM_SYSCALL__ 336
#define __INCR_MEMCG_MARGIN_SYSCALL__ 337
#define __RESIZE_QUOTA_SYSCALL_ 338
#define __READ_QUOTA_SYSYCALL__ 339

using namespace google::protobuf::io;

namespace ec {
    namespace agent {
        //TODO: these should not be typedefs
        using string = std::string;
        class Handler {
        public:
            void run(int64_t clifd);
            static void *run_handler(void *server_args);

        private:
            char *handle_request(char *buff, int &tx_size);
            static uint64_t handle_mem_req(uint64_t cgroup_id);
            static uint64_t handle_cpu_req(uint64_t cgroup_id, uint64_t quota, uint64_t &updated_quota);
            static uint64_t handle_resize_max_mem(uint16_t cgroup_id, uint64_t new_limit, int is_memsw);

            static int64_t handle_read_quota(uint16_t cgroup_id);

            static std::string connect_container(const std::string &server_ip, const std::string &container_name);
            static std::string exec(std::string &command);
            static google::protobuf::uint32 readHdr(char *buf);
            std::mutex protolock;
        };

        struct serv_thread_args {
            serv_thread_args()      = default;
            int64_t clifd           = 0;
            Handler* req_handler    = nullptr;
        };
    }
}
#endif //AGENT_HANDLER_H
