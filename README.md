# Angular Build Buildpack

A buildpack that runs `ng build` so you can containerize your Angular apps.

## Building a sample app
Let's build a [sample Angular app](/samples/hello-world-angular) into a container using the Paketo Node Engine, Paketo NPM Install, Angular Build, and Paketo NGINX Buildpacks.

### Package the Angular Build Buildpack into a buildpackage artifact

```
./scripts/package.sh --version <version-number>
```

### Building
```
pack build hello-world-angular --path ./samples/hello-world-angular --buildpack gcr.io/paketo-buildpacks/node-engine --buildpack gcr.io/paketo-buildpacks/npm-install --buildpack </path/to/angular/build/buildpack> --buildpack gcr.io/paketo-buildpacks/nginx
```

### Running
```
docker run -e PORT=8080 -p 8080:8080 -it hello-world-angular
```

### Viewing
Visit `http://localhost:8080/`
