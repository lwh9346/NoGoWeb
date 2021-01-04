<template>
  <div style="margin-top: 50px; text-align: center">
    <board-line
      v-for="ln in lns"
      :key="ln"
      :lineNum="ln"
      :board="board"
      @click="clickStone"
      :disabled="false"
    ></board-line>
    <v-chip style="margin-top: 10px" :color="chipColor">{{ message }}</v-chip>
    <v-spacer></v-spacer>
    <v-chip style="margin-top: 10px" @click="changeDifficulty">{{difficultyMessage[difficulty]}}</v-chip>
    <v-dialog v-model="dialog" max-width="300"
      ><v-card max-width="300">
        <v-card-title>{{ winnerDialogData.title }}</v-card-title>
        <v-card-text>{{ winnerDialogData.text }}</v-card-text>
        <v-card-actions
          ><v-btn @click="restart()" text :color="color.blue.base"
            >重新开始游戏</v-btn
          ></v-card-actions
        >
      </v-card></v-dialog
    >
  </div>
</template>

<script>
import BoardLine from "./BoardLine.vue";
import colors from "vuetify/lib/util/colors";

const axios = require("axios").default;

const stoneColors = {
  base: colors.grey.lighten3,
  player: colors.blue.base,
  computer: colors.red.base,
  playerAllow: colors.green.lighten4,
};

const chipColors = {
  playerRound: colors.blue.lighten3,
  computerRound: colors.yellow.lighten3,
  error: colors.red.lighten3,
  gameOver: colors.purple.lighten3,
};

const messages = {
  playerRound: "现在是你的回合",
  computerRound: "计算机思考中",
  error: "无法连接到服务器",
  gameOver: "游戏结束",
};

export default {
  components: { BoardLine },
  data: function () {
    return {
      board: [],
      lns: [0, 1, 2, 3, 4, 5, 6, 7, 8],
      disabled: false,
      message: messages.playerRound,
      dialog: false,
      winnerDialogData: { title: "title", text: "text" },
      color: colors,
      chipColor: chipColors.playerRound,
      computerFirst: false, //计算机执黑
      difficulty: "easy",
      difficultyMessage:{"easy":"简单","normal":"普通","hard":"困难"}
    };
  },
  created: function () {
    for (let i = 0; i < 81; i++) {
      this.board.push({
        state: 0,
        allow: true,
        color: stoneColors.playerAllow,
      });
    }
  },
  methods: {
    clickStone: function (event) {
      if (this.disabled) {
        return;
      }
      if (!this.board[event].allow) {
        return;
      }
      this.disabled = true; //加锁
      this.chipColor = chipColors.computerRound;
      this.$set(this.board, event, {
        state: 1,
        allow: false,
        color: stoneColors.player,
      });
      this.message = messages.computerRound;
      let req = { board: [], difficulty: this.difficulty };
      for (let index = 0; index < this.board.length; index++) {
        const e = this.board[index];
        req.board.push(e.state);
      }
      axios({
        method: "post",
        url: "/api/nogo",
        data: req,
      })
        .then((resp) => {
          resp = resp.data;
          if (resp.winner != "none") {
            this.winnerDialogData.title = "游戏结束";
            let t = "获胜方是：";
            if (resp.winner == "player") {
              t += "玩家";
            } else {
              t += "计算机";
            }
            this.winnerDialogData.text = t;
            this.dialog = true;
            this.chipColor = chipColors.gameOver;
            this.message = messages.gameOver;
            return;
          }
          let index = resp.x * 9 + resp.y;
          this.$set(this.board, index, {
            state: -1,
            allow: false,
            color: stoneColors.computer,
          });
          for (let i = 0; i < this.board.length; i++) {
            const element = this.board[i];
            if (element.state == 0) {
              if (resp.allow[i]) {
                this.$set(this.board, i, {
                  state: 0,
                  allow: true,
                  color: stoneColors.playerAllow,
                });
              } else {
                this.$set(this.board, i, {
                  state: 0,
                  allow: false,
                  color: stoneColors.base,
                });
              }
            }
          }
          this.disabled = false;
          this.chipColor = chipColors.playerRound;
          this.message = messages.playerRound;
        })
        .catch((error) => {
          console.log(error);
          //还原
          this.$set(this.board, event, {
            state: 0,
            allow: true,
            color: stoneColors.playerAllow,
          });
          this.disabled = false;
          this.chipColor = chipColors.error;
          this.message = messages.error;
        });
    },
    restart: function () {
      for (let i = 0; i < 81; i++) {
        this.$set(this.board, i, {
          state: 0,
          allow: true,
          color: stoneColors.playerAllow,
        });
      }
      this.dialog = false;
      this.chipColor = chipColors.playerRound;
      this.disabled = false;
      this.message = messages.playerRound;
    },
    computerFirstRestart: function () {
      this.restart();
      this.$set(this.board, 1, {
        state: -1,
        allow: false,
        color: stoneColors.computer,
      });
    },
    changeDifficulty:function(){
      switch (this.difficulty) {
        case "easy":
          this.difficulty="normal"
          break;
        case "normal":
          this.difficulty="hard"
          break;
        case "hard":
          this.difficulty="easy"
          break;
        default:
          break;
      }
    }
  },
};
</script>

<style>
</style>