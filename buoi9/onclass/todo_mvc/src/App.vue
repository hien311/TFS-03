<template>
  <div id="app">
    <Header/>
    <TodoList/>
    <Footer v-if="this.$store.state.todos.length"/>
  </div>
</template>

<script>
export default {
  name: 'App',
  components: {
    Header: () => import('./components/Header.vue'),
    TodoList: () => import('./components/TodoList.vue'),
    Footer: () => import('./components/Footer.vue'),

  }
}
</script>

<style>
*{
  padding: 0;
  margin: 0 auto;
  box-sizing: border-box;
}

body {
    font: 300 14px / 1.4em "Helvetica Neue", Helvetica, Arial, sans-serif;
    background: rgb(245, 245, 245);
    color: rgb(77, 77, 77);
    min-width: 230px;
    max-width: 550px;
    margin: 0px auto;
    -webkit-font-smoothing: antialiased;
}

#app {
  background: rgb(255, 255, 255);
  max-width: 550px;
  margin: 160px 0 40px;
  position: relative;
  box-shadow: rgb(0 0 0 / 20%) 0px 2px 4px 0px, rgb(0 0 0 / 10%) 0px 25px 50px 0px;
}
</style>
