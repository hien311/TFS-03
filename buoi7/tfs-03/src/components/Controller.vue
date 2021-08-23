<template>
  <div class="controller">
    <div class="controller-item" @click="clean">AC</div>
    <div class="controller-item del" @click="del">Del</div>
    <div class="controller-item" @click="power">Power</div>
  </div>
</template>

<script>
export default {
  name: "Controller",
  props: ["isTyping"],
  methods: {
    clean() {
      const result = document.querySelectorAll(".result");
      result[0].innerHTML = "";
      result[1].innerHTML = "";
      this.$emit("change-focus", true);
      const screen = document.querySelector('.focused')
      screen.style.fontSize = '28px'
    },
    del() {
      const result = document.querySelectorAll(".result");
      if (this.isTyping) {
        if (result[0].innerHTML[result[0].innerHTML.length - 1] == " ") {
          result[0].innerHTML = result[0].innerHTML.slice(0, -3);
        } else {
          result[0].innerHTML = result[0].innerHTML.slice(0, -1);
        }
      }
    },
    power() {
      this.clean()
      this.$emit("change-isOff");
    }
  }
};
</script>
<style scopped>
.controller {
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
}

.controller-item {
  border: 1px solid;
  border-radius: 8px;
  padding: 2px 8px;
  height: 18px;
  font-size: 18px;
  line-height: 18px;
  font-weight: bold;
  color: rgb(218, 186, 186);
  background-color: rgb(240, 37, 23);
}

.controller-item:hover {
  cursor: pointer;
  opacity: 0.8;
}

.controller-item:active {
  opacity: 1;
}

.controller-item + .controller-item {
  margin-left: 6px;
}
</style>
