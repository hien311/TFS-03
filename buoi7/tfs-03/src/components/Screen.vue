<template>
  <div :class="{screen: true, display: isOff }">
           <div class="result" :class={focused:isTyping}></div>
           <div class="result" :class={focused:!isTyping}></div>
   </div>
</template>

<script>
export default {
    name: "Screen",
    props: ['isOff', 'isTyping'],
}
</script>

<style scoped>
.screen {
    margin-bottom: 10px;
    min-height: 80px;
    border-radius: 4px;
    background-color: rgb(218, 204, 186);

    font-family: arial, sans-serif;
    transition: background 300ms linear;
}

.display {
    background-color: black;
}

.result {
    height: 30%;
    line-height: 24px;
    font-size: 16px;
    padding: 0 4px;
    text-align: right;  
    transition: all 200ms linear;
}

.focused {
    height: 70%;
    font-size: 28px;
    line-height: 56px;
}

</style>