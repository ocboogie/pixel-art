<template>
  <v-container class="fill-height" fluid>
    <v-row align="center" justify="center">
      <v-col cols="12" sm="8" md="4">
        <Proflie />
      </v-col>
    </v-row>
  </v-container>
</template>
<script>
import Proflie from "../components/Profile.vue";

export default {
  components: {
    Proflie
  }
};
</script>