build:
	go build cmd/systelem.go
	go build cmd/monitors/systelem_edac.go

install:
	@install -D systelem $(DESTDIR)/usr/bin/systelem
	@install -D systelem_edac $(DESTDIR)/usr/bin/systelem_edac
	@install -D -m 644 init/systelem.service $(DESTDIR)/etc/systemd/system/systelem.service
	@install -D -m 644 init/systelem_edac.service $(DESTDIR)/etc/systemd/system/systelem_edac.service

clean:
	rm -f systelem systelem_edac
	go clean -i ./...

test:
	go test ./...