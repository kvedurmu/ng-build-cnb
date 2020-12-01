# Angular Build Buildpack

A buildpack that runs `ng build` so you can containerize your Angular apps.

## Building a sample app
Let's build a [sample Angular app](/samples/hello-world-angular) into a container using the Paketo Node Engine, Paketo NPM Install, Angular Build, and Paketo NGINX Buildpacks.

### Package the Angular Build Buildpack into a buildpackage artifact

```
./scripts/package.sh --version <version-number>
```

### Create a Builder
```
pack create-builder angular-builder --config builder.toml
```

### Building
```
pack build hello-world-angular --path ./samples/hello-world-angular --builder angular-builder
```

### Running
```
docker run -e PORT=8080 -p 8080:8080 -it hello-world-angular
```

### Viewing
Visit `http://localhost:8080/`
