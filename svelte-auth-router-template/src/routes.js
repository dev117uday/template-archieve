import Home from './routes/Home.svelte';
//import Route from './routes/Route.svelte';
import NotFound from './routes/NotFound.svelte';

export default {
    '/': Home,
    // '/auth/:jwt': Auth,
    // The catch-all route must always be last
    '*': NotFound
};
