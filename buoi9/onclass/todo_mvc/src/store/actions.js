import axios from "axios";

axios.defaults.baseURL = "http://localhost:3000";

function randomID() {
    return Math.random()
        .toString()
        .substr(2, 5);
}

export default {
    async loadTodos({ commit }) {
        const response = await axios.get("tasks");
        const tasks = await response.data;
        commit("SET_TODOS", tasks);
    },
    async addTask({ commit }, e) {
        if (!e.target.value) {
          return
        }
        let id = parseInt(randomID());

        const newTask = {
            id: id,
            name: e.target.value,
            status: false,
        };
        await axios.post("tasks/add", JSON.stringify(newTask));
        commit("ADD_TASK", newTask);
        e.target.value = "";
    },
    async completeAllTasks({ commit }) {
        await axios.post("tasks/complete-all");
        commit("COMPLETE_ALL_TASKS");
    },
    async clearCompleted({ commit }) {
        await axios.delete("tasks/delete/completed-task");
        commit("CLEAR_COMPLETED");
    },
    async updateTaskStatus({ commit }, task) {
        const newTask = {
          id : task.id,
          name: task.name,
          status: !task.status
        }
        console.log(newTask)
        await axios.put(`tasks/update/${newTask.id}`, JSON.stringify(newTask));
        commit("UPDATE_TASK_STATUS", newTask);
    },
    async removeTask({ commit }, task) {
        await axios.delete(`tasks/delete/${task.id}`);
        commit("REMOVE_TASK", task);
    },
    changeMode({ commit }, mode) {
        commit("CHANGE_MODE", mode);
    },
    async updateTaskName({ commit }, value) {
        const newTask = {
          id: value.task.id,
          name: value.newName,
          status: value.task.status
        }

        await axios.put(`tasks/update/${newTask.id}`, JSON.stringify(newTask));
        commit("UPDATE_TASK_NAME", newTask);
    },
};
