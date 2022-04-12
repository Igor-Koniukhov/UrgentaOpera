<template>
  <v-form
      ref="form"
      v-model="valid"
      lazy-validation
  >

    <v-text-field
        v-model="login.email"
        :rules="emailRules"
        label="E-mail"
        required
    ></v-text-field>

    <v-text-field
        v-model="login.password"
        :rules="passRules"
        label="Password"
        required
    ></v-text-field>

    <v-btn
        :disabled="!valid"
        color="success"
        class="mr-4"
        @click="Login"
    >Login
    </v-btn>




  </v-form>
</template>

<script>
export default {
  data: () => ({
    valid: true,
   login:{
     password: '',
     email: '',
   },
    passRules: [
      v => !!v || 'Password is required',
    ],

    emailRules: [
      v => !!v || 'E-mail is required',
      v => /.+@.+\..+/.test(v) || 'E-mail must be valid',
    ],
    items: [
      'Item 1',
      'Item 2',
    ],
  }),
  methods: {
    Login () {
      this.$refs.form.validate()
      this.$store.dispatch('auth/Login',{
        email: this.login.email,
        password: this.login.password
      })
      this.reset()

    },
    reset () {
      this.$refs.form.reset()
      this.login.email=''
      this.login.password=''
    },

  },
}
</script>