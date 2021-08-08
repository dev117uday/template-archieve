<template>
  <div>
    <HelloWorld />

    <div v-if="$auth.isAuthenticated">
      {{ saveUser() }}
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from "@/components/HelloWorld.vue";

let savebool = false;

export default {
  name: "home",
  components: { HelloWorld },
  methods: {
    saveUser: function () {
      if (this.$auth.user.email != undefined && !savebool) {
        savebool = true;
        let data = this.$auth.user.sub.split("|");

        var myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");

        var raw = JSON.stringify({
          email: this.$auth.user.email,
          name: this.$auth.user.name,
          gid: data[1],
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
    },
  },
};
</script>
