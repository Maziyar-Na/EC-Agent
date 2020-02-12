//
// Created by maaz on 9/25/19.
//

#include <unistd.h>
#include "Handler.h"

uint64_t ec::agent::Handler::handle_mem_req(ec_msg_t* req) {
    uint64_t ret = 0, avail_mem = 0;
    //for ( count = 0; count < req -> num_of_cgroups; ++count)
    //{

    std::cout << "cgroup_id: " << req->cgroup_id << std::endl;
    ret = syscall(__NR_SYSCALL__, req->cgroup_id, false);

    cout << "[INFO] EC Agent: Reclaimed memory is: " << ret << endl;
    avail_mem += ret;
    //}
    return avail_mem;
}

uint64_t ec::agent::Handler::connect_container(ec_msg_t* req) {
    std::cout << "[dbg] In connect_container: " << std::endl;
    std::cout << "Container name: " << std::to_string(req->cont_name) << std::endl;
    
    // std::cout << "request type: " << std::to_string(req->req_type) << std::endl;
    // std::cout << "request ip: " << std::to_string(req->client_ip) << std::endl;
    // std::cout << "cgroup id: " << std::to_string(req->cgroup_id) << std::endl;
    // std::cout << "req_type: " << std::to_string(req->req_type) << std::endl;
    std::cout << "Image Type: " << std::to_string(req->runtime_remaining) << std::endl;
    // std::cout << "request request/response: " << std::to_string(req->request) << std::endl;
    
    // Again, This needs to be changed when we implement sending strings across in the correct way, this is an awful way
    // to pass the image and container name
    std::string image_name_type = std::to_string(req->runtime_remaining);
    std::string image_name;
    if (image_name_type == "1") {
        image_name = "nginx";
    } else if (image_name_type == "2"){
        image_name = "redis";
    } else {
        image_name = "nginx";
    }

    std::string cont_name_string = std::to_string(req->cont_name);
    // This is the format of the name that the pod created the container on the GCM master node..
    cont_name_string = cont_name_string + "-" + image_name;
    char cont_name_cstring[cont_name_string.size() + 1];
    strcpy(cont_name_cstring, cont_name_string.c_str());

    std::string cmd = "sudo docker ps -a | grep ";
    cmd = cmd + "k8s_" + cont_name_cstring + " | awk '{print $1,$3}'";
    std::cout << "docker cmd: " << cmd << std::endl;
    // Puzzle here: what is the best way to wait for the container to be created? I might be sending stuff
    // to agent too quickly here..
    sleep(5);
    std::string container_id = exec(cmd);
    // This is the where we can confirm whether the container was successfully created and deployed
    std::cout << "[should not hang]: " << std::endl;
    std::cout << "[dbg]: " << container_id << std::endl;
    if (container_id.size() == 0) {
        std::cout << "[dbg]: No container found with name:" << cont_name_string << std::endl;
        return 1;
    }
    size_t pos = container_id.find(" ");    
    container_id = container_id.substr(0, pos);
    std::cout << "[dbg] Docker Container ID: " << container_id << std::endl;
    
    // Get the PID of the container to call sys_connect with..
    std::string cmd_pid = "sudo docker inspect --format '{{ .State.Pid }}' " + container_id;
    std::string pid_string = exec(cmd_pid);
    // every container created should have a PID but in the case that it hasn't started, it won't have one. This is another error
    if (pid_string.size() == 0) {
        std::cout << "[dbg]: Error in getting PID for container with name:" << cont_name_string << std::endl;
        return 1;
    }
    pos = pid_string.find(":");    
    pid_string = pid_string.substr(pos + 1, pid_string.size()-1);
    //std::cout << pid_string << std::endl;

    std::string cmd_sysconnect = "./../../../ec_syscalls/sys_connect ";

    std::string ip_address_string = req->client_ip.to_string();
    cmd_sysconnect += ip_address_string + " " + pid_string + " 4444";

    // Todo: This needs to change and needs to be passed in from the GCM. 
    // i.e a different port number corresponds to a different distributed container
    // Debugging purposes

    std::cout << "sysconnect command: " << cmd_sysconnect << "" << std::endl;

    std::string sys_connect_output = exec(cmd_sysconnect);
    // Finally, this is the case for when sys_connect fails.. another source for error
    // if (sys_connect_output == "-1") {
    //     std::cout << "[dbg]: Error in calling sys_connect for container with name:" << cont_name_string << std::endl;
    //     return 1;
    // }
    std::cout << "[DEBUG]: SYS_CONNECT output: " << sys_connect_output << "" << std::endl;
    return 0;
}

