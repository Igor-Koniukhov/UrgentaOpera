import axios from "axios";

const authentication = {
    namespaced: true,
    state: {
        registration: {
            name: '',
            phoneNumber: '',
            email: '',
            password: '',
            confirmPassword: ''
        },
        isSignedUp: false,
        signUpError: null,
        isLogin: false,
        loginError: null
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
            state.isLogin = false
        },
        setIsLogin(state) {
            state.isLogin = true
        },
        setLoginError(state, payload) {
            state.loginError = payload.error
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
                localStorage.setItem('userId', response.data.id)
                commit('userRegistration', payload)
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
                console.log("login success", payload)
                console.log("login success", response)
                commit('setIsLogin')
                commit('setLoginError', {
                    error: null,
                })
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