{
  nixConfig = {
    extra-experimental-features = "nix-command flakes impure-derivations recursive-nix";
    substituters = [ "https://mirrors.tuna.tsinghua.edu.cn/nix-channels/store" ];
  };

  description = "A basic gomod2nix flake";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";
  inputs.gomod2nix.url = "github:nix-community/gomod2nix";
  inputs.gomod2nix.inputs.nixpkgs.follows = "nixpkgs";
  inputs.gomod2nix.inputs.flake-utils.follows = "flake-utils";

  outputs =
    { self
    , nixpkgs
    , flake-utils
    , gomod2nix
    ,
    }:
    let
      version =
        let
          lastModified = builtins.substring 0 8 self.lastModifiedDate or self.lastModified or "19700101";
        in
        toString (self.shortRev or self.dirtyShortRev or lastModified);
    in
    flake-utils.lib.eachDefaultSystem
      (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        gomod2nix-pkgs = gomod2nix.legacyPackages.${system};

        go-template = gomod2nix-pkgs.buildGoApplication {
          pname = "go-template";
          version = version;
          pwd = ./.;
          src = ./.;
          modules = ./gomod2nix.toml;

          ldflags = [ "-s" "-w" "-X github.com/yoshino-s/go-template/common.Version=v${version}" ];
        };

        protoc-gen-grpc-gateway = pkgs.buildGoModule rec {
          pname = "grpc-gateway";
          version = "2.19.1";

          src = pkgs.fetchFromGitHub {
            owner = "grpc-ecosystem";
            repo = "grpc-gateway";
            rev = "v${version}";
            sha256 = "sha256-CdGQpQfOSimeio8v1lZ7xzE/oAS2qFyu+uN+H9i7vpo=";
          };

          vendorHash = "sha256-no7kZGpf/VOuceC3J+izGFQp5aMS3b+Rn+x4BFZ2zgs=";

          nativeBuildInputs = [ pkgs.installShellFiles ];

          subPackages = [ "protoc-gen-grpc-gateway" "protoc-gen-openapiv2" ];
        };

        buildDeps = with pkgs; [ git go_1_22 gnumake ];
        devDeps = with pkgs;
          buildDeps
          ++ [
            nil
            alejandra

            gotools

            golangci-lint
            golines

            (gomod2nix-pkgs.mkGoEnv { pwd = ./.; })
            gomod2nix-pkgs.gomod2nix

            protobuf
            protoc-gen-go
            protoc-gen-go-grpc
            protoc-gen-grpc-gateway
            buf
          ];
      in
      rec {
        packages = {
          inherit go-template;
        };
        packages.default = go-template;

        apps.go-template = flake-utils.lib.mkApp {
          drv = go-template;
        };
        apps.default = apps.go-template;

        devShells.default = pkgs.mkShell {
          buildInputs = devDeps;
          shellHook = ''
            export PATH="$PWD/result/bin:$PATH"
          '';
        };

        checks = {
          format =
            pkgs.runCommand "check-format"
              {
                buildInputs = with pkgs; [
                  go_1_22

                  alejandra

                  golangci-lint
                  golines
                  clang-tools
                ];
              } ''
              cd ${./.}
              ${pkgs.alejandra}/bin/alejandra --quiet . >> $out
              ${pkgs.golines}/bin/golines --max-len=88 -w . >> $out
              HOME=$TMPDIR ${pkgs.golangci-lint}/bin/golangci-lint run --fix --timeout 10m ./... >> $out
              ${pkgs.clang-tools}/bin/clang-format -style="{BasedOnStyle: Google, IndentWidth: 4, AlignConsecutiveDeclarations: true, AlignConsecutiveAssignments: true, ColumnLimit: 0}" -i ./proto/*/*.proto >> $out
            '';
        };
      });
}
