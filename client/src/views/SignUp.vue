<template>
  <v-container class="fill-height" fluid>
    <v-row align="center" justify="center">
      <v-col cols="12" sm="8" md="4">
        <SignUp />
      </v-col>
    </v-row>
  </v-container>
</template>
<script>
import SignUp from "../components/SignUp.vue";

export default {
  components: {
    SignUp
  }
};
</script>
