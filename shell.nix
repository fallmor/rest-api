let
  sources = import ./nix/sources.nix;
  pkgs = import sources.nixpkgs { };
  inherit (pkgs.lib) optional optionals;
   mygoapp = import ./app;

in  pkgs.mkShell{
    buildInputs = [
      sources.go
      pkgs.postgresql
      #pkgs.docker

    ];
    shellHook = ''
      echo "Hello you are running your app in a ${builtins.currentSystem} and the version is "
        ${mygoapp}
        export DB_USERNAME=postgres
        export DB_PASSWORD=postgres
        export DB_TABLE=postgres
        export DB_PORT=5432
        export DB_HOST=localhost
    '';
}