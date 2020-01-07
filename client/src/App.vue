<template>
  <v-app>
    <template v-if="!loading">
      <v-app-bar app color="primary" dark>
        <div class="d-flex align-center">
          <v-toolbar-title>Pixel Art</v-toolbar-title>
        </div>

        <v-spacer></v-spacer>

        <v-menu open-on-hover v-if="loggedIn" down offset-y>
          <template v-slot:activator="{ on }">
            <v-btn icon v-on="on">
              <v-icon>mdi-account</v-icon>
            </v-btn>
          </template>
          <v-list width="150">
            <v-list-item to="/profile">
              <v-list-item-title>Profile</v-list-item-title>
            </v-list-item>
            <v-list-item @click="logout">
              <v-list-item-title>Logout</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
        <v-btn to="/login" v-else text>
          Login
        </v-btn>
      </v-app-bar>

      <v-content>
        <router-view />
      </v-content>
    </template>
  </v-app>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  name: "App",
  data: () => ({ loading: true }),
  computed: mapGetters(["loggedIn"]),
  async created() {
    try {
      await this.fetchMe();
    } catch {
      // This most likely means that we are logged out which is fine
    }
    this.loading = false;
  },
  methods: mapActions(["logout", "fetchMe"])
};
</script>
