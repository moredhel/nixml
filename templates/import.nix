with (import (builtins.fetchTarball {
  name = "||.Name||-||.Version||";
  url = ||.Url||;
  sha256 = "||.Sha256||";
}) {});
let
  # TODO: move these into an attrSet
  packages = {|| range .PackageSets ||
   # Language: || .Name ||
    ||.Name|| = [
      || range .Modules ||||.|||| end ||
    ];
              || end ||
  };
in
{
  # shell environment
  shell = pkgs.mkShell {
    name = "||.Name||-||.Version||";
    buildInputs = lib.lists.concatLists (lib.attrsets.attrValues packages);
  };


  # package
  pkg = stdenv.mkDerivation {
    name = "||.Name||-||.Version||";
    propogatedBuildInputs = lib.lists.concatLists (lib.attrsets.attrValues packages);
  };
}
