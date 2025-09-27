# Palettify

Extract colors from any image and create beautiful palettes in less than 5 seconds!

The flow is simple: upload image -> extract colors -> copy the colors RGBA codes or the generated CSS code!

## Install and Launch

> _If you are not interested in replicating this application but you want to use it, you can head to the [live demo](#)_

### From source code

To install and build the application from source code, you first need to clone it:

```bash
git clone https://github.com/AstraBert/palettify
cd palettify
```

Then you need to build the application with the following command (Go 1.24.5+ is required for the build):

```bash
go build -tags netgo -ldflags '-s -w' -o palettify
```

Then you can run the application:

```bash
./palettify
```

> [!IMPORTANT]
> If you wish to change something in the frontend of the  application, you need to do so by modifying the `.templ` files in the `templates` folder. Once you are done, make sure to run `templ generate` to update the templates (you might need to [install Templ](https://templ.guide/quick-start/installation))

### With `npm` (recommended)

You can directly install the application binary from NPM, using a simple:

```bash
npm install @cle-does-things/palettify
```

Note that, with this installation, you cannot customize the application itself.

Once the binary is installed, you can run it with:

```bash
palettify
```

### With `go`

You can directly install the application binary from GitHub, using a simple:

```bash
go install github.com/AstraBert/palettify@latest
```

You need Go 1.24.5+ for this operation to be successfull.

Once the binary is installed, you can run it with:

```bash
palettify
```

### With Docker

#### Pull from GitHub Registry

You can pull and run the Docker image from the GitHub Registry in the following way:

```bash
docker pull ghcr.io/AstraBert/palettify:main # use the main tag to have the version updated to the latest commit to main, otherwise use a version tag (v0.1.0 e.g.)
docker run -p 8000:8000 ghcr.io/AstraBert/palettify:main
```

#### Build and Run Locally

To build the Docker image locally, you first need to clone the GitHub repository:

```bash
git clone https://github.com/AstraBert/palettify
cd palettify
```

Then you can launch the build command:

```bash
docker build . -t username/imagename:tag
```

Once you are done, you can run it with:

```bash
docker run -p 8000:8000 username/imagename:tag
```

## Contributing

Contributions (both for the blog and for the source code) are more than welcome! You can find a detail contribution guide [here](./CONTRIBUTING.md).

## License

This project is distributed under [MIT license](./LICENSE).

---

_Built with love and [Templ](https://templ.guide), [HTMX](https://htmx.org) and [AlpineJS](https://alpinejs.dev)_
