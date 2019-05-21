
class Auth(object):
    def __init__(self, *args, **kw):
        for k, v in kw.items():
            setattr(self, k, v)
        
        self.auth = {
            "Db": self.dbname,
            "Source": self.auth_database,
            "User": self.username,
            "Psw": self.password,
            "Host": self.host,
            "Rs": self.replica_set,
        }

# # s1001
# auth1 = Auth(host='dds-t4n63fd619057f541.mongodb.singapore.rds.aliyuncs.com:3717,dds-t4n63fd619057f542.mongodb.singapore.rds.aliyuncs.com:3717',
# replica_set='mgset-300240369',
# dbname = "s1001",
# auth_database="admin",
# username='gameadmin',
# password='I0Cgsng4')  

# # s1002
# auth2 = Auth(host='dds-t4n63fd619057f541.mongodb.singapore.rds.aliyuncs.com:3717,dds-t4n63fd619057f542.mongodb.singapore.rds.aliyuncs.com:3717',
# replica_set='mgset-300240369',
# dbname = "s1002",
# auth_database="admin",
# username='gmtwo',
# password='oj2btg9n')

# s1001
s1001 = Auth(host='dds-t4n63fd619057f541.mongodb.singapore.rds.aliyuncs.com:3717,dds-t4n63fd619057f542.mongodb.singapore.rds.aliyuncs.com:3717',
replica_set='',
dbname = "s1001",
auth_database="admin",
username='radium',
password='KkZz0GF5') 

# s1002
s1002 = Auth(host='dds-t4n63fd619057f541.mongodb.singapore.rds.aliyuncs.com:3717,dds-t4n63fd619057f542.mongodb.singapore.rds.aliyuncs.com:3717',
replica_set='',
dbname = "s1002",
auth_database="admin",
username='radium',
password='KkZz0GF5') 
