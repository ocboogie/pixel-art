<template>
  <v-card>
    <v-toolbar color="primary" dark flat class="justify-center">
      <v-toolbar-title>Login</v-toolbar-title>
    </v-toolbar>
    <v-card-text>
      <v-form>
        <v-text-field
          label="Email"
          name="email"
          v-model="form.email"
          prepend-icon="mdi-email"
          type="text"
        />
        <v-text-field
          id="password"
          label="Password"
          name="password"
          v-model="form.password"
          prepend-icon="mdi-lock"
          type="password"
        />
        <v-checkbox
          hide-details
          v-model="form.rememberMe"
          label="Remember me"
        />
      </v-form>
      <v-alert v-if="error" class="mt-3 mb-0" type="error" dense>
        {{ error }}
      </v-alert>
    </v-card-text>
    <v-card-actions>
      <v-btn :loading="loading" color="primary" x-large block @click="submit"
        >Login</v-btn
      >
    </v-card-actions>
    <div class="text-center pb-3">
      Don't have an account?
      <router-link to="/signup">Sign up</router-link>
    </div>
  </v-card>
</template>
<script>
import { mapActions } from "vuex";

export default {
  data: () => ({
    loading: false,
    error: null,
    form: {
      email: "",
      password: "",
      rememberMe: false
    }
  }),
  methods: {
    ...mapActions(["login"]),
    async submit() {
      this.loading = true;
      try {
        await this.login(this.form);
      } catch (error) {
        this.loading = false;
        this.error = error.response.data.error.message;
        return;
      }
      this.loading = false;
      this.$router.replace({ name: "home" });
    }
  }
};
</script>
