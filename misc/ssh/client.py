
import socket
import os 
import os.path as path

def send(data):
    c = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    c.connect(('localhost', 45000))
    c.send(bytes(data))

def list_files():
    files = os.listdir('.')
    for f in files:
        pass
