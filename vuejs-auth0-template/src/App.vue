<template>
  <div>
    <NavBar />

    <router-view />
  </div>
</template>

<script>
import NavBar from "./components/NavBar.vue";

export default {
  components: { NavBar },
};
</script>

<!-- <router-link to="/">Home</router-link>|
<router-link to="/about">About</router-link>| -->
