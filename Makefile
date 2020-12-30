.PHONY: copy-abi
copy-abi:
	if [ ! -d abi ]; then mkdir abi ; fi
	cp -r ./indexed-js/src/abi/*.json abi

.PHONY: gen-bindings
gen-bindings:
	if [ ! -d bindings ]; then mkdir bindings ; fi
	if [ ! -d bindings/pool ]; then mkdir bindings/pool ; fi
	abigen --pkg poolbindings --out bindings/pool/bindings.go
