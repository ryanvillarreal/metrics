{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  name = "metrics-env";

  buildInputs = [
    pkgs.go
    pkgs.sqlc
  ];

  shellHook = ''
    echo "Go $(go version) | sqlc $(sqlc --version)"
  '';
}
