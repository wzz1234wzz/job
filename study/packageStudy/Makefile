condict = $(wildcard ./*.go)

run:mainer
	./mainer

mainer:$(condict)
	go build -o mainer $(condict)

.PHONY : clean
clean:
	rm mainer -f
	-rm -rf testdir
