
import agent.net as net
import json
class Query(object):
    def __init__(self, auth, collection, sshagent=("127.0.0.1", 40000)):
        self.auth = auth
        self.collection = collection
        self.ssh = sshagent
    
    # return an array or none obj
    def find(self, query={}, selector={}, gridfs=False):
        data = self.pack(query, selector, gridfs)
        cli = net.Client(data=data, ssh=self.ssh)
        return cli.parse()
        
    def pack(self, query, selector,gridfs=False):
        p = {"C":self.collection,
        "A":self.auth,
        "Q":query,
        "S":selector,
        }
        if gridfs:
            p['M'] = 'fs'
        return json.dumps(p).encode('utf-8')

