<template>
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
        <a class="navbar-brand" href="#">Template</a>
        <button
            class="navbar-toggler"
            type="button"
            data-toggle="collapse"
            data-target="#navbarSupportedContent"
            aria-controls="navbarSupportedContent"
            aria-expanded="false"
            aria-label="Toggle navigation"
        >
            <span class="navbar-toggler-icon"></span>
        </button>

        <div class="collapse navbar-collapse" id="navbarSupportedContent">
            <ul class="navbar-nav mr-auto">
                <li class="nav-item active">
                    <a class="nav-link" href="#"
                        >Home <span class="sr-only">(current)</span></a
                    >
                </li>
            </ul>

            <div v-if="Vue3GoogleOauth.isInit">
                <button
                    v-if="!Vue3GoogleOauth.isAuthorized"
                    @click="handleClickSignIn"
                    class="btn btn-success"
                >
                    Login
                </button>

                <button
                    v-if="Vue3GoogleOauth.isAuthorized"
                    class="btn btn-outline-danger"
                    @click="handleClickSignOut"
                >
                    Logout
                </button>
            </div>
        </div>
    </nav>
</template>

<script>
import { inject, toRefs } from "vue";
import { env } from "../dev.env";

export default {
    name: "NavBar",

    methods: {
        async handleClickSignIn() {
            try {
                const googleUser = await this.$gAuth.signIn();
                this.googleuser = googleUser;
                if (!googleUser) {
                    return null;
                }

                let idToken = googleUser.getAuthResponse()["id_token"];

                var myHeaders = new Headers();
                myHeaders.append("Content-Type", "application/json");
                myHeaders.append("Access-Control-Allow-Origin", "*");

                var raw = JSON.stringify({
                    token: idToken,
                });

                var requestOptions = {
                    method: "POST",
                    headers: myHeaders,
                    body: raw,
                    redirect: "follow",
                };

                console.log(env.BackendURL);

                fetch(`${env.BackendURL}/auth/authenticate`, requestOptions)
                    .then((response) => response.text())
                    .then((result) => {
                        localStorage.setItem("jwtToken", result);
                    })
                    .catch((error) => {
                        console.log("error", error);
                        alert("some error occured");
                    });
            } catch (error) {
                //on fail do something
                console.log(error);
                return null;
            }
        },

        async handleClickSignOut() {
            try {
                await this.$gAuth.signOut();
                await localStorage.clear();
            } catch (error) {
                console.error(error);
            }
        },
    },

    setup(props) {
        const { isSignIn } = toRefs(props);
        const Vue3GoogleOauth = inject("Vue3GoogleOauth");

        const handleClickLogin = () => {};
        return {
            Vue3GoogleOauth,
            handleClickLogin,
            isSignIn,
        };
    },
};
</script>