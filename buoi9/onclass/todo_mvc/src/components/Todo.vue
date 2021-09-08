<template>
    <div class="view" :data-id=todo.id
        @mouseover="xOn"
        @mouseout="xOff"
    >
        <input @click="updateTaskStatus(todo)" type="checkbox"  class="toggle" :checked="todo.status">
        <label @dblclick="editLabel" :class="{completed: todo.status, editing: isEdit}">{{todo.name}}</label>
        <button :class="{destroy:true, editing: isEdit}" @click="removeTask(todo)"></button>
        <input type="text" @keydown.enter="saveEdit" :class="{edit: isEdit, editing: !isEdit}" v-focus>
    </div>
</template>

<script>
import { mapActions } from "vuex"
export default {
    name: "Todo",
    data() {
        return {
            isEdit: false,
        }
    },
    created: function () {
        window.addEventListener('mousedown', this.saveEditMouse)
    },
    destroyed: function() {
        window.removeEventListener('mousedown',this.saveEditMouse)
    },
    props: ['todo'],
    methods: {
        ...mapActions(['updateTaskStatus', 'removeTask',"updateTaskName"]),
        xOn(e) {
            const destroy = e.target.closest(".view").children[2]
            destroy.innerHTML = "×"
        },
        xOff(e) {
            const destroy = e.target.closest(".view").children[2]
            destroy.innerHTML = ""
        },
        editLabel(e) {
            this.isEdit = true
            const input = e.target.closest('.view').lastChild
            this.oldTaskName = e.target.innerHTML
            input.value = e.target.innerHTML
        },
        saveEdit(e) {
            this.isEdit = false
            this.updateTaskName({task: this.todo, newName: e.target.value})
        },
        saveEditMouse(e) {
            const classList = e.target.classList
            if (classList.value !== 'edit' && this.isEdit === true) {
                const newName = document.querySelector('.edit').value
                this.updateTaskName({task: this.todo, newName: newName})
                this.isEdit = false
            }
        }
    },
    directives: {
        focus: function(el) {
            el.focus()
        } 
    }
}
</script>

<style scoped>
    .view {
        position: relative;
        word-break: break-all;
        line-height: 1;
        transition: color 0.4;
    }

    .toggle {
        position: absolute;
        opacity: 0;
        top: 0;
        left: 0;
        appearance: none;
        background: none;
        border: none;  
        height: 40px;
        width: 40px;
        margin: 10px 0
    }

    label {
        font-size: 24px;
        border-bottom: 1px solid rgb(237, 237, 237) ;
        display: block;
        padding: 15px 15px 15px 60px;
        line-height: 1.2;
        transition: color 0.4s ease;
        background: url("../assets/toggle_img.svg") no-repeat left;
    }

    input:checked + label{
        background: url("../assets/toggle_done_img.svg") no-repeat left;
    }

    .editing {
        display: none;
        
    }

    label.completed {
        color: rgb(224, 218, 218);
        text-decoration:line-through;
    }

    .edit {
        font-size: 24px;
        color: #777;
        line-height: 1.2;
       
        border: none;
        width: calc(100% - 45px);
        padding: 14px 15px;
        margin-left: 45px;

        outline: none;
        box-shadow: 1px 1px 4px #777 inset,1px 1px 4px #777;
        transition: color 0.4s ease;
    }

    .destroy {
        appearance: none;
        border: none;
        background: none;
        
        position: absolute;
        top: 0;
        right: 0;
        
        width: 40px;
        height: 40px;
        margin-top: 8px;

        font-size: 30px;
        color: #cc9a9a;
    }

    .destroy:hover {
        color: #af5b5e;
    }
</style>