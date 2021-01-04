<template>
  <div>
    <v-dialog fullscreen hide-overlay v-model="value" key="helper">
      <v-card>
        <v-toolbar>
          <v-btn icon @click="$emit('click')"><v-icon>mdi-close</v-icon></v-btn>
          <v-toolbar-title>帮助</v-toolbar-title>
        </v-toolbar>
        <v-spacer></v-spacer>
        <v-card-title>不围棋简介</v-card-title>
        <v-card-text
          >按字面意思，不围棋就是围棋的反面，围棋要吃子，而不围棋不能吃子。</v-card-text
        >
        <v-card-text
          >我们将棋盘上连在一起的同色棋子看作一个整体，将与其相邻的空格称为“气”，当某一组棋子没有“气”的时候就被“吃”了，这时候被吃者的对手就输了。</v-card-text
        >
        <v-card-text
          >点击左上角的菜单键可以呼出菜单，重新查看本教程或重新开始游戏、切换计算机先手后手等。点击难度标签可以切换难度。</v-card-text
        >
        <v-divider></v-divider>
        <v-card-title>棋子图例</v-card-title>
        <v-row v-for="n in [0,1,2,3]" :key="n" dense>
          <v-col cols="3" offset="1" align-self="center"
            ><stone :id="n" :color="exampleColor[n]"></stone
          ></v-col>
          <v-col cols="8" align-self="center"
            ><v-card-text>{{ exampleText[n] }}</v-card-text></v-col
          >
        </v-row>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import Stone from "./Stone.vue";
import colors from "vuetify/lib/util/colors";
export default {
  components: { Stone },
  props: {
    value: Boolean,
  },
  data: function () {
    return {
      exampleText: [
        "玩家可以下的位置",
        "玩家不可以下的位置",
        "玩家已经下的位置",
        "计算机已经下的位置",
      ],
      exampleColor: [
        colors.green.lighten4,
        colors.grey.lighten3,
        colors.blue.base,
        colors.red.base,
      ],
    };
  },
};
</script>