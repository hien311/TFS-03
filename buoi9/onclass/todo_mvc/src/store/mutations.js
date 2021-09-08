export default {
    SET_TODOS(state, tasks) {
        state.todos = tasks;
    },
    COMPLETE_ALL_TASKS(state) {
        state.todos.forEach((t) => (t.status = true));
    },
    CLEAR_COMPLETED(state) {
        state.todos = state.todos.filter((x) => x.status === false);
    },
    UPDATE_TASK_STATUS(state, task) {
        let t = state.todos.find((t) => (t.id === task.id));
        t.status = task.status;
    },
    UPDATE_TASK_NAME(state, newTask) {
        let t = state.todos.find((x) => x.id === newTask.id);
        t.name = newTask.name;
    },
    ADD_TASK(state, newTask) {
        state.todos.push(newTask);
    },
    REMOVE_TASK(state, task) {
        const index = state.todos.indexOf(task);
        state.todos.splice(index, 1);
    },
    CHANGE_MODE(state, mode) {
        state.mode = mode;
    },
};
