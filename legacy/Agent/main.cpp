
#include "Server.h"
#define _PORT_ 4445

int main() {

    cout << "[INFO] EC Agent started!" << endl;

    //Initialize and start the elastic container agent
    auto *s = new ec::agent::Server(_PORT_);
    s->init_agent_server();
    s->run();

    return 0;
}