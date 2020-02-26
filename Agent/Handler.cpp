//
// Created by maaz on 9/25/19.
//



#include <unistd.h>
#include "Handler.h"

uint64_t ec_agent::Handler::handle_mem_req(ec_reclaim_msg_t* req) {
    uint64_t ret = 0, avail_mem = 0;
    //for ( count = 0; count < req -> num_of_cgroups; ++count)
    //{
    ret = syscall(__NR_SYSCALL__, req->cgroup_id, false);

    cout << "[INFO] EC Agent: Reclaimed memory is: " << ret << endl;
    avail_mem += ret;
    //}
    return avail_mem;
}

uint64_t ec_agent::Handler::handle_cpu_req(ec_agent::ec_reclaim_msg_t *req) {
    uint64_t ret = 0;
    std::cout << "setting quota to: " << req->_quota << std::endl;
    ret = syscall(__RESIZE_QUOTA_SYSCALL_, req->cgroup_id, req->_quota / 1000); //must divide since kernel multiplies
    if(ret) {
        std::cout << "quota set failed" << std::endl;
    }

    return 0;//req->_quota;
}


//Helper function to handle request
uint64_t  ec_agent::Handler::handle_request(char* buff){

    uint64_t ret = 0;
    auto *req = new ec_reclaim_msg_t;
    req = (ec_reclaim_msg_t*)buff;
    std::cout << "req: id, is_mem, quota: " << req->cgroup_id << ", " << req->is_mem << ", " << req->_quota << std::endl;
    std::cout << "reclaim msg size: " << sizeof(*req) << std::endl;

    switch (req -> is_mem) {
        case TRUE:
            ret = handle_mem_req(req);
            break;

        case FALSE:
            std::cout << "[Agent DBG]: handle cpu_req" << std::endl;
            ret = handle_cpu_req(req);
            break;

        default:
            cerr << "[ERROR] Not going in the right way! request type is neither true nor false!" << endl;
    }
//    delete req;
    return ret;
}

void ec_agent::Handler::run(int64_t clifd) {
    char buff[__BUFFSIZE__];
    int64_t bytes_read;
    bzero(buff, __BUFFSIZE__);
    cout << "[dbg] run: We are ready to accept request from GCM! fd is: " << clifd << endl;
    while( (bytes_read = read(clifd, buff, __BUFFSIZE__) ) > 0 ) {
        uint64_t ret = 0;

        ret = handle_request(buff);

        //const char* res = ret > 0 ? (const char*)ret : NOMEM;
        if(ret) {
            if (write(clifd, (const char *) &ret, sizeof(ret)) < 0) {
                cout << "[ERROR] writing to socket connection (Agent -> GCM) Failed! " << endl;
            }
        }

    }

    pthread_exit(nullptr);
}

void* ec_agent::Handler::run_handler(void* server_args)
{
    cout << "[dbg] run_handler: thread executed!" << endl;
    auto args = static_cast<serv_thread_args_t*>(server_args);
    args->req_handler->run(args->clifd);
    return nullptr;
}

