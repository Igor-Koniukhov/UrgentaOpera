import axios from "axios";

const authentication = {
    namespaced: true,
    state: {
        user:{
            id: '' || localStorage.getItem('UserId'),
            name: '' || localStorage.getItem('UserName'),
            email: '' || localStorage.getItem('email')
        },
        registration: {
            name: '',
            phoneNumber: '',
            email: '',
            password: '',
            confirmPassword: ''
        },
        isSignedUp: false,
        signUpError: null,
        isLogin: false || localStorage.getItem('status'),
        loginError: null,
        snackbar: {
            show: false,
            text: "",
        },
    },
    mutations: {
        userRegistration(state, payload) {
            state.registration = payload
        },
        setIsNotSignedUp(state) {
            state.isSignedUp = false
        },
        setIsSignedUp(state) {
            state.isSignedUp = true
        },
        setSignUpError(state, payload) {
            state.signUpError = payload.error
        },
        setIsNotLogin(state) {
            localStorage.removeItem('UserName')
            localStorage.removeItem('UserId')
            localStorage.removeItem('email')
            localStorage.removeItem('status')
            state.isLogin = false
        },
        setIsLogin(state) {
            state.isLogin = localStorage.getItem('status')
        },
        setLoginError(state, payload) {
            state.loginError = payload.error
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

    },
    actions: {
        async userRegistration({commit}, payload) {
            try {
                console.log(payload, " from userRegistration")
                const response = await axios.post('/registration', {
                    name: payload.name,
                    phoneNumber: payload.phoneNumber,
                    email: payload.email,
                    password: payload.password,
                    confirmPassword: payload.confirmPassword
                })
                commit('setIsSignedUp')
                commit('setSignUpError', {
                    error: null,
                })
                console.log(response)
                localStorage.setItem('userId', response.data.id)
                commit('userRegistration', payload)
                commit('showSnackbar', "Success! Now login!")
            } catch (err) {
                commit('setIsNotSignedUp')
                commit('setSignUpError', {
                    error: err.message
                })
            }

        },
        async Login({commit}, payload) {
            try {
                const response = await axios.post('/login', {
                    email: payload.email,
                    password: payload.password,
                })
                console.log(response)
                localStorage.setItem("UserId", response.data.user_id)
                localStorage.setItem("UserName", response.data.user_name)
                localStorage.setItem("status", response.data.status)
                localStorage.setItem("email", response.data.email)
                commit('setIsLogin')
                commit('setLoginError', {
                    error: null,
                })
                commit('showSnackbar', "Success! You are login!")
            } catch (err) {
                console.log(err.message)
                commit('setIsNotLogin')
                commit('setLoginError', {
                    error: err.message,
                })
            }

        }
    },
    getters: {
        isLogin: (state) => state.isLogin
    }

}


export default authentication