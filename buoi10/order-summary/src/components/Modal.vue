<template>
  <div class="modal" ref="yo">
    <div class="modal-above">
      <img src="../assets/success_tick.svg" alt="ok">
    </div>
    <div class="modal-body">
      <h2>Success!</h2>
      <button @click="close">Enjoy</button>
    </div>
  </div>
</template>

<script>
export default {
    name: 'Modal',
    methods: {
      close: function() {
        this.$emit('close-modal')
      }
    }

}
</script>

<style scoped>
  .modal {
    position: fixed;
    top: 100px;
    left: 50%;
    transform: translateX(-50%);

    background: whitesmoke;
    width: 400px;
    height: 240px;
    border-radius: 16px;
    overflow: hidden;
    box-shadow: 5px 5px 8px rgba(107, 107, 107,0.5);
  }

  .modal .modal-above {
    height: 110px;
    background: rgb(119 210 233 / 64%);
  }

  .modal-above img {
    height: 110px;
  }

  .modal-body h2 {
    margin-top: 20px;
    font-size: 30px;
    color: #41ad49
  }

  .modal-body button {
    outline: none;
    border: none;
    border-radius: 32px;

    font-size: 22px;
    font-weight: 700;
    letter-spacing: 1px;
    color: white;
    background-color: rgb(109, 120, 219);
    
    width: 300px;
    height: 40px;
    margin-top: 10px;
    transition: all .5s ease;
  }

  .modal-body button:hover {
    background: yellowgreen ;
  }
</style>