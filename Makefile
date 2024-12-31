.PHONY: wire
wire:
	wire ./cmd/admin/wire
	wire ./cmd/front/wire

.PHONY: goconvey
goconvey:
	@cd test && goconvey -port 5555
