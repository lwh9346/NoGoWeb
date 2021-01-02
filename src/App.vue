<template>
  <v-app id="inspire">
    <v-navigation-drawer v-model="drawer" app>
      <v-list-item>
        <v-list-item-content
          ><v-list-item-title class="title"
            >操作</v-list-item-title
          ></v-list-item-content
        >
      </v-list-item>
      <v-divider></v-divider>
      <v-list nav>
        <v-list-item
          @click="
            $refs.gameboard.restart();
            drawer = !drawer;
          "
        >
          <v-list-item-icon><v-icon>mdi-restart</v-icon></v-list-item-icon>
          <v-list-item-content
            ><v-list-item-title
              >重新开始游戏</v-list-item-title
            ></v-list-item-content
          >
        </v-list-item>
        <v-list-item
          @click="
            $refs.gameboard.computerFirstRestart();
            drawer = !drawer;
          "
        >
          <v-list-item-icon
            ><v-icon>mdi-desktop-classic</v-icon></v-list-item-icon
          >
          <v-list-item-content
            ><v-list-item-title
              >由计算机先手</v-list-item-title
            ></v-list-item-content
          >
        </v-list-item>
        <v-list-item @click="helperShow()">
          <v-list-item-icon><v-icon>mdi-help</v-icon></v-list-item-icon>
          <v-list-item-content
            ><v-list-item-title>帮助</v-list-item-title></v-list-item-content
          >
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar app>
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>

      <v-toolbar-title>NoGo Web</v-toolbar-title>
    </v-app-bar>

    <v-main style="text-align: center">
      <div style="margin-top: 50px"></div>
      <board ref="gameboard"></board>
      <helper @click="helperShow()" :value="helper"></helper>
    </v-main>
  </v-app>
</template>

<script>
import Board from "./components/Board.vue";
import Helper from "./components/Helper.vue";
export default {
  name: "App",

  components: { Board, Helper },

  data: () => ({
    drawer: false,
    helper: false,
  }),
  methods: {
    helperShow: function () {
      this.helper = !this.helper;
    },
  },
  mounted:function(){
    document.title="NoGo Web"
  }
};
</script>
