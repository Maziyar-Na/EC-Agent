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

uint64_t ec::agent::Handler::connect_container(string server_ip, string container_name) {
    std::cout << "[dbg] In connect_container: " << std::endl;
    std::cout << "Container name: " << container_name << std::endl;

    string cmd = "sudo docker ps -a | grep k8s_" + container_name + " | awk '{print $1, $3}'";
    std::cout << "docker cmd: " << cmd << std::endl;
    // Puzzle here: what is the best way to wait for the container to be created? I might be sending stuff
    // to agent too quickly here..
    sleep(5);
    std::string container_id = exec(cmd);
    // This is the where we can confirm whether the container was successfully created and deployed
    std::cout << "[dbg]: " << container_id << std::endl;
    if (container_id.size() == 0) {
        std::cout << "[dbg]: No container found with name:" << container_name << std::endl;
        return (uint64_t) -1;
    }
    size_t pos = container_id.find(" ");    
    container_id = container_id.substr(0, pos);
    // std::cout << "[dbg] Docker Container ID: " << container_id << std::endl;

    cmd = "sudo docker inspect --format '{{ .State.Pid }}' " + container_id;
    string pid = exec(cmd);
    // // every container created should have a PID but in the case that it hasn't started, it won't have one. This is another error
    if (pid.size() == 0) {
        std::cout << "[dbg]: Error in getting PID for container with name:" << container_name << std::endl;
        return (uint64_t) -1;
    }
    cout << "[dbg] Docker Container PID: " << pid << endl;
    pid.erase(remove(pid.begin(), pid.end(), '\n'), pid.end());
    cmd = "../../ec_syscalls/sys_connect " + server_ip + " " + pid + " 4444";

    std::cout << "sysconnect command: " << cmd << std::endl;

    string output = exec(cmd);

    if (output.find("ERROR") != std::string::npos) {
        std::cout << "[dbg]: Error in calling sys_connect for container with name:" << cont_name_string << std::endl;
        return (uint64_t) -1;
    }
    
    uint64_t pid_return_value;
    std::istringstream iss(pid);
    iss >> pid_return_value;
    return pid_return_value;
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
    
    // cout<< "[RX MESSAGE DBG]: " << endl;
    // cout<< "[ip]: " << rx_msg.client_ip() << endl;
    // cout<< "[request type]: " << rx_msg.req_type() << endl;
    // cout<< "[payload string]: " << rx_msg.payload_string() << endl;
    // cout<< "[cgroup id]: " << rx_msg.cgroup_id() << endl;

    uint64_t ret = 0;
    switch (rx_msg.req_type() ) {
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
            ret = connect_container(rx_msg.client_ip(), rx_msg.payload_string());
            break;
        default:
            cerr << "[ERROR] Not going in the right way! request type is invalid!" << endl;
    }

    msg_struct::ECMessage tx_msg;
    tx_msg.set_req_type(rx_msg.req_type());
    tx_msg.set_rsrc_amnt(ret);
    tx_msg.set_payload_string(rx_msg.payload_string());

    // cout<< "[TX MESSAGE DBG]: " << endl;
    // cout<< "[ip]: " << tx_msg.client_ip() << endl;
    // cout<< "[rsrc amt]: " << tx_msg.rsrc_amnt() << endl;
    // cout<< "[payload string]: " << tx_msg.payload_string() << endl;
    
    int tx_size = tx_msg.ByteSizeLong()+4;
    char* tx_buf = new char[tx_size];
    google::protobuf::io::ArrayOutputStream arrayOut(tx_buf, tx_size);
    google::protobuf::io::CodedOutputStream codedOut(&arrayOut);
    codedOut.WriteVarint32(tx_msg.ByteSizeLong());
    tx_msg.SerializeToCodedStream(&codedOut);
    return tx_buf;
}

void ec::agent::Handler::run(int64_t clifd) {
    char buff[__BUFFSIZE__];
    int64_t bytes_read;
    bzero(buff, __BUFFSIZE__);
    cout << "[dbg] run: We are ready to accept request from GCM! fd is: " << clifd << endl;
    char* tx_buff;

    while( (bytes_read = read(clifd, buff, __BUFFSIZE__) ) > 0 ) {
        tx_buff = handle_request(buff);

        if (write(clifd, (char*) tx_buff, sizeof(*tx_buff)) < 0)
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

std::string ec::agent::Handler::exec(string command) {
    char buffer[128];
    string result = "";

    // Open pipe to file
    FILE* pipe = popen(command.c_str(), "r");
    if (!pipe) {
        return "popen failed!";
    }

    // read till end of process:
    while (!feof(pipe)) {

        // use buffer to read and add to result
        if (fgets(buffer, 128, pipe) != NULL)
            result += buffer;
    }

    pclose(pipe);
    return result;
}



google::protobuf::uint32 ec::agent::Handler::readHdr(char *buf)
{
  google::protobuf::uint32 size;
  google::protobuf::io::ArrayInputStream ais(buf,4);
  CodedInputStream coded_input(&ais);
  coded_input.ReadVarint32(&size);//Decode the HDR and get the size
  cout<<" size of RX payload is "<<size<<endl;
  return size;
}

