# mDNS Handler for Caddy

The `mdns_handler` Caddy module is a middleware that allows Caddy to resolve `.local` hostnames using the multicast DNS (mDNS) protocol. This can be useful in local network environments where DNS resolution is not available or not working properly.

To use the `mdns_handler` module, you'll need to follow these steps:

## Step 1: Install Caddy

To use the `mdns_handler` module, you'll need to have [Caddy](https://caddyserver.com/) installed on your system. You can download Caddy from the [official website](https://caddyserver.com/download) or use a package manager on your operating system to install it.

## Step 2: Install the `mdns_handler` module

The `mdns_handler` module is not included with the default Caddy installation, so you'll need to install it separately. You can do this by adding the following line to your Caddyfile:

```
{
order mdns_handler first
}
```


This line tells Caddy to load the `mdns_handler` module first, before any other middleware.

You'll also need to install the module by running the following command:

```
go get github.com/<your-github-username>/<your-module-repo>/mdnshandler
```


This will download and install the `mdnshandler` module to your system.

## Step 3: Configure the `mdns_handler` module

Once the module is installed, you can configure it in your Caddyfile. Here's an example configuration that resolves the `example.local` hostname using the `_http._tcp.local` service:

```
example.local {
    reverse_proxy / http://localhost:8080
    mdns_handler {
        name    example.local
        service _http._tcp.local
    }
}

```


This configuration sets up a reverse proxy for the `/` path to `http://localhost:8080`, and adds the `mdns_handler` middleware to resolve the `example.local` hostname using the `_http._tcp.local` service. You can replace these values with your own server and service settings as needed.

## Step 4: Restart Caddy

After you've configured the `mdns_handler` module, you'll need to restart Caddy for the changes to take effect. You can do this by running the following command:

```
sudo systemctl restart caddy
```


This command restarts the Caddy service on Linux systems. You may need to use a different command or tool to restart Caddy on your system.

## Conclusion

With the `mdns_handler` module installed and configured, your Caddy server should now be able to resolve `.local` hostnames using the mDNS protocol. This can be a useful tool for local network environments where DNS resolution is not working properly.
