<script>
	export let authState;
	export let userInfo;
	export let authFunc;
	export let authClient;

	function login() {
		authFunc.loginWithPopup(authClient).then(()=>{
			saveUser()
		})
	}

	function logout() {
		authFunc.logout(authClient);
	}

	function saveUser(){
		setTimeout(()=>{},1000)
		var myHeaders = new Headers();
		myHeaders.append("Content-Type", "application/json");

		var raw = JSON.stringify({
			email: $userInfo["email"],
			name: $userInfo["name"],
			gid: $userInfo["sub"],
		});

		var requestOptions = {
			method: "POST",
			headers: myHeaders,
			body: raw,
			redirect: "follow",
		};

		fetch("http://localhost:8080/user", requestOptions)
			.then((response) => response.text())
			.then((result) => console.log(result))
			.catch((error) => console.log("error", error));
	}

	// $: {
	// 	console.log($userInfo);
	// }
</script>

<nav class="navbar navbar-expand-lg navbar-light bg-light">
	<a class="navbar-brand" href="/">Template</a>
	<button
		class="navbar-toggler"
		type="button"
		data-toggle="collapse"
		data-target="#navbarSupportedContent"
		aria-controls="navbarSupportedContent"
		aria-expanded="false"
		aria-label="Toggle navigation"
	>
		<span class="navbar-toggler-icon" />
	</button>

	<div class="collapse navbar-collapse" id="navbarSupportedContent">
		<ul class="navbar-nav mr-auto">
			<!-- <li class="nav-item active">
				<a class="nav-link" href="/"
					>Home <span class="sr-only">(current)</span></a
				>
			</li> -->
		</ul>
		{#if $authState}
			<span class="text-black">{$userInfo.name}</span>
		{:else}<span>&nbsp;</span>{/if}
		<br /> &nbsp;
		{#if $authState}
			<button class="btn btn-danger" href="/" on:click={logout}
				>Log Out</button
			>
		{:else}
			<button class="btn btn-success" href="/#" on:click={login}
				>Log In</button
			>
		{/if}
	</div>
</nav>

<br>

<!-- <button class="btn" on:click={saveUser}>Click Me</button> -->