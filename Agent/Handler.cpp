//
// Created by maaz on 9/25/19.
//

#include <unistd.h>
#include "Handler.h"

uint64_t ec::agent::Handler::handle_mem_req(uint64_t cgroup_id) {
    uint64_t ret = 0, avail_mem = 0;

    std::cout << "cgroup_id: " << cgroup_id << std::endl;
    ret = syscall(__RESIZE_MAX_MEM_SYSCALL__, cgroup_id, false);

    std::cout << "[INFO] EC Agent: Reclaimed memory is: " << ret << std::endl;
    avail_mem += ret;
    
    return avail_mem;
}

uint64_t ec::agent::Handler::handle_cpu_req(uint64_t cgroup_id, uint64_t quota, uint64_t &update_quota) {
    uint64_t ret = 0;
    std::cout << "setting quota to: " << quota << std::endl;
    ret = syscall(__RESIZE_QUOTA_SYSCALL_, cgroup_id, quota / 1000); //must divide since kernel multiplies
    if(ret == 1) {
        std::cout << "quota set failed, ret: " << ret << std::endl;
        ret = 1;
    }
    else {
        std::cout << "quota set success! quota set to: " << ret << std::endl;
        update_quota = ret;
        ret = 0;
    }
    return ret;
}

//int64_t ec::agent::Handler::handle_read_quota(uint16_t cgroup_id) {
//    long ret = 0;
//    std::cout << "reading quota" << std::endl;
//    ret = syscall(__READ_QUOTA_SYSYCALL__, cgroup_id);
//    if(ret == -1) {
//        std::cout << "quota is RUNTIME_INF" << std::endl;
//        return -1;
//    }
//    else if(ret == 1) {
//        std::cout << "error reading quota!" << std::endl;
//        return 1;
//    }
//    return ret;
//}

uint64_t ec::agent::Handler::handle_resize_max_mem(uint16_t cgroup_id, uint64_t new_limit, int is_memsw) {
    uint64_t ret = 0;
    std::cout << "setting new mem limit to: " << new_limit << std::endl;
    ret = syscall(__RESIZE_MAX_MEM_SYSCALL__, cgroup_id, new_limit, false);

    if(!ret) {
        std::cout << "resize_max_mem failed!" << std::endl;
    }
    return ret;
}


std::string ec::agent::Handler::connect_container(const string &server_ip, const string &container_name) {
    std::cout << "server_ip: " << server_ip << ". " << "container_name: " << container_name << std::endl;
    string cmd = "sudo docker ps -a | grep k8s_" + container_name + " | awk '{print $1, $3}'";

    std::cout << "cmd: " << cmd << std::endl;
    // Todo: Change this so that we looop for maximum of 5 seconds or until a new container is created 
    
    sleep(5);

    std::string container_id = exec(cmd);
     std::cout << "[dbg]: container_id:  " << container_id << std::endl;
    // This is the where we can confirm whether the container was successfully created and deployed
    if (container_id.empty()) {
        std::cout << "[dbg]: No container found with name: " << container_name << std::endl;
        return "error";
    }
    size_t pos = container_id.find(" ");    
    container_id = container_id.substr(0, pos);
     std::cout << "[dbg] Docker Container ID: " << container_id << std::endl;

    cmd = "sudo docker inspect --format '{{ .State.Pid }}' " + container_id;
    string pid = exec(cmd);
    // // every container created should have a PID but in the case that it hasn't started, it won't have one. This is another error
    if (pid.empty()) {
        std::cout << "[dbg]: Error in getting PID for container with name:" << container_name << std::endl;
        return "error";
    }

    std::cout << "calling sysconnect" << std::endl;

    pid.erase(remove(pid.begin(), pid.end(), '\n'), pid.end());
    cmd = "../../../ec_syscalls/sys_connect " + server_ip + " " + pid + " 4444 " + "eno1";//"eno1";

    std::cout << "sysconnect command: " << cmd << std::endl;

    string output = exec(cmd);

    if (output.find("ERROR") != std::string::npos) {
        std::cout << "[dbg]: Error in calling sys_connect for container with name:" << container_name << std::endl;
        return "error";
    }
    return container_id;
}

