<template>
    <div>
        <div v-if="!Vue3GoogleOauth.isInit">
            <div class="alert alert-danger" role="alert">
                Please Enable Cookie to login
            </div>
        </div>
    </div>
</template>

<script>
import { inject, toRefs } from "vue";

export default {
    name: "HelloWorld",

    data() {
        return {
            user: "",
        };
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