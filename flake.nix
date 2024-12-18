{
  description = "Advent of Code 2024";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-24.11";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {
        inherit system;
      };
    in with pkgs; {
      devShell = pkgs.mkShell {
        buildInputs = [
          go
        ];
      };
    });
}
