<script>
	import { onMount } from "svelte";
	import auth from "./authService";
	import { isAuthenticated, user } from "./store";
	import NavBar from "./components/NavBar.svelte";

	let auth0Client;

	onMount(async () => {
		auth0Client = await auth.createClient();
		//auth.loginWithPopup(auth0Client);
		isAuthenticated.set(await auth0Client.isAuthenticated());
		user.set(await auth0Client.getUser());
	});

	// $: {
	// 	console.log($user);
	// }
</script>

<main>
	<!-- App Bar -->
	<NavBar
		authState={isAuthenticated}
		userInfo={user}
		authFunc={auth}
		authClient={auth0Client}
	/>
</main>
