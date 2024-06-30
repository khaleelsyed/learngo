FROM golang:1.22.4 as base

ARG UID=1000
ARG GID=1001
ENV CODE_DIR=/src
ENV SCRIPTS_DIR=/scripts

USER root

RUN groupadd -g $GID gophers
RUN useradd -rm -d /home/gopher -g $GID -u $UID -s /bin/bash gopher

COPY .$CODE_DIR $CODE_DIR
RUN chown -R gopher:gophers $CODE_DIR
RUN chmod -R 775 ${CODE_DIR}

FROM base as development

COPY .$SCRIPTS_DIR $SCRIPTS_DIR
RUN chown -R gopher:gophers $SCRIPTS_DIR
RUN chmod -R 775 ${SCRIPTS_DIR}

RUN ${SCRIPTS_DIR}/dev_setup.sh

USER gopher

WORKDIR $CODE_DIR
