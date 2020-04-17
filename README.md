# EC-Agent
EC Agent implemented in Go before merge into Kubelet

## Steps to run EC Agent
These instructions assume that you wish to run the Agent as a standalone process and not as part of the kubelet. Additionally, it is assumed that the machine you're running the Agent on is running the modified version of the EC-4.20.16 Kernel and you've inserted the appropriate kernel modules you wish to test (i.e. cgroup_connection is required to make a conneciton to GCM, resize_quota for CPU requests, etc.)

### Step 1: Setting up your Environment
1. Install [Docker](https://docs.docker.com/engine/install/ubuntu/)
2. Install [Kubernetes](https://blog.sourcerer.io/a-kubernetes-quick-start-for-people-who-know-just-enough-about-docker-to-get-by-71c5933b4633)
    ```
    sudo apt-get update && sudo apt-get install -y apt-transport-https
    curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
    ```
    Add the following line to `/etc/apt/sources.list.d/kubernetes.list`
    - `deb http://apt.kubernetes.io/ kubernetes-xenial main`
    ```
    sudo apt-get update
    sudo apt-get install -y kubelet kubeadm kubectl
    ```
3. Install latest version of Go (note: testing was done on v1.14.1)
    ```
        sudo apt-get update
        cd /tmp
        wget https://dl.google.com/go/go1.14.1.linux-amd64.tar.gz
        sudo tar -xvf go1.14.1.linux-amd64.tar.gz
        sudo mv go /usr/local
    ```
    Add the following lines in `~/.profile`:
    ```
        export GOROOT=/usr/local/go
        export GOPATH=$HOME/go
        export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
    ```
    Finally, `source ~/.profile` to update $GOPATH

<!-- 4. Install Protoc-gen-go
    - To install protoc-gen-go:
        ```
        go get -u github.com/golang/protobuf/protoc-gen-go
        ```
        The compiler plugin, protoc-gen-go, will be installed in `$GOPATH/bin` unless `$GOBIN` is set. It must be in your `$PATH` for the protocol compiler, `protoc`, to find it. -->

### Step 2: Running the Agent
1.  Make sure the node on which the agent is running is included in your Kubernetes cluster (i.e. can do `kubctl get nodes` on the master node to confirm)
2. Make sure the agent_ip has been updated correctly in ec_gcm/tests/app_def.json.
3. Also, make sure ec_gcm is on branch `ref-agent-go`. This is because I beleive there WAS a protobuf error in sending header information and so I've removed that bit from the gcm in that specific branch.
4. Finally, run the Agent on the node via: 
    - `go run main.go` 
    -  And build/run the ec_gcm code as you normally would