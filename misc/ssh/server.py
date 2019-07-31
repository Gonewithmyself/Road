

import socket
import json

TAG = chr(254)
def send():
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.bind(('localhost', 45000))
    s.listen()
    while True:
        c , addr = s.accept()
        print(addr)
        buf = []
        while True:
            d = c.recv(1024)
            if d:
                buf.append(d)
            else:
                c.close()
                break
        msg = bytes().join(buf)
        on_msg(msg)
def on_msg(msg):
    msgs = str(msg).split(TAG)
    for m in msgs:
        try:
            obj = json.loads(m)
        except json.JSONDecodeError:
            print("not json", m)
            continue
            
        try:
            with open(obj.name, "w")as f:
                f.write(obj.data)
        except AttributeError:
            print("unkown json", obj)
            continue


if __name__ == "__main__":
    send()

#####
# nc -l 9999 > xxx.file
# nc localhost 9999 < xxx.file
# ssh -CqtnNf
#   ssh -L 7001:localhost:9999  sshhost
#   listen local 7001 port, data send to 7001, will transefer to sshhost 9999 via ssh turnnel

#   ssh -R 7001:localhost:9999  sshhost
#   sshhost listen 7001 port , data send to 7001, will transefer to localhost 9999 via ssh turnnel

#####


