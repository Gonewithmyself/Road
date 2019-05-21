import json

def dump(obj):
    print(json.dumps(obj))

def proc_users(res):
    print(res[:2])

def proc_heros(res):
    dump(res)

def proc_level(res):
    res.sort(key=lambda x: x.get('exp'), reverse=True)
    # print(res[0].keys())
    print(res[:10])

    # res.sort(key=lambda x: x.get('power'), reverse=True)
    # write_csv(res[:10])


import csv

def write_csv(rows):
    with open ('xxx.csv', "w+") as f:
        w = csv.writer(f)
        w.writerows(rows)
