# _*_ coding: utf-8 _*_


from loop import Loop
from channel import Channel, Socket

class Client(object):
    def __init__(self, addr, coder=None):
        self.sock = Socket()
        self.addr = addr
        self.loop = Loop()
        self.ch = Channel(self.sock, self.loop, coder)
        self.ch.set_read_callback(self.on_msg_in)
        self.ch.set_write_callback(self.on_msg_sent)
        self.ch.set_error_callback(self.on_error)
    
    def start(self):
        self.ch.connect(self.addr)
        self.loop.loop()

    def on_msg_in(self, msg):
        pass
    
    def on_msg_sent(self):
        pass
    
    def on_error(self):
        pass
    
    def send(self, msg):
        self.ch.send(msg)