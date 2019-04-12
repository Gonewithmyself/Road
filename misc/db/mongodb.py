


import pymongo
from mysql import log
import gridfs
import ctypes

host='192.168.15.93'
port=27017
username="test"
password="123456"
dbname='s1001'

auth = {
    'host':host,
    'username':username,
    'password':password,
    'authSource':dbname,
    'authMechanism':'SCRAM-SHA-1',
}

class Mongo(object):
    def __init__(self, auth, db=None, col=None):
        self.conn = pymongo.MongoClient(**auth)
        self.db = self.conn.get_database(db)
        self.check_collection(col)
        self.fs = gridfs.GridFS(self.db)
        self.fstream = gridfs.GridFSBucket(self.db)
    
    def check_collection(self, name):
        if not name:
            return
        self.col = self.db.get_collection(name)
    
    def find(self, col=None, *args, **kw):
        self.check_collection(col)
        cursor =  self.col.find(*args, **kw)
        return [doc for doc in cursor]
    
    def insert_many(self, docs, col=None):
        self.check_collection(col)
        self.col.insert_many(docs, bypass_document_validation=True)
    
    def insert(self, doc, col=None):
        self.check_collection(name)
        self.col.insert_one(doc)


    @log
    def listdb(self):
        return self.db, self.col
    
    @log
    def list_collection(self):
        return [ x for x in filter(lambda x: x != 'fs.files' and x != 'fs.chunks',
        self.db.list_collection_names())]

    @log
    def list_gridfs(self):
        return self.fs.list()
    
    def fs_insert(self, data, name):
        prev = self.fs.find_one({"filename":name})
        self.fs.put(data, filename=name)
        if prev:
            self.fs.delete(prev._id)
    
    def fs_load(self, filename):
        return self._fs_load(filename).read()
    
    def _fs_load(self, filename):
        return self.fs.find_one({"filename":filename})

# write girdfs with raw str
class FsWriter(object):
    def __init__(self):
        self.dll = ctypes.cdll.LoadLibrary("../test.so")
        self.dll.InitSession(
            cstr('s1001'), 
            cstr("192.168.15.93:27017"), 
            cstr("test"), 
            cstr("123456"))
    
    def write(self, filename, data):
        self.dll.WriteGridFS(
            cstr(filename),
            cstr(data),
        )

def cstr(src):
    return ctypes.c_char_p(src.encode('utf-8'))

def copy():
    auth['host'] = '192.168.15.227'
    src = Mongo(auth, db='s1001')

    auth['host'] = '192.168.15.93'
    dst = Mongo(auth, db='s1001')
    cols = src.list_collection()
    for col in cols:
        docs = src.find(col)
        if not docs:
            print("no data in", col)
            continue
        dst.insert_many(docs, col)
    
    w = FsWriter()
    for col in src.list_gridfs():
        data = src.fs_load(col)
        if not data:
            print('none gridfs', col)
            continue
        w.write(col, data)
        





import json
if __name__ == '__main__':
    # copy()
    # auth['host'] = '192.168.15.93'
    # h = Mongo(auth,db='s1002', col="users")
    # h.list_collection()
    # h.list_gridfs()
    # h.fs_insert(bytes('go to hell', encoding='utf-8'), "test")
    # data = h.fs_load("user.txt")
    # print(type(data), json.loads(data))
    # t = Mongo(auth, db='s1003', col='users')
    # t.col.insert_many(users)

    h = FsWriter()
    h.write("python", "this is data!")
    auth['host'] = '192.168.15.93'
    db0 = Mongo(auth, dbname)

    res = db0.fs_load("python")
    print(res)

    copy()

