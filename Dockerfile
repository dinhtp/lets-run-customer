FROM centos:7

EXPOSE 9090

ADD ./bin/run-customer-service /usr/bin/run-customer-service

CMD ["run-customer-service", "serve"]