{
  description = "gobat - Go 1.24 dev shell";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05";
  };

  outputs = { self, nixpkgs, ... }: let
    system = "x86_64-linux";
    pkgs = import nixpkgs { inherit system; };

    go = pkgs.go_1_24;  # Go 1.24.x series
  in {
    devShells.${system}.default = pkgs.mkShell {
      packages = [ go ];

      shellHook = ''
        export GOCACHE=$PWD/.cache/go-build
        export GOPATH=$PWD/.gopath
        export GOMODCACHE=$PWD/.gopath/pkg/mod
        export PATH=$GOPATH/bin:$PATH

        echo "🦫 Go dev shell ready (Go version: $(go version))"
      '';
    };
  };
}
