# remote executor

## ssh
maybe better than libssh is libssh2 since seems better supported.
[good example here](https://www.libssh2.org/examples/ssh2.html)

## Bridge to serial devices and ssh hosts
The idea is a microservice that 
* opens serial/ssh connection 
* open listening socket for unix domain socket

Why:
* clients connect to UDS and get serial/ssh communication
* One UDS socket for multiple ssh or serial connections
* UDS channel is multiplexer
    * communication via JSON
    * JSON packages: 
        * "type: data | control | upload | download"
        * "channel: serial | ssh"
        * "remote: /dev/ttyS0 | 192.168.0.1:22"
        * "payload-type: binary | text"
        * "payload: Last login: Tue..."