//Helper function to handle request
ec::agent::ec_msg_t* ec::agent::Handler::handle_request(char* buff){

    uint64_t ret = 0;
    ec_msg_t* req;
    req = (ec_msg_t*)buff;
    
    ec_msg_t* res;
    res = req;

    switch (req->req_type) {
        case _CPU_:
            cout << "[MAYBE TODO] Handling CPU request in the agent!" << endl;
            break;
        case _MEM_:
            ret = handle_mem_req(req);
            ret = ret > 0 ? (uint64_t) ret: -1;
            res->request = 0;
            res->rsrc_amnt = ret;
            break;
        case _INIT_:
            cout << "[DBG] Init message, not sure if we need this" << endl;
            break;
        case _SLICE_:
            cout << "[DBG] Slice message, not sure if we need this" << endl;
            break;
        case _CONNECT_:
            ret = connect_container(req);
            res->request = 0;
            res->rsrc_amnt = ret;
            break;
        case _MEM_LIMIT_:
            ret = 2061374;//TODO: temporary. for testing purpose. we need a syscall to extract mem limit based on cgroup id
            res->rsrc_amnt = ret;
            break;
        default:
            cerr << "[ERROR] Not going in the right way! request type is invalid!" << endl;
    }
    return res;
}

void ec::agent::Handler::run(int64_t clifd) {
    char buff[__BUFFSIZE__];
    int64_t bytes_read;
    bzero(buff, __BUFFSIZE__);
    cout << "[dbg] run: We are ready to accept request from GCM! fd is: " << clifd << endl;
    while( (bytes_read = read(clifd, buff, __BUFFSIZE__) ) > 0 ) {
        ec_msg_t* res;

        res = handle_request(buff);

        std::cout << "Container Status: (0=success, 1=fail)... " << res->rsrc_amnt << std::endl;
        if (write(clifd, (char*) res, sizeof(*res)) < 0)
            cout <<"[ERROR] writing to socket connection (Agent -> GCM) Failed! " << endl;
    }

    pthread_exit(NULL);
}

void* ec::agent::Handler::run_handler(void* server_args)
{
    cout << "[dbg] run_handler: thread executed!" << endl;
    auto *args = static_cast<serv_thread_args*>(server_args);
    args->req_handler->run(args->clifd);
    return NULL;
}

std::string ec::agent::Handler::exec( std::string cmd) {
    /*char buffer[128];
    std::string result = "";
    FILE* pipe = popen(cmd, "r");
    if (!pipe) throw std::runtime_error("popen() failed!");
    try {
        while (fgets(buffer, sizeof buffer, pipe) != NULL) {
            std::cout << "[dbg] hangs in here! " << std:: endl;
            result += buffer;
        }
    } catch (...) {
        pclose(pipe);
        throw;
    }
    pclose(pipe);
    return result;*/
    std::string file_name = "result.txt" ;
    std::system( ( cmd + " > " + file_name ).c_str() ) ; // redirect output to file

    // open file for input, return string containing characters in the file
    std::ifstream file(file_name) ;
    return { std::istreambuf_iterator<char>(file), std::istreambuf_iterator<char>() } ;
}