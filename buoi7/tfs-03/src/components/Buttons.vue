<template>
    <div class="buttons">
            <Button 
                v-for="(value,index) in values"
                :value="value"
                :styles="!(index-values.length + 1)"
                :key="value"
                @change-focus="changeFocus"
            ></Button>
    </div>
        
</template>

<script>
import Button from "./Button.vue"
export default {
    name: "Buttons",
    data: function (){
        return {
            values:[7,8,9," + ",4,5,6,' - ',1,2,3,' x ',0,'.',' / ','='],
        }
    },
    methods: {
        changeFocus(e){
            this.$emit('change-focus', e)
        }
    },
    components: {Button}
}
</script>

<style scoped>
    .buttons {
        display: flex;
        flex-wrap: wrap;
        margin-top: 10px;
    }
</style>