# Svelte App with Auth and Routing

![Svelte](/assets/svelte_js.png)
![Auth0](/assets/auth0.png)

This is a project template for [Svelte](https://svelte.dev) apps that includes [svelte-spa-router](https://github.com/italypaleale/svelte-spa-router) for client-side routing, and Rollup as bundler and [Auth0](https://auth0.com/) for authentication including login, logout

This template is based on the official template for Svelte: [sveltejs/template](https://github.com/sveltejs/template).

*Note that you will need to have [Node.js](https://nodejs.org) installed.*

## About svelte-spa-router

svelte-spa-router is a client-side router for Svelte 3 apps that leverages hash-based routing (i.e. stores the current view in the URL after the `#` symbol).

You can read more about the router, and the reasons why you might want to use hash-based routing (or not), in the [documentation](https://github.com/italypaleale/svelte-spa-router).

## Get started

**Install the dependencies…**

```bash
cd svelte-app
npm install
```

**Setup Auth Config**
- Create a auth config file
```bash
touch auth_config.js

## add the following content

const config = {
	domain: "",
	clientId: ""
};

export default config;
```

…then start [Rollup](https://rollupjs.org):

```bash
npm run dev
```

Navigate to [localhost:5000](http://localhost:5000). You should see your app running. Edit a component file in `src`, save it, and reload the page to see your changes.

