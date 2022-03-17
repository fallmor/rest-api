let
  sources = import ../nix/sources.nix;
  pkgs = import sources.nixpkgs { };
  inherit (pkgs.lib) optional optionals;
in  pkgs.buildGoModule {
    pname = "mygo-app";
    version = "1.0.0";
    src = ./rest;
   # vendorSha256 = "122333";
    
    ldflags = [ "-w" "-extldflags=-static"];
    CGO_ENABLED = 0;
    #comment this line if not building for linux system
    preBuild = '' 
    export GOOS="linux"
    export GOARCH="amd64"
    ''; 
}