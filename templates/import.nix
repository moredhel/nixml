with (import (builtins.fetchTarball {
  name = "||.Name||-||.Version||";
  url = ||.Url||;
  sha256 = "||.Sha256||";
}) {});
let
  # TODO: move these into an attrSet
  # <<<>>>
  || range .PackageSets ||
   # Language: || .Name ||
  ||.Name|| = [
    || range .Modules ||||.||
    || end ||
  ];|| end ||
  # <<<>>>
in
{
  shell = pkgs.mkShell {
    name = "||.Name||-||.Version||";
    buildInputs = lib.lists.concatLists [
      || range .PackageSets |||| .Name ||
      || end ||
    ];
  };
}
