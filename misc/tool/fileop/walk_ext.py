import os
import sys



def walk():
    root = os.path.realpath('.')
    todo = [root]
    ext = {}
    while not todo:
        curr = todo.pop()
        files = os.listdir(curr)
        for f in files:
            fpath = curr + '/' + f
            if os.path.isdir(fpath):
                todo.append(curr + '/' + f)
                continue
            xx, ex = os.path.splitext(f)
            print(xx, ex)

def do_walk(root, curr):
    files = os.listdir(curr)


if __name__ == "__main__":
    walk()