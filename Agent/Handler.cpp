//
// Created by maaz on 9/25/19.
//

#include <unistd.h>
#include "Handler.h"

uint64_t ec::agent::Handler::handle_mem_req(uint64_t cgroup_id) {
    uint64_t ret = 0, avail_mem = 0;

    std::cout << "cgroup_id: " << cgroup_id << std::endl;
    ret = syscall(__NR_SYSCALL__, cgroup_id, false);

    cout << "[INFO] EC Agent: Reclaimed memory is: " << ret << endl;
    avail_mem += ret;
    
    return avail_mem;
}

std::string ec::agent::Handler::connect_container(string server_ip, string container_name) {
    string cmd = "sudo docker ps -a | grep k8s_" + container_name + " | awk '{print $1, $3}'";

    // Todo: Change this so that we looop for maximum of 5 seconds or until a new container is created 
    
    sleep(5);

    std::string container_id = exec(cmd);
    // std::cout << "[dbg]: container_id:  " << container_id << std::endl;
    // This is the where we can confirm whether the container was successfully created and deployed
    if (container_id.size() == 0) {
        std::cout << "[dbg]: No container found with name: " << container_name << std::endl;
        return "error";
    }
    size_t pos = container_id.find(" ");    
    container_id = container_id.substr(0, pos);
    // std::cout << "[dbg] Docker Container ID: " << container_id << std::endl;

    cmd = "sudo docker inspect --format '{{ .State.Pid }}' " + container_id;
    string pid = exec(cmd);
    // // every container created should have a PID but in the case that it hasn't started, it won't have one. This is another error
    if (pid.size() == 0) {
        std::cout << "[dbg]: Error in getting PID for container with name:" << container_name << std::endl;
        return "error";
    }

    pid.erase(remove(pid.begin(), pid.end(), '\n'), pid.end());
    cmd = "./../../ec_syscalls/sys_connect " + server_ip + " " + pid + " 4444 " + "eno1";

    std::cout << "sysconnect command: " << cmd << std::endl;

    string output = exec(cmd);

    if (output.find("ERROR") != std::string::npos) {
        std::cout << "[dbg]: Error in calling sys_connect for container with name:" << container_name << std::endl;
        return "error";
    }
    return container_id;
}

//Helper function to handle request
char* ec::agent::Handler::handle_request(char* buff){

    google::protobuf::uint32 siz = readHdr(buff);
    msg_struct::ECMessage rx_msg;
    google::protobuf::io::ArrayInputStream arrayIn(buff, siz);
    google::protobuf::io::CodedInputStream codedIn(&arrayIn);
    codedIn.ReadVarint32(&siz);
    google::protobuf::io::CodedInputStream::Limit msgLimit = codedIn.PushLimit(siz);
    rx_msg.ParseFromCodedStream(&codedIn);
    codedIn.PopLimit(msgLimit);

    /* 
    cout<< "[RX MESSAGE DBG]: " << endl;
    cout<< "[request type]: " << rx_msg.req_type() << endl;
    cout<< "[cgroup id]: " << rx_msg.cgroup_id() << endl;
    */
    
    uint64_t ret = 0;
    std::string container_id;

    msg_struct::ECMessage tx_msg;
    tx_msg.set_req_type(rx_msg.req_type());
    switch (rx_msg.req_type()) {
        case _CPU_:
            cout << "[MAYBE TODO] Handling CPU request in the agent!" << endl;
            break;
        case _MEM_:
            ret = handle_mem_req(rx_msg.cgroup_id());
            ret = ret > 0 ? (uint64_t) ret: -1;
            break;
        case _INIT_:
            cout << "[DBG] Init message, not sure if we need this" << endl;
            break;
        case _SLICE_:
            cout << "[DBG] Slice message, not sure if we need this" << endl;
            break;
        case _CONNECT_:
            container_id = connect_container(rx_msg.client_ip(), rx_msg.payload_string());
            std::cout << "docker container docker id: " << container_id << std::endl;
            tx_msg.set_payload_string(container_id);
            break;
        default:
            cerr << "[ERROR] Not going in the right way! request type is invalid!" << endl;
    }
    tx_msg.set_rsrc_amnt(ret);
    tx_msg.set_cgroup_id(0);

    /*
    cout<< "[TX MESSAGE DBG LOG] " << endl;
    cout<< "[request type]: " << tx_msg.req_type() << endl;
    cout<< "[rsrc amt]: " << tx_msg.rsrc_amnt() << endl;
    cout<< "[payload string]: " << tx_msg.payload_string() << endl;
    */

    int tx_size = tx_msg.ByteSizeLong()+4;
    char* tx_buf = new char[tx_size];
    google::protobuf::io::ArrayOutputStream arrayOut(tx_buf, tx_size);
    google::protobuf::io::CodedOutputStream codedOut(&arrayOut);
    codedOut.WriteVarint32(tx_msg.ByteSizeLong());
    tx_msg.SerializeToCodedStream(&codedOut);
    
    //std::cout << "[EC Init] Sending Message to GCM with message of length: " << tx_size << std::endl; 

    return tx_buf;
}

void ec::agent::Handler::run(int64_t clifd) {
    char buff[__BUFFSIZE__];
    int64_t bytes_read;
    bzero(buff, __BUFFSIZE__);
    cout << "[RUN log] We are ready to accept request from GCM! fd is: " << clifd << endl;
    char* tx_buff;

    while( (bytes_read = read(clifd, buff, __BUFFSIZE__) ) > 0 ) {
        tx_buff = handle_request(buff);

        if (write(clifd, (void*) tx_buff, __BUFFSIZE__) < 0) {
            cout <<"[ERROR] writing to socket connection (Agent -> GCM) Failed! " << endl;
        }
    }

    pthread_exit(NULL);
}

void* ec::agent::Handler::run_handler(void* server_args)
{
    // cout << "[dbg] run_handler: thread executed!" << endl;
    auto *args = static_cast<serv_thread_args*>(server_args);
    args->req_handler->run(args->clifd);
    return NULL;
}

std::string ec::agent::Handler::exec(string command) {
    // char buffer[128];
    // string result = "";

    // // Open pipe to file
    // FILE* pipe = popen(command.c_str(), "r");
    // if (!pipe) {
    //     return "popen failed!";
    // }

    // // read till end of process:
    // while (!feof(pipe)) {

    //     // use buffer to read and add to result
    //     if (fgets(buffer, 128, pipe) != NULL)
    //         result += buffer;
    // }

    // pclose(pipe);
    // return result;
    // run a process and create a streambuf that reads its stdout and stderr
    string data;
    FILE * stream;
    const int max_buffer = 256;
    char buffer[max_buffer];
    command.append(" 2>&1");

    stream = popen(command.c_str(), "r");
    if (stream) {
    while (!feof(stream))
    if (fgets(buffer, max_buffer, stream) != NULL) data.append(buffer);
    pclose(stream);
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

