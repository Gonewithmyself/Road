# _*_ coding: utf-8 _*_


from loop import EPOLLALL, EPOLLIN
from channel import Socket
import socket

class Acceptor(object):
    def __init__(self, addr, loop):
        self.listen_sock = Socket()
        self.listen_fd = self.listen_sock.fd
        self.addr = addr
        self.loop = loop
    
    def listen(self):
        self.listen_sock.listen(self.addr)
        self.loop.register(self.listen_fd, EPOLLALL, self.process_new)
    
    def set_error_callback(self, func):
        self.error_callback = func
    
    def set_connect_callback(self, func):
        self.connect_callback = func

    def process_new(self, fd, evt):
        if not evt & EPOLLIN:
            msg = 'event happened %d on listen fd.' % evt
            self.error_callback(msg)
            return
        while True:
            try:
                sock, addr = self.listen_sock.sock.accept()
            except socket.error as e:
                if e.errno == socket.errno.EAGAIN:
                    break
                if e.errno == socket.errno.EINTR:
                    continue
                self.error_callback(e.errno)
            self.connect_callback(sock)