//Helper function to handle request
char* ec::agent::Handler::handle_request(char* buff, unsigned long &tx_size){


    google::protobuf::uint32 siz = readHdr(buff);
    msg_struct::ECMessage rx_msg;
    google::protobuf::io::ArrayInputStream arrayIn(buff, siz);
    google::protobuf::io::CodedInputStream codedIn(&arrayIn);
    codedIn.ReadVarint32(&siz);
    google::protobuf::io::CodedInputStream::Limit msgLimit = codedIn.PushLimit(siz);
    rx_msg.ParseFromCodedStream(&codedIn);
    codedIn.PopLimit(msgLimit);

//    std::cout<< "[RX MESSAGE DBG]: " << std::endl;
//    std::cout<< "[request type]: " << rx_msg.req_type() <<  std::endl;
//    std::cout<< "[cgroup id]: " << rx_msg.cgroup_id() << std::endl;
//    std::cout << "rx_msg.quota: " << rx_msg.quota() << std::endl;
//    std::cout << "---------------------" << std::endl;
//    std::cout << "in quota: " << rx_msg.quota() << std::endl;
//    std::cout << "in seq num: " << rx_msg.request() << std::endl;
//    std::cout << "in time: " << std::chrono::system_clock::to_time_t(t) << std::endl;
    
    uint64_t ret = 0;
    uint64_t updated_quota = 0;
    std::string container_id;
    std::cout << "rx_msg.req_type(): " << rx_msg.req_type() << std::endl;

    msg_struct::ECMessage tx_msg;
    switch (rx_msg.req_type() ) {
        case _CPU_:
            std::cout << "[Agent DBG]: handle cpu_req" << std::endl;
//            protolock.lock();
            ret = handle_cpu_req(rx_msg.cgroup_id(), rx_msg.quota(), updated_quota);
//            protolock.unlock();
            break;
        case _MEM_:
            ret = handle_mem_req(rx_msg.cgroup_id());
            ret = ret > 0 ? (uint64_t) ret: -1;
            break;
        case _INIT_:
            std::cout << "[DBG] Init message, not sure if we need this" << std::endl;
            break;
        case _SLICE_:
            std::cout << "[DBG] Slice message, not sure if we need this" << std::endl;
            break;
        case _CONNECT_:
            container_id = connect_container(rx_msg.client_ip(), rx_msg.payload_string());
            std::cout << "docker container docker id: " << container_id << std::endl;
            tx_msg.set_payload_string(container_id);
            break;
        case _SET_MAX_MEM_:
            ret = handle_resize_max_mem(rx_msg.cgroup_id(), rx_msg.rsrc_amnt(), false);
            break;
//        case _READ_QUOTA_:
//            ret = handle_read_quota(rx_msg.cgroup_id());
//            break;
        default:
            std::cerr << "[ERROR] Not going in the right way! request type is invalid!" << std::endl;
    }

    tx_msg.set_req_type(rx_msg.req_type());
    tx_msg.set_rsrc_amnt(ret);
    tx_msg.set_quota(updated_quota);
    tx_msg.set_request(rx_msg.request());

//    std::cout << "out rsrc_amnt: " << tx_msg.rsrc_amnt() << std::endl;
//    std::cout << "out seq num: " << tx_msg.request() << std::endl;
//    std::cout << "out quota: " << tx_msg.quota() << std::endl;


    /*
    cout<< "[TX MESSAGE DBG LOG] " << endl;
    cout<< "[request type]: " << tx_msg.req_type() << endl;
    cout<< "[rsrc amt]: " << tx_msg.rsrc_amnt() << endl;
    cout<< "[payload string]: " << tx_msg.payload_string() << endl;
    */

    tx_size = tx_msg.ByteSizeLong()+4;
//    char* tx_buf = new char[tx_size];
    char *tx_buf{ new char[tx_size]{} };
    std::cout << *tx_buf << std::endl;
    google::protobuf::io::ArrayOutputStream arrayOut(tx_buf, (int)tx_size);
    google::protobuf::io::CodedOutputStream codedOut(&arrayOut);
    codedOut.WriteVarint32(tx_msg.ByteSizeLong());
    tx_msg.SerializeToCodedStream(&codedOut);
    
    //std::cout << "[EC Init] Sending Message to GCM with message of length: " << tx_size << std::endl; 
    return tx_buf;
}

void ec::agent::Handler::run(int64_t clifd) {
    char buff[__BUFFSIZE__] = {0};
    int64_t bytes_read;
//    bzero(buff, __BUFFSIZE__);
    std::cout << "[RUN log] We are ready to accept request from GCM! fd is: " << clifd << std::endl;
    char* tx_buff;
    unsigned long tx_size = 0;

    while( (bytes_read = read(clifd, buff, __BUFFSIZE__) ) > 0 ) {
        std::cout << "rx req!" << std::endl;
        tx_buff = handle_request(buff, tx_size);
//        std::cout << "handled req. tx_buff: " << tx_buff << std::endl;

        //TODO: temp fix. handle_request returns nullptr if it's a CPU req (on purpose)
        if(tx_buff) {
            if (write(clifd, (void*) tx_buff, tx_size) < 0) {
                std::cout <<"[ERROR] writing to socket connection (Agent -> GCM) Failed! " << std::endl;
            }
            delete[] tx_buff;
        }
        else {
            std::cout << "tx_buff == NULL" << std::endl;
        }
    }

    pthread_exit(nullptr);
}

void* ec::agent::Handler::run_handler(void* server_args)
{
    // cout << "[dbg] run_handler: thread executed!" << endl;
    auto *args = static_cast<serv_thread_args*>(server_args);
    args->req_handler->run(args->clifd);
    return nullptr;
}

std::string ec::agent::Handler::exec(string &command) {
    // run a process and create a streambuf that reads its stdout and stderr
    std::cout << "in exec" << std::endl;
    string data;
    FILE * stream;
    const int max_buffer = 256;
    char buffer[max_buffer];
//    command.append(" 2>&1");
    std::cout << "command: " << command << std::endl;

    stream = popen(command.c_str(), "r");
    if (stream) {
        while (!feof(stream)) {
            if (fgets(buffer, max_buffer, stream) != nullptr) {
                data.append(buffer);
            }
        }
        pclose(stream);
    }
    else {
        std::cout << "stream == NULL" << std::endl;
    }
    return data;
}

google::protobuf::uint32 ec::agent::Handler::readHdr(char *buf)
{
  google::protobuf::uint32 size;
  google::protobuf::io::ArrayInputStream ais(buf,4);
  CodedInputStream coded_input(&ais);
  coded_input.ReadVarint32(&size);//Decode the HDR and get the size
  return size;
}



