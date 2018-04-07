# _*_ coding: utf-8 _*_

import socket
import select
import time
import errno


# from pynet import Client, Coder
from pynet.channel import Coder, Stdin
from pynet.server import Server


HOST = '127.0.0.1'
PORT = 8888
# print select.EPOLLIN, select.EPOLLOUT, select.EPOLLPRI, select.EPOLLHUP, select.EPOLLERR




class Aserver(Server):
    def __init__(self, addr, coder=None):
        super(Aserver, self).__init__(addr, coder)
        self.stdin = Stdin(self.loop, self.on_stdin)
    
    def on_stdin(self, msg):
        if msg == 'exit':
            self.quit()
    
    def on_msg_in(self, msg, ch):
        print 'recved' + msg
    
    def on_msg_sent(self, ch):
        print 'msg sent'


if __name__ == '__main__':
    addr = (HOST, PORT)
    serv = Aserver(addr, Coder())
    serv.start()






