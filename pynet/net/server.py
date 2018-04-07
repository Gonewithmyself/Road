# _*_ coding: utf-8 _*_


from loop import Loop
from channel import Channel, Socket
from acceptor import Acceptor


class Server(object):
    def __init__(self, addr, coder=None):
        self.loop = Loop()
        self.acceptor = Acceptor(addr, self.loop)
        self.coder = coder
        self.clients = {}
        self.acceptor.set_error_callback(self.fatal_error)
        self.acceptor.set_connect_callback(self.process_new)

    def start(self):
        self.acceptor.listen()
        self.loop.loop()
    
    def fatal_error(self, msg):
        print msg
        self.quit()

    
    def quit(self):
        for fd in self.clients.keys():
            ch = self.clients.pop(fd)
            ch.close()
        self.loop.quit()

    def process_new(self, sock):
        ch = Channel(Socket(sock), self.loop, self.coder)
        ch.set_read_callback(self.on_msg_in)
        ch.set_write_callback(self.on_msg_sent)
        ch.set_error_callback(self.on_error)
        ch.set_peer_closed(self.peer_closed)
        self.clients[ch.fd] = ch
        self.on_connect(ch)

    def on_connect(self, ch=None):
        print 'client {} connect.'.format(ch.peer_addr)
        pass
             
    def on_msg_in(self, msg, ch=None):
        pass

    def on_msg_sent(self, ch=None):
        pass
    
    def on_error(self, ch=None):
        print 'error'
        self.clients.pop(ch.fd)

    def peer_closed(self, ch=None):
        print 'client {} disconnect.'.format(ch.peer_addr)
        self.clients.pop(ch.fd)
