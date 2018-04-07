# _*_ coding: utf-8 _*_

import socket
import time
import select
import os, sys, fcntl

EPOLLIN = select.EPOLLIN
EPOLLPRI = select.EPOLLPRI
EPOLLOUT = select.EPOLLOUT
EPOLLET = select.EPOLLET
EPOLLALL = EPOLLIN|EPOLLOUT|EPOLLET|EPOLLPRI
EPOLLRD = EPOLLIN|EPOLLET|EPOLLPRI
EPOLLWR = EPOLLOUT|EPOLLET

class Socket(object):
    def __init__(self, sock=None, protocol=socket.AF_INET):
        if not sock:
            sock= socket.socket(protocol, socket.SOCK_STREAM)
            self.stat = 0
        else:
            self.stat = 3
        self.sock = sock
        self.sock.setblocking(False)
        self.fd = sock.fileno()
    
    def __del__(self):
        self.close()
    
    def close(self):
        if self.stat > 0:
            self.sock.close()
            self.stat = 0

    def connect(self, addr, times=3):
        if self.stat > 0:
            return self.stat
        try:
            self.sock.connect(addr)
            self.stat = 1
        except socket.error as e:
            if e.errno == socket.errno.EISCONN:
                self.stat = 1   
        return self.stat
    
    def listen(self, addr, num=5):
        if self.stat != 0:
            return
        self.sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        self.sock.bind(addr)
        self.sock.listen(num)
        self.stat = 2
        print 'listening...'

class Channel(object):
    def __init__(self, sock, loop, coder=None):
        self.sock = sock
        self.fd = sock.fd
        self.loop = loop
        self.coder = coder
        self.rbuf = ''
        self.wbuf = ''
        self.read_callback = None
        self.peer_closed = None
        self.register()
        self._peer_addr = None

    def set_peer_closed(self, func):
        self.peer_closed = func
    
    def set_read_callback(self, func):
        self.read_callback = func
    
    def set_write_callback(self, func):
        self.write_callback = func
    
    def set_error_callback(self, func):
        self.error_callback = func

    @property
    def peer_addr(self):
        if self._peer_addr:
            return self._peer_addr
        try:
            self._peer_addr = self.sock.sock.getpeername()
        except socket.error as e:
            print e.message
        return self._peer_addr
    
    def send(self, msg):
        if self.coder:
            msg = self.coder.pack(msg)
        self.wbuf += msg
        if self.handle_write():
            self.register()

    def recv(self):
        if self.coder:
            ret = self.coder.unpack(self.rbuf)
            if not ret:
                return None
            msg = self.rbuf[:ret]
            self.rbuf = self.rbuf[ret+1:]
        else:
            msg = self.rbuf
            self.rbuf = ''
        return msg

    def handle_write(self):
        sock = self.sock.sock
        send_len = len(self.wbuf)
        if send_len == 0:
            return
        try:
            sent = sock.send(self.wbuf)
        except socket.error as e:
            if e.errno == socket.errno.EAGAIN or \
            e.errno == socket.errno.EINTR:
                return
            self.handle_error(e.message)
            return
        if sent < send_len:
            self.wbuf = self.wbuf[sent:]
            return 1
        self.write_callback(self)
    
    def handle_read(self):
        sock = self.sock.sock
        while True:
            try:
                buf = sock.recv(1024)
            except socket.error as e:
                if e.errno == socket.errno.EAGAIN:
                    break
                if e.errno == socket.errno.EINTR:
                    continue
                self.handle_error(e.message)
                return
            if not buf:
                self.close()
                if self.peer_closed:
                    self.peer_closed(self)
                break
            self.rbuf +=buf
        
        msg = self.recv()
        if msg:
            self.read_callback(msg, self)

    def handle_error(self, msg=None):
        if msg:
            print msg
        self.close()
        self.error_callback(self) 

    def close(self):
        self.unregister()
        self.sock.close()
    
    def disable_read(self):
        self.register(EPOLLWR)
    
    def disable_write(self):
        self.register(EPOLLWR)
    
    def enable_all(self):
        self.register()
    
    def register(self, evt=EPOLLALL):
        self.loop.register(self.sock.fd, evt, self.run_once)
    
    def unregister(self):
        self.loop.unregister(self.sock.fd)
    
    def run_once(self, fd, evt):
        if evt & EPOLLOUT:
            self.handle_write()
        if evt | (EPOLLIN |EPOLLPRI):
            self.handle_read()
        if evt > 7:
            self.handle_error()

class Coder(object):
    def __init__(self):
        self.tail = chr(254)
    
    def pack(self, msg):
        return msg + self.tail
    
    def unpack(self, msg):
        pos = msg.find(self.tail)
        if pos == -1:
            return None
        return pos

def set_noblocking(fd):
    flag = fcntl.fcntl(fd, fcntl.F_GETFL)
    fcntl.fcntl(fd, fcntl.F_SETFL, flag|os.O_NONBLOCK)

class Stdin(object):
    def __init__(self, loop, callback):
        self.buf = ''
        self.fd = sys.stdin.fileno()
        set_noblocking(self.fd)
        self.loop = loop
        loop.register(self.fd, 3, self.handle_read)
        self.cb = callback
    
    def __del__(self):
        self.close()
    
    def close(self):
        self.loop.unregister(self.fd)

    def handle_read(self, fd, evt):
        while True:
            try:
                buf = sys.stdin.read(1024)
            except IOError as e:
                if e.errno == 11:
                    break
                self.close()
            self.buf += buf
        self.cb(self.buf.strip('\n'))
        self.buf = ''
