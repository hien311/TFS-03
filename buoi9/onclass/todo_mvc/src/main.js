import Vue from 'vue'
import App from './App.vue'
import Vuex from 'vuex'

Vue.use(Vuex)
Vue.config.productionTip = false

const store = new Vuex.Store({
  state: {
    todos: [
    ],
    mode: 2,
  },
  getters: {
    todos: (state) => state.todos,
    todosCompleted: (state) => state.todos.filter(x => x.status === true),
    todosIncompleted: (state) => state.todos.filter(x => x.status === false),
    mode: (state) => state.mode,
    count: (state) => {
      const count = state.todos.filter(x => x.status === false).length
      return count === 1 ? `${count} item left` : `${count} items left`
    },
  },
  actions: {
    completeAllTasks({commit}) {
      commit('completeAllTask')
    },
    clearCompleted({commit}) {
      commit('clearCompleted')
    },
    updateStatusTask({commit}, task) {
      commit('updateStatusTask', task)
    },
    addTask({commit},task) {
      commit('addTask', task)
    },
    removeTask({commit}, task) {
      commit('removeTask', task)
    },
    changeMode({commit}, mode) {
      commit('changeMode', mode)
    },
    updateNameTask({commit},value){
      commit('updateNameTask',value)
    }
  },
  mutations: {
    completeAllTask(state) {
      state.todos.forEach( t => t.status = true)
    },
    clearCompleted(state) {
      state.todos = state.todos.filter(x => x.status === false)
    } 
    ,
    updateStatusTask(state,task) {
      const taskName = task.target.closest('.view').children[1].innerHTML
      const t = state.todos.find(x => x.name === taskName)
      t.status = task.target.checked
    },
    updateNameTask(state, value){
      const t = state.todos.find(x => x.name === value.oldName)
      t.name = value.newName
    },
    addTask(state,task) {
      const newTask = task.target.value
      if (!state.todos.find(x => x.name === newTask) && newTask !== "" ) {
        state.todos.push({name: newTask, status: false})
      }
      task.target.value = ""
    },
    removeTask(state, task) {
      const taskName = task.target.closest('.view').children[1].innerHTML
      const t = state.todos.find(x => x.name === taskName)
      const index = state.todos.indexOf(t)
      state.todos.splice(index,1)
    },
    changeMode(state, mode) {
      state.mode = mode
    },
  }
})

new Vue({
  render: h => h(App),
  store
}).$mount('#app')
