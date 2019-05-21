import socket
import json
from io import BytesIO
# Tag = chr(254)
# NameTag = chr(253)
class Client(object):
    def __init__(self, *args, **kw):
        self.sock = socket.socket()
        self.sock.connect(kw['ssh'])
        self.send(kw['data'])
        # self.parse()

    def send(self, data):
        self.sock.sendall(data)
        self.sock.shutdown(socket.SHUT_WR)
    
    def parse(self):
        res = self.recv()
        if len(res) == 0:
            return None
        return json.loads(res)
        
    
    def recv(self):
        buf = BytesIO()
        while True:
            data = self.sock.recv(1024)
            if data:
                buf.write(data)
            else:
                self.sock.close()
                break
        return buf.getvalue().decode('utf-8')