


import pymongo
from mysql import log

host='192.168.15.93'
port=27017
username="test"
password="123456"
dbname='s1001'

class Mongo(object):
    def __init__(self, **kw):
        self.conn = pymongo.MongoClient(
            host=host,
            username=username,
            password=password,
            port=port,
            authSource=dbname,
            authMechanism='SCRAM-SHA-1')
        
        self.db = self.conn.get_database(dbname)
        self.col = self.db.get_collection("test")
    
    def check_collection(self, name):
        self.col = self.db.get_collection(name)
    
    @log
    def find(self, *args, **kw):
        cursor =  self.col.find(*args, **kw)
        return [doc for doc in cursor]

    @log
    def listdb(self):
        return self.db, self.col
    
    @log
    def list_collection(self):
        return self.db.list_collection_names()

if __name__ == '__main__':
    h = Mongo()
    h.check_collection("users")
    h.find({"_id":21001})
    h.list_collection()
