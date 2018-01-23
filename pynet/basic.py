# 

import socket
import select

EPOLLHUP = select.EPOLLHUP
EPOLLERR = select.EPOLLERR
EPOLLIN = select.EPOLLIN
EPOLLOUT = select.EPOLLOUT

EPOLLDFT = EPOLLHUP|EPOLLERR|EPOLLIN

HOST = '127.0.0.1'
PORT = 8888
serv_fd = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
serv_fd.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
serv_fd.setblocking(False)
serv_fd.bind((HOST, PORT))

serv_fd.listen(5)
eph = select.epoll()


eph.register(serv_fd.fileno(), select.EPOLLIN|select.EPOLLERR|select.EPOLLHUP)

cli_map = {}
def handle_new(fd, evt):
    if evt & select.EPOLLERR:
        fd.close()
    elif evt & select.EPOLLHUP:
        fd.close()
    elif evt & select.EPOLLIN:
        cli, addr = fd.accept()
        #print 'addr %s, cli %s' % addr, cli
        cli.setblocking(False)
        eph.register(cli.fileno(), EPOLLDFT)
        cli_map[cli.fileno()] = cli
    else:
        fd.close()

def handle_cli(fd, evt):
    fd = cli_map[fd]
    if evt & select.EPOLLERR:
        fd.close()
    elif evt & select.EPOLLHUP:
        print 'cli disconnect'
        fd.close()
    elif evt & select.EPOLLIN:
        buf = fd.recv(100)
        if not buf:
            print 'cli disconnect'
            fd.close()
            return
        print 'cli said: %s' % buf
        fd.send('recv %s' % buf)
    else:
        fd.close()

while True:
    ret = eph.poll()
    if not ret:
        continue
    for fd, evt in ret:
        if fd == serv_fd.fileno():
            handle_new(serv_fd, evt)
        else:
            handle_cli(fd, evt)
    



serv_fd.close()