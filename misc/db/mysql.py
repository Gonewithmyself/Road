# _*_ coding: utf-8 _*_

import pymysql
host='127.0.0.1'
port=3306
username="john"
password="since1999"
dbname="test"

'''
CREATE TABLE `user` (
      `id` int(11) NOT NULL AUTO_INCREMENT,
      `name` varchar(16) DEFAULT NULL,
      `password` varchar(45) NOT NULL,
      `role` varchar(45) NOT NULL,
      `salt` char(10) NOT NULL,
      PRIMARY KEY (`id`),
      UNIQUE KEY `name` (`name`)

) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
insert into user set id=1, name='admin', password='2dc4a8162eb8ad6918603d37f1dd2c93', role=1, salt='kmcB';
'''

def log(fn):
    def wrapper(*args, **kw):
        res = fn(*args, **kw)
        print(res)
        return res
    return wrapper

class Db(object):
    def __init__(self):
        self.db = pymysql.connect(
            host=host,
            user=username,
            password=password,
            db=dbname,
            port=port,
            charset='utf8')
        
        self.cursor = self.db.cursor()
    
    
    def do(self, sql, *args):
        return self._do(self.cursor.execute, sql, *args)
    
    def do_many(self, sql, *args):
        return self._do(self.cursor.executemany, sql, *args)

    @log
    def _do(self, fn, sql, *args):
        try:
            fn(sql, *args)
            self.db.commit()
            return self.cursor
        except Exception as e:
            print(e)
            print(sql)
            return "error"
    
    def bulk(self, *args):
        for sql in args:
            try:
                self.cursor.execute(sql)
            except Exception as e:
                print(sql)
                print(e)
                self.db.rollback()
            self.db.commit()
    
    def insert(self, name):
        sql = "insert into user set name=%s, password='2dc4a8162eb8ad6918603d37f1dd2c93', role=1, salt='kmcB'"
        self.do(sql, name)

    
    @log
    def all(self):
        return self.do("select * from user").fetchall()


sqlfmt =  "insert into user set name=%s, password='2dc4a8162eb8ad6918603d37f1dd2c93', role=1, salt='kmcB'"
def insert(db):
    db.do(sqlfmt, "test123")

def insert_many(db):
    values = [(str(i)+"xx",) for i in range(3)]
    db.do_many(sqlfmt, values)

        
if __name__ == '__main__':
    db = Db()
    # db.insert("test")
    # db.all()
    # insert(db)
    insert_many(db)
