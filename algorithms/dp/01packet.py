

# capacity

class Packet1(object):
    def __init__(self, weight=9, items=[2,2,4,6,3]):
        self.weight = weight
        self.items = items
        self.n = len(self.items)
        self.max = 0
        self.map = [[0 for i in range(self.weight)] for i in range(self.n)]
        print(self.map)
        self.put_into_packet(0, 0)
        print(self.max)
    
    def put_into_packet(self, i, cw):
        if cw == self.weight or i == self.n :
            if self.max < cw:
                self.max = cw
            return
        if self.map[i][cw]:
            return
        self.map[i][cw] = 1
        self.put_into_packet(i+1, cw)
        if cw + self.items[i] <= self.weight:
            print('put itme ', i)
            self.put_into_packet(i+1, cw+self.items[i])

class Packet2(Packet1):
    def __init__(self):
        super().__init__()
        self.
    def put_into_packet(self):

if __name__ == "__main__":
    Packet1()
    pass