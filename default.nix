with (import (builtins.fetchTarball {
  name = "nixml-1.0";
  url = https://github.com/nixos/nixpkgs/archive/85f820d6e41a301a06a44cd72430cce0b3e4e5f5.tar.gz;
  sha256 = "0iwn8lrhdldgdgz5rg7k8h5wavxw5y73j453ji7z84z805falfwi";
}) {});
let
  # TODO: move these into an attrSet
  packages = {
   # Language: dev-pkgs
    dev-pkgs = [
      bat
    ];
              
   # Language: packages
    packages = [
      go
    ];
              
  };
in
{
  # shell environment
  shell = pkgs.mkShell {
    name = "nixml-1.0";
    buildInputs = lib.lists.concatLists (lib.attrsets.attrValues packages);
  };


  # package
  pkg = stdenv.mkDerivation {
    name = "nixml-1.0";
    propogatedBuildInputs = lib.lists.concatLists (lib.attrsets.attrValues packages);
  };
}
