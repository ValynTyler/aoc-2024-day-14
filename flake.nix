{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = { self, nixpkgs }:
  let
    pkgs = nixpkgs.legacyPackages.x86_64-linux;
  in
  {
    devShells.x86_64-linux.default = pkgs.mkShell {
      buildInputs = with pkgs; [ go ];
      # must be run with `--impure` flag to work
      env.GOPATH = "${builtins.getEnv "HOME"}/.cache/go";
    };
  };
}
