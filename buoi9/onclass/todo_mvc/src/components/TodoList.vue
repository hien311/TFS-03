<template>
    <section>
        <input @click="completeAllTasks" type="checkbox" id="toggle-all" class="toggle-all">
        <label for="toggle-all" class="toggle-all">Mark all as complete</label>
        <ul id="todo-list" >
            <li :class = "{ completed: todo.status}"
                v-for="todo in tasks"
                :key="todo.id"
            >
                <Todo :todo="todo"/>
            </li>
        </ul>
    </section>
</template>

<script>
import {mapGetters} from 'vuex'
import {mapActions} from 'vuex'
export default {
    name: "TodoList",
    components: {
        Todo: () => import('./Todo.vue')
    },
    computed: {
        ...mapGetters(['todos', 'todosCompleted', 'todosIncompleted','mode']),
        tasks() {
            switch (this.mode) {
                case 1:
                    return this.todos
                case 2:
                    return this.todosIncompleted
                default:
                    return this.todosCompleted
            }
        }
    },
    methods: {
        ...mapActions(['completeAllTasks',]),
    }
}
</script>

<style scoped>
   #toggle-all {
       opacity: 0;
   }
   
   .toggle-all {
    width: 1px;
    height: 1px;
    border: none;
    position: absolute;
    right: 100%;
    bottom: 100%;
   }

   label {
       width: 60px;
       height: 34px;
       font-size: 0px;
       top: -2px;
       left: 40px;
       transform: rotate(90deg);
       cursor: default;
       background: url("../assets/toggle_img.svg") no-repeat left;
   }
   
    label::before {
        content: "❯";
        font-size: 24px;
        padding: 10px 27px;
        opacity: 0.5;
   }

    li {
        list-style-type: none;
        position: relative;
    }
</style>