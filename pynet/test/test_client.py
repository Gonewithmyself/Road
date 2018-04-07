# _*_ coding: utf-8 _*_

import socket
import select
import time
import errno


# from pynet import Client, Coder
from foo.channel import Coder
from foo.client import Client

HOST = '127.0.0.1'
PORT = 8888

class Aclient(Client):
    
    def on_msg_in(self, msg, ch):
        print 'recved' + msg
    
    def on_msg_sent(self, ch):
        print 'msg sent'
    
    def on_error(self, ch):
        print 'error occurred'

if __name__ == '__main__':
    addr = (HOST, PORT)
    client = Aclient(addr, Coder())
    client.start()




