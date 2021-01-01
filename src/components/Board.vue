<template>
  <div style="margin-top: 50px;text-align:center">
    <board-line
      v-for="ln in lns"
      :key="ln"
      :lineNum="ln"
      :board="board"
      @click="clickStone"
    ></board-line>
  </div>
</template>

<script>
import BoardLine from "./BoardLine.vue";

export default {
  components: { BoardLine },
  data: function () {
    return {
      board: [],
      lns: [0, 1, 2, 3, 4, 5, 6, 7, 8],
    };
  },
  created: function () {
    for (let i = 0; i < 81; i++) {
      this.board.push({ disabled: false, color: "blue" });
    }
  },
  methods: {
    clickStone: function (event) {
      window.alert("hi from stone "+event);
    },
  },
};
</script>

<style>
</style>