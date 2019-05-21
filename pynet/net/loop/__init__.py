import select
class Loop(object):
    def __init__(self, timeout=-1):
        self.eph = select.epoll()
        self.fd_map = {}
        self._quit = False
        self.timeout = timeout
        self.hand_timeout = None
    
    def __del__(self):
        self.eph.close()
    
    def register(self, fd, event, handler):
        if not self.fd_map.get(fd):
            self.eph.register(fd, event)
            self.fd_map[fd] = handler
        else:
            eph.modify(fd, event)
    
    def unregister(self, fd):
        if not self.fd_map.get(fd):
            return
        else:
            self.eph.unregister(fd)
            self.fd_map.pop(fd)
    
    def set_timeout(self, interval, func=None):
        self.timeout = interval
        self.hand_timeout = func
    
    def quit(self):
        self._quit = True
    
    def loop(self):

        select.select()
        count = 0
        while not self._quit:
            ret = self.eph.poll(self.timeout)
            if not ret:
                if self.hand_timeout:
                    self.hand_timeout()
                continue
            count += 1
            print '%-5d, wakeup evt %s' % (count, ret)
            for fd, evt in ret:
                self.fd_map[fd](fd, evt)