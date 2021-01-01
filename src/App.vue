<template>
  <v-app id="inspire">
    <v-navigation-drawer v-model="drawer" app>
      <!--  -->
    </v-navigation-drawer>

    <v-app-bar app>
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>

      <v-toolbar-title>NoGo Web</v-toolbar-title>
    </v-app-bar>

    <v-main>
      <div style="margin-top:50px">
      </div>
      <board></board>
    </v-main>
  </v-app>
</template>

<script>
import Board from './components/Board.vue';
export default {
  name: "App",

  components: {Board},

  data: () => ({
    drawer:null
  }),
};
</script>
