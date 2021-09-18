import proxy from 'http2-proxy';

const proxyOptions = {
  hostname: 'localhost', port: 8080
}

/** @type {import("snowpack").SnowpackUserConfig } */
export default {
  mount: {
    // directory name: 'build directory'
    'web/public': '/',
    'web': '/dist',
  },
  plugins: [
    '@snowpack/plugin-svelte',
  ],
  routes: [
    {src: '/api/.*', dest: (req, res) => proxy.web(req, res, proxyOptions).catch(() => res.end())},
    // {"match": "routes", "src": ".*", "dest": "/index.html"},
  ],
  optimize: {
    /* Example: Bundle your final build: */
    // "bundle": true,
  },
  packageOptions: {
    /* ... */
  },
  devOptions: {
    port: 3000,
    open: 'none'
  },
  buildOptions: {
    /* ... */
  },
};
