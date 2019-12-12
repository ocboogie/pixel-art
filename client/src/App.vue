<template>
  <v-app>
    <v-app-bar app color="primary" dark>
      <div class="d-flex align-center">
        <v-toolbar-title>Pixel Art</v-toolbar-title>
      </div>

      <v-spacer></v-spacer>

      <v-btn v-if="loggedIn" icon>
        <v-icon>mdi-account</v-icon>
      </v-btn>
      <v-btn to="/login" v-else text>
        Login
      </v-btn>
    </v-app-bar>

    <v-content>
      <router-view />
    </v-content>
  </v-app>
</template>

<script>
import HelloWorld from "./components/HelloWorld";
import { mapState } from "vuex";

export default {
  name: "App",
  computed: mapState({
    loggedIn: state => state.auth.loggedIn
  }),

  components: {
    HelloWorld
  }
};
</script>
