ROOT=..
CC=gcc
CFLAGS= -w -std=gnu99 -o3

PROGS =	test

all:	clean ${PROGS}

test:test.o dlq_list.o main.o sort.o hash.o
	$(CC) $(CFLAGS) -o $@ $^

clean:
	rm -f $(PROGS) $(TEMPFILES) *.o file.hole


