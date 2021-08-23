<template>
  <main @mouseup="deFontSize">
    <div class="container">
      <Screen :isOff="this.isOff" :isTyping="isTyping"></Screen>
      <Controller :isTyping="isTyping" @change-isOff="changeIsOff" @change-focus="changeFocus" />
      <Buttons @change-focus="changeFocus" />
    </div>
  </main>
</template>

<script>
import Screen from "./components/Screen.vue";
import Buttons from "./components/Buttons.vue";
import Controller from "./components/Controller.vue";
export default {
  name: "App",
  data: function() {
    return {
      isOff: true,
      isTyping: true
    };
  },
  methods: {
    changeIsOff: function() {
      this.$data.isOff = !this.$data.isOff;
    },
    changeFocus(e) {
      this.$data.isTyping = e;
    },
    deFontSize (){
      const screen = document.querySelector('.focused')
      let fontSize = parseFloat(window.getComputedStyle(screen,null).getPropertyValue('font-size'))
      if (screen.offsetWidth > 228) {
        fontSize -= fontSize/12
        screen.style.fontSize = fontSize + 'px'
      }
    } 
  },
  components: { Screen, Buttons, Controller }
};
</script>
<style scoped>
main {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  display: flex;
  justify-content: center;
  font-family: sans-serif;
}

.container {
  margin-top: 200px;
  padding: 10px 6px;
  border-radius: 6px;

  display: flex;
  flex-direction: column;

  background-color: rgb(163, 163, 163);
}
</style>
