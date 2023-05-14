FROM golang:1.20.4

ARG PROJECT_NAME=
ARG COMPANY_NAME=

WORKDIR /mnt/kvex/v4/${COMPANY_NAME}/${PROJECT_NAME}

RUN apt install -y make

COPY . .
RUN make build-release-x64 WITH_STDERR=true

RUN rm -rf /mnt/kvex/v4/${COMPANY_NAME}/${PROJECT_NAME}/src/
RUN mv ./run.sh /run.sh
RUN chmod +x /run.sh

ENV PROJECT_NAME=${PROJECT_NAME}
ENV COMPANY_NAME=${COMPANY_NAME}

CMD ["/run.sh"]