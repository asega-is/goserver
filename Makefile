
DESTDIR	= "/usr/local/bin"

all:
	go build goserver.go
	strip goserver

install: all
	cp goserver $(DESTDIR)


clean:
	rm -fr *~ *\# goserver

uninstall: clean
	rm -fr $(DESTDIR)/goserver


