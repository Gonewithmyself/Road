import time
from agent.auth import *
import agent.query as query
from agent.handlers import *
def proc__xxx(res):
    proc_heros(res)
    pass

def get_account(id):
    res = query.Query(s1001.auth, "users").find(query={"_id":41001})
    print(res)

if __name__ == "__main__":
    # 1556726400 1557331199
    # print(time.localtime(1555268065))
    collection = 'draw_pack'
    q = query.Query(s1001.auth, collection)
    # res = q.find(query={'created':{'$gte':1556726400, '$lte':1557331199}}, selector={'level':1, 'power':1, 'name':1})
    res = q.find(query={'_id':29391001}, selector={})
    proc__xxx(res)

    # get_account(41001)