FROM ubuntu:latest
ADD Agent /tmp/
RUN chmod +x /tmp/Agent
RUN apt-get update
# RUN apt-get -y install curl && curl -s https://packagecloud.io/install/repositories/akopytov/sysbench/script.deb.sh |  bash && apt -y install sysbench
ENTRYPOINT ["/tmp/Agent"]