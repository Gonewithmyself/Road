# _*_ coding: utf-8 _*_


from loop import Loop
from channel import Channel, Socket
from connector import Connector

class Client(object):
    def __init__(self, addr, coder=None):
        self.sock = None
        self.addr = addr
        self.coder = coder
        self.loop = Loop()
        self.ch = None
    
    def start(self):
        self.connect()
        self.loop.loop()
    
    def connect(self, retry=3):
        self.sock = Connector(self.addr).connect()
        if not self.sock:
            print 'connect timeout.'
            exit()
        self.ch = Channel(self.sock, self.loop, self.coder)
        self.ch.set_read_callback(self.on_msg_in)
        self.ch.set_write_callback(self.on_msg_sent)
        self.ch.set_error_callback(self.on_error)
        self.ch.set_peer_closed(self.peer_closed)        

    def on_msg_in(self, msg, ch=None):
        pass

    def on_msg_sent(self, ch=None):
        pass
    
    def on_error(self, ch=None):
        pass

    def peer_closed(self, ch=None):
        print 'server offline.'
        self.loop.quit()
    
    def send(self, msg):
        self.ch.send(msg)