<template>
    <footer v-flex>
        <span class="todo-count" v-verticalCenter>{{count}}</span>
        <ul class="filters" v-flex>
            <li @click="changeMode(1)" :class="{selected: mode === 1}">All</li>
            <li @click="changeMode(2)" :class="{selected: mode === 2}">Active</li>
            <li @click="changeMode(3)" :class="{selected: mode === 3}">Completed</li>
        </ul>
        <div class="clear-completed" @click="clearCompleted" v-verticalCenter v-if="todosCompleted.length">Clear completed</div>
    </footer>
</template>

<script>
import { mapActions } from 'vuex'
import { mapGetters } from "vuex"
export default {
    name: "Footer",
    computed: {
        ...mapGetters(['count', 'mode','todosCompleted'])
    },
    directives: {
        flex: function(el){
            el.style.display = "flex"
        },
        verticalCenter: function (el) {
            el.style.cssText = `
                position: absolute;
                top: 50%;
                -ms-transform: translateY(-50%);
                transform: translateY(-50%);
            `
        }
    },
    methods: {
        ...mapActions(['clearCompleted','changeMode'])
    }

}
</script>

<style scoped>
    footer::before {
        content: '';
        position: absolute;
        right: 0;
        bottom: 0px;
        left: 0;
        height: 50px;
        overflow: hidden;
        box-shadow: 0 1px 1px rgb(0 0 0 / 20%), 0 8px 0 -3px #f6f6f6, 0 9px 1px -3px rgb(0 0 0 / 20%), 0 16px 0 -6px #f6f6f6, 0 17px 2px -6px rgb(0 0 0 / 20%);
        z-index: 0;
    }

    footer {
        position: relative;
        padding: 10px 15px;
        text-align: center;
        border-top: 1px solid #e6e6e6;
        color: #777;
        font-size: 14px;  
    }

    .todo-count {
        text-align: left;
    }

    .filters {
        z-index: 99;
    }

    .filters li {
        list-style-type: none;
        padding: 3px 7px;
        margin: 0 3px;
        border: 1px solid white;
    }

    .filters li:hover {
        cursor: pointer;
        border: 1px solid rgba(175, 47, 47, 0.1);
    }

    .filters .selected {
        border: 1px solid rgba(175, 47, 47, 0.2) !important;
    }

    .clear-completed {
        right: 15px;
        background: none;
        display: block;
        border:none;

        color: #777;
        text-align: right;
        font-size: 14px;
        z-index: 99;
    }

    .clear-completed:hover {
        text-decoration: underline;
    }
</style>