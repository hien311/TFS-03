<template>
    <div id="Header">
      <h1>todo</h1>
      <input type="text"  @keyup.enter="addTask" placeholder="What needs to be done?">
    </div>
</template>

<script>
import { mapActions } from 'vuex'
export default {
  name: "Header",
  methods: {
    ...mapActions(['addTask'])
  }
}
</script>

<style scoped>

  h1 {
    position:absolute;
    width: 100%;
    top: -80px;
    font-size: 100px;
    font-weight: 100;
    text-align: center;
    color: rgba(175, 47, 47, 0.15);
     background: rgb(245, 245, 245)
  }

  input {
    position: relative;
    width: 100%;
    padding: 16px 16px 16px 60px;
    border: none;
    background: rgba(0, 0, 0, 0.004);
    box-shadow: rgb(0 0 0 / 3%) 0px -2px 1px inset;
    
    font-size: 24px;
    line-height: 1.4em;
    outline: none;
    color: rgba(0, 0, 0, 0.5)
  }

  input::placeholder {
    opacity: 0.5;
  }
</style>