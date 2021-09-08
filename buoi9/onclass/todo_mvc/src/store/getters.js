export default {
    todos: (state) => state.todos,
    todosCompleted: (state) => state.todos.filter(x => x.status === true),
    todosIncompleted: (state) => state.todos.filter(x => x.status === false),
    mode: (state) => state.mode,
    count: (state) => {
      const count = state.todos.filter(x => x.status === false).length
      return count === 1 ? `${count} item left` : `${count} items left`
    },
}