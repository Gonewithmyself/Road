import socket
import json
from io import BytesIO

# Tag = chr(254)
# NameTag = chr(253)
class Client(object):
    def __init__(self, *args, **kwargs):
        self.sock = socket.socket()
        self.sock.connect(ADDR)
        self.send()
        # self.parse()

    def send(self):
        data = pack_data()
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


# host='dds-t4n63fd619057f541.mongodb.singapore.rds.aliyuncs.com:3717,dds-t4n63fd619057f542.mongodb.singapore.rds.aliyuncs.com:3717'
# replica_set='mgset-300240369'
# dbname = "s1001"
# auth_database="admin"
# username='gameadmin'
# password='I0Cgsng4'

host='dds-j6c5fd84ccef14641.mongodb.rds.aliyuncs.com:3717,dds-t4n63fd619057f542.mongodb.singapore.rds.aliyuncs.com:3717'
replica_set='mgset-13182279'
dbname = "s1001"
auth_database="admin"
username='gmone'
password='nji5y6jm'

# host = "127.0.0.1:27017"
# replica_set=''
# dbname = "s1001"
# auth_database=dbname
# username='test'
# password='123456'

auth = {
    "Db": dbname,
    "Source":auth_database,
    "User": username,
    "Psw": password,
    "Host": host,
    "Rs": replica_set,
}

q = {}
selector = {}
m=''
def pack_data():
    global c,q, selector,m
    if not q:
        q = None

    p = {"C":c,
        "A":auth,
        "Q":q,
        "S":selector,
        "M":m
    }
    # print(p)
    return json.dumps(p).encode('utf-8')


# class 


c = "user_cache"
m = ''
def find():
    global q, c, selector
    c = "users"
    q = {}
    selector = {}
    obj = Client().parse()
    show(obj)

def readfs():
    global c, m
    c = "user_cache"
    m = "fs"
    obj = Client().parse()
    show(obj)

def show(obj):
    if not obj:
        print("not found")
        return
    
    if isinstance(obj, list):
        print(obj[-1])
    else:
        print(obj["41001"])
    

ADDR = ("127.0.0.1", 40000)
if __name__ == "__main__":
    # find()
    readfs()
    pass