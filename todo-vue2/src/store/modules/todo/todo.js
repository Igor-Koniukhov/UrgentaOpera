import Localbase from "localbase";
let db = new Localbase("db");
db.config.debug = false;
const todoModule = {
    namespaced: true,
    state: {
        appTitle: process.env.VUE_APP_TITLE,
        search: null,
        board: {
            list: 1,
            title: "",
            backgroundColor: "",
        },
        boards: [],
        defaultBoards: [
            {
                list: 1,
                title: "Created",
                backgroundColor: "#fffff",
                color: "green",
            },
        ],

        tasks: [],
        addBoard: false,
        sorting: false,
        boardSorting: false,
        snackbar: {
            show: false,
            text: "",
        },
    },
    mutations: {
        addTask(state, newTask) {
            state.tasks.push(newTask);
        },
        doneTask(state, id) {
            let task = state.tasks.filter((task) => task.id === id)[0];
            task.done = !task.done;
        },
        deleteTask(state, id) {
            state.tasks = state.tasks.filter((task) => task.id !== id);
        },
        deleteBoard(state, list) {
            state.boards = state.boards.filter((board) => board.list !== list);
        },
        showSnackbar(state, text) {
            let timeout = 0;
            if (state.snackbar.show) {
                state.snackbar.show = false;
                timeout = 300;
            }
            setTimeout(() => {
                state.snackbar.show = true;
                state.snackbar.text = text;
            }, timeout);
        },
        hideSnackBar(state) {
            state.snackbar.show = false;
        },
        updateTasktitle(state, payload) {
            let task = state.tasks.filter((task) => task.id === payload.id)[0];
            task.title = payload.title;
        },
        updateTaskDueDate(state, payload) {
            let task = state.tasks.filter((task) => task.id === payload.id)[0];
            task.dueDate = payload.dueDate;
        },
        setTasks(state, tasks) {
            state.tasks = tasks;
        },
        updateBoard(state, payload) {
            let board = state.boards.filter(
                (board) => board.list === payload.list
            )[0];
            console.log(board.backgroundColor, payload.backgroundColor, "mutations");
            board.list = payload.list;
            board.title = payload.title;
            board.backgroundColor = payload.backgroundColor;
            board.color = payload.color;
        },
        setSearch(state, value) {
            state.search = value;
        },
        toggleSorting(state) {
            state.sorting = !state.sorting;
        },
        toggleBoardSorting(state) {
            state.boardSorting = !state.boardSorting;
        },
        toggleAddBoard(state) {
            state.addBoard = !state.addBoard;
        },
        addBoard(state, payload) {
            state.boards.push(payload);
            console.log(state.boards);
        },
        setBoards(state, payload) {
            state.boards = payload;
        },
    },
    actions: {
        addTask({ commit }, newTaskTitle) {
            let newTask = {
                id: Date.now(),
                list: 1,
                title: newTaskTitle,
                done: false,
                dueDone: null,
            };
            db.collection("tasks")
                .add(newTask)
                .then(() => {
                    commit("addTask", newTask);
                    commit("showSnackbar", "Task Added!");
                });
        },
        doneTask({ state, commit }, id) {
            let task = state.tasks.filter((task) => task.id === id)[0];
            db.collection("tasks")
                .doc({ id: id })
                .update({
                    done: !task.done,
                })
                .then(() => {
                    commit("doneTask", id);
                });
        },
        deleteTask({ commit }, id) {
            db.collection("tasks")
                .doc({ id: id })
                .delete()
                .then(() => {
                    commit("deleteTask", id);
                    commit("showSnackbar", "Task Deleted!");
                });
        },
        deleteBoard({ commit }, id) {
            commit("deleteBoard", id);
            commit("showSnackbar", "Board Deleted!");
        },
        updateTasktitle({ commit }, payload) {
            db.collection("tasks")
                .doc({ id: payload.id })
                .update({ title: payload.title })
                .then(() => {
                    commit("updateTasktitle", payload);
                    commit("showSnackbar", "Task updated!");
                });
        },
        updateTaskDueDate({ commit }, payload) {
            db.collection("tasks")
                .doc({ id: payload.id })
                .update({ dueDate: payload.dueDate })
                .then(() => {
                    commit("updateTaskDueDate", payload);
                    commit("showSnackbar", "Date updated!");
                });
        },
        updateBoard({ commit }, payload) {
            commit("updateBoard", payload);
            commit("showSnackbar", "Board updated!");
        },
        setTasks({ commit }, tasks) {
            db.collection("tasks").set(tasks);
            commit("setTasks", tasks);
        },
        setBoards({ state, commit }, boards) {
            db.collection("boards").set(boards);
            commit("setBoards", boards);
        },
        getBoards({ state, commit }) {
            db.collection("boards")
                .get()
                .then((boards) => {
                    if (boards.length === 0) {
                        db.collection("boards")
                            .add(state.defaultBoards[0])
                            .then(() => {
                                commit("addBoard", state.defaultBoards[0]);
                                commit("showSnackbar", "Board Added!");
                            });
                    }
                    commit("setBoards", boards);
                });
        },
        getTasks({ commit }) {
            db.collection("tasks")
                .get()
                .then((tasks) => {
                    console.log(tasks, " getTasks");
                    commit("setTasks", tasks);
                });
        },
        addBoard({ commit }, board) {
            db.collection("boards")
                .add(board)
                .then(() => {
                    commit("addBoard", board);
                    commit("showSnackbar", "Board Added!");
                });
        },
    },
    getters: {
        tasksFiltered(state, id) {
            if (!state.search) {
                return state.tasks;
            }
            return state.tasks.filter((task) =>
                task.title.toLowerCase().includes(state.search.toLowerCase())
            );
        },
    },

}
export default todoModule