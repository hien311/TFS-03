<template>
  <div id="app">
    <img class="app-bg" src="./assets/pattern-background-desktop.svg" alt="bg">
    <OrderSummary/>
  </div>
</template>

<script>
export default {
  name: 'App',
  data() {
    return { 
    }
  },
  components: {
    OrderSummary: () => import('./components/OrderSummary.vue')
  }
}
</script>

<style >
*{
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

#app {
  font-family: Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  
  background: rgb(224,232,255);
  color: #2c3e50;
  margin: auto;
  position: relative;
  width: 1440px;
  height: 100vh;
}
</style>
