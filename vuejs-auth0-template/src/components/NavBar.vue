<template>
  <nav class="navbar navbar-expand-lg navbar-light bg-light">
    <div class="container-fluid">
      <router-link style="text-decoration: none; color: inherit" to="/"
        ><div class="navbar-brand">Template</div></router-link
      >
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarSupportedContent"
        aria-controls="navbarSupportedContent"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li class="nav-item"></li>
        </ul>
        <div v-if="!$auth.loading">
          <button
            class="btn btn-success"
            v-if="!$auth.isAuthenticated"
            @click="login"
          >
            Log in
          </button>
          <router-link v-if="$auth.isAuthenticated" to="/profile"
            ><button class="btn btn-primary">Profile</button></router-link
          >
          &nbsp;
          <button
            class="btn btn-danger"
            v-if="$auth.isAuthenticated"
            @click="logout"
          >
            Log out
          </button>
        </div>
      </div>
    </div>
  </nav>
</template>

<script>
// @ is an alias to /src

export default {
  name: "navbar",
  methods: {
    login() {
      this.$auth.loginWithRedirect();
    },

    logout() {
      this.$auth.logout({
        returnTo: window.location.origin,
      });
    },
  },
};
</script>
