# Vue 3 Google Auth

## Project setup

### 1. install

```sh
yarn install
```

### 2. set your Google ClientId

create a file in src folder called : `dev.env.js` and add the following content

```javascript
'use strict'
export const env = {
	GAuthClientId: "<ADD YOUR CLIENT ID>",
	BackendURL : "<BACKEND URL>"
}
```

### Compiles and hot-reloads for development
```sh
yarn run dev
```

### Compiles and minifies for production
```sh
yarn run build
```