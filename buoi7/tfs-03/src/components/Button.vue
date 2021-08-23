<template>
  <div class="buttons-item" @click="press" :class="{ done: styles }">{{ value }}</div>
</template>

<script>
export default {
  name: "Button",
  props: ["value", "styles"],
  data: function() {
    return {};
  },
  methods: {
    press(e) {
      const result = document.querySelectorAll(".result");
      const button = e.target;
      if (
        [" + ", " - ", " x ", " / ", ".", "="].includes(e.target.innerHTML) &&
        result[0].innerHTML == ""
      ) {
      } else {
        if (button.innerHTML != "=") {
          if (
            !(
              result[0].innerHTML[result[0].innerHTML.length - 1] == "." && button.innerHTML == "."
            ) &&
            !(
              result[0].innerHTML[result[0].innerHTML.length - 1] == " " &&
              [" + ", " - ", " x ", " / ", ".", "="].includes(e.target.innerHTML)
            )
          ) {
            result[0].innerHTML += button.innerHTML;
            this.$emit("change-focus", true);
          }
        } else {
          this.done();
        }
      }
    },
    done() {
      this.$emit("change-focus", false);
      const result = document.querySelectorAll('.result')
       let url = "http://localhost:3000/result/" + result[0].innerHTML;
        fetch(url)
            .then((res) => res.json())
            .then((data) => {
                data = Math.round(data * 100) / 100;
                result[1].innerHTML = data;
            })
            .catch((err) => console.log(err));
    }
  }
};
</script>

<style>
.buttons-item {
  text-align: center;
  font-size: 24px;

  width: 23%;
  border: 1px solid;
  border-radius: 6px;
  margin: 2px auto;

  font-weight: 600;
  line-height: 32px;
  background-color: rgb(63, 62, 62);
  color: rgb(218, 186, 186);
}

.buttons-item:hover {
  cursor: pointer;
  opacity: 0.8;
}

.buttons-item:active {
  opacity: 1;
}

.done {
  background-color: rgb(228, 114, 49);
}
</style>
