# 

import socket
import select
import time

EPOLLHUP = select.EPOLLHUP
EPOLLERR = select.EPOLLERR
EPOLLIN = select.EPOLLIN
EPOLLOUT = select.EPOLLOUT

EPOLLDFT = EPOLLHUP|EPOLLERR|EPOLLIN
EPOLLALL = EPOLLDFT|EPOLLOUT

HOST = '127.0.0.1'
PORT = 8888
serv_fd = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
serv_fd.setblocking(False)



serv_fd.connect_ex((HOST, PORT))

eph = select.epoll()
eph.register(serv_fd.fileno(), EPOLLDFT|EPOLLOUT)

def handle_write(fd):
    serv_fd.send('hi %.3f' % time.time())
    eph.modify(serv_fd, EPOLLDFT)

def handle_read(fd):
    buf = serv_fd.recv(100)
    if not buf:
        serv_fd.close()
    print 'from sever: %s' % buf
    

while True:
    ret = eph.poll(2)
    if not ret:
        eph.modify(serv_fd, EPOLLALL)
        continue
    for fd, evt in ret:
        if evt & EPOLLOUT:
            handle_write(fd)
        elif evt & EPOLLIN:
            handle_read(fd)
        else:
            break
