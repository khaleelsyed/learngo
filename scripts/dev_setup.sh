#!/bin/bash

mkdir /home/gopher/.ssh

touch /home/gopher/.ssh/known_hosts

chown -R gopher:gophers /home/gopher/.ssh/
chmod -R 775 /home/gopher/.ssh/
