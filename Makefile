.PHONY: build install
DESTDIR=$(CURDIR)/build-root/
GOPATH=$(CURDIR)/build/go-build

build:
	GOPATH=$(GOPATH) go get github.com/smallfish/simpleyaml
	GOPATH=$(GOPATH) go build -o gonurse

install:
	install -d --mode=755 $(DESTDIR)
	install -d --mode=755 $(DESTDIR)/usr/sbin
	install -d --mode=755 $(DESTDIR)/etc/gonurse
	install -v --mode=755 gonurse $(DESTDIR)/usr/sbin/gonurse
	install -v --mode=644 gonurse/nuser.conf $(DESTDIR)/etc/gonurse/gonurse.conf

clean:
	rm --preserve-root -rf $(DESTDIR)
	rm --preserve-root -rf $(GOPATH)
