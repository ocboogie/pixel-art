<template>
  <v-card>
    <v-toolbar color="primary" dark flat class="justify-center">
      <v-toolbar-title>Sign Up</v-toolbar-title>
    </v-toolbar>
    <v-card-text>
      <v-form v-model="valid" ref="form">
        <v-text-field
          label="Name"
          name="name"
          v-model="form.name"
          prepend-icon="mdi-account"
          type="text"
          :rules="rules.name"
        />
        <v-text-field
          label="Email"
          name="email"
          v-model="form.email"
          prepend-icon="mdi-email"
          type="text"
          :rules="rules.email"
        />
        <v-text-field
          id="password"
          label="Password"
          name="password"
          v-model="form.password"
          prepend-icon="mdi-lock"
          :rules="rules.password"
          type="password"
        />
        <v-text-field
          id="confirmPassword"
          label="Confirm Password"
          name="confirmPassword"
          v-model="form.confirmPassword"
          prepend-icon="mdi-lock"
          :rules="[passwordConfirmationRule]"
          type="password"
        />
      </v-form>
      <v-alert v-if="error" class="mt-3 mb-0" type="error" dense>
        {{ error }}
      </v-alert>
    </v-card-text>
    <v-card-actions>
      <v-btn
        color="primary"
        x-large
        block
        @click="submit"
        :loading="loading"
        :disabled="!valid"
        >Sign Up</v-btn
      >
    </v-card-actions>
    <div class="text-center pb-3">
      Already have an account?
      <router-link to="/login">Login</router-link>
    </div>
  </v-card>
</template>
<script>
import store from "../store";
import { mapActions } from "vuex";

export default {
  computed: {
    passwordConfirmationRule() {
      return () =>
        this.form.password === this.form.confirmPassword ||
        "Password must match";
    }
  },
  data: () => ({
    loading: false,
    valid: false,
    form: {
      name: "",
      email: "",
      password: "",
      confirmPassword: ""
    },
    rules: {
      name: [
        v => (v && v.length > 2) || "Names must have more than 2 characters",
        v => (v && v.length < 256) || "Names must have less than 256 characters"
      ],
      email: [
        v => Boolean(v) || "E-mail is required",
        v =>
          /(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])/.test(
            v
          ) || "E-mail must be valid"
      ],
      password: [
        v => (v && v.length > 5) || "Password must have more than 5 characters"
      ]
    }
  }),
  watch: {
    "form.password": "validateField"
  },
  methods: {
    ...mapActions(["loggedIn"]),
    async submit() {
      this.loading = true;

      try {
        await this.axios.post("/auth/signUp", this.form);
      } catch (error) {
        this.loading = false;
        this.error = error.response.data.error.message;
        return;
      }
      this.loading = false;
      this.loggedIn();
      this.$router.replace({ name: "home" });
    },
    validateField() {
      this.$refs.form.validate();
    }
  }
};
</script>
