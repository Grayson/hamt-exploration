DUMMY: main.o
	./main.o

clean:
	rm main.o

main.o: main.c
	cc main.c -g -o main.o