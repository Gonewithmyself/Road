# _*_ coding: utf-8 _*_


from loop import Loop, EPOLLOUT, EPOLLIN, EPOLLET
from channel import Socket
import socket
import time

class Connector(object):
    def __init__(self, addr, retry_times=3):
        self.sock = Socket()
        self.fd = self.sock.fd
        self.addr = addr
        self.loop = Loop(2)
        self.loop.set_timeout(2, self.on_timeout)
        self.times = retry_times
        self.retry_time = 0
        self.stat = 0
    
    def on_timeout(self):

        if self.retry_time > self.times:
            self.close()
            return
        self.retry_time += 1
        self.loop.set_timeout(self.retry_time, self.on_timeout)
        if self.sock.connect(self.addr):
            self.on_connect()

    def run_onece(self, fd, evt):
        if evt < 8:
            self.on_connect()
        # else:
        #     print 'connector ', evt
    
    def on_connect(self):
        self.close()
        self.stat = 1
        self.sock.stat = 1
    
    def close(self):
        self.loop.unregister(self.fd)
        self.loop.quit()
    
    def poll(self):
        self.loop.register(self.fd, EPOLLET|EPOLLOUT, self.run_onece)
        self.loop.loop()
    
    def connect(self):
        self.sock.connect(self.addr)
        self.poll()
        if self.stat:
            print 'connected'
            return self.sock
        return None
    
    # def connect1(self):
    #     cnt = 1
    #     while self.times:
    #         self.times -= 1
    #         cnt += 1
    #         if self.sock.connect(self.addr):
    #             self.poll()
    #             break         
    #         else:
    #             time.sleep(cnt)
    #     if self.stat:
    #         print 'connected'
    #         return self.sock
    #     return None

