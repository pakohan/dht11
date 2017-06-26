CC=gcc
CFLAGS=-c -std=gnu99
LDFLAGS=
SOURCES=$(wildcard *.c)
SOURCES+=$(wildcard Raspberry_Pi/*.c)
SOURCES+=$(wildcard Raspberry_Pi_2/*.c)
OBJECTS=$(SOURCES:.c=.o)
LIBRARY=dht.a
TARGET=dht

$(LIBRARY): $(OBJECTS)
	ar cr $@ $(OBJECTS)

%.o: %.c
	$(CC) $(CFLAGS) $^ -o $@

clean:
	rm -f $(LIBRARY) $(OBJECTS) $(TARGET)
